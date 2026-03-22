package controller

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"api-go-test/app"
	"api-go-test/helper"

	"github.com/julienschmidt/httprouter"
)

type HealthController struct {
	Config app.Config
	DB     *sql.DB
}

func NewHealthController(cfg app.Config, db *sql.DB) *HealthController {
	return &HealthController{
		Config: cfg,
		DB:     db,
	}
}

func (c *HealthController) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.WriteSuccess(w, http.StatusOK, "service is healthy", map[string]any{
		"env":                 c.Config.AppEnv,
		"database_configured": c.DB != nil,
	})
}

func (c *HealthController) Ready(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if c.DB == nil {
		helper.WriteSuccess(w, http.StatusOK, "service is ready", map[string]string{
			"database": "not configured",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	if err := c.DB.PingContext(ctx); err != nil {
		helper.WriteError(w, r, http.StatusServiceUnavailable, "service is not ready", map[string]string{
			"database": "database connection is unavailable",
		})
		return
	}

	helper.WriteSuccess(w, http.StatusOK, "service is ready", map[string]string{
		"database": "ok",
	})
}
