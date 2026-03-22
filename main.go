package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"api-go-test/app"
	"api-go-test/controller"
	"api-go-test/exception"
	"api-go-test/helper"
	"api-go-test/middleware"
	"api-go-test/repository"
	"api-go-test/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config := app.LoadConfig()

	if strings.TrimSpace(config.DBDSN) == "" {
		log.Fatal("DB_DSN is required")
	}

	db, err := app.OpenDB(config.DBDSN)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	muslimRepository := repository.NewMuslimRepository(db)
	muslimService := service.NewMuslimService(muslimRepository)
	muslimController := controller.NewMuslimController(muslimService)
	healthController := controller.NewHealthController(config, db)

	router := httprouter.New()

	router.GET("/health", healthController.Health)
	router.GET("/healthz", healthController.Health)
	router.GET("/readyz", healthController.Ready)
	router.GET("/duas", muslimController.FindDuas)
	router.GET("/duas/random", muslimController.FindRandomDua)
	router.GET("/prayer-times", muslimController.FindPrayerTime)

	router.PanicHandler = exception.ErrorHandler
	router.NotFound = http.HandlerFunc(exception.NotFoundHandler)
	router.MethodNotAllowed = http.HandlerFunc(exception.MethodNotAllowedHandler)

	server := http.Server{
		Addr:         config.Address(),
		Handler:      middleware.RequestID(middleware.Logger(router)),
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	go func() {
		log.Printf("running at http://localhost%s env=%s", config.Address(), config.AppEnv)

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	helper.PanicIfErr(err)
}
