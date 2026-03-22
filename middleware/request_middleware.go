package middleware

import (
	"log"
	"net/http"
	"time"

	"api-go-test/helper"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = helper.NewRequestID()
		}

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(helper.ContextWithRequestID(r.Context(), requestID)))
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		log.Printf(
			"request_id=%s method=%s path=%s status=%d duration=%s",
			helper.RequestIDFromContext(r.Context()),
			r.Method,
			r.URL.Path,
			recorder.statusCode,
			time.Since(start),
		)
	})
}
