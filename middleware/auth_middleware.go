package middleware

import (
	"net/http"

	"api-go-test/helper"

	"github.com/julienschmidt/httprouter"
)

func RequireAPIKey(apiKey string, next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if r.Header.Get("X-API-Key") == apiKey {
			next(w, r, params)
			return
		}

		helper.WriteError(w, r, http.StatusUnauthorized, "invalid api key", nil)
	}
}
