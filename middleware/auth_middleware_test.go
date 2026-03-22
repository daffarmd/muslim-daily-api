package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api-go-test/helper"
	"api-go-test/model/web"

	"github.com/julienschmidt/httprouter"
)

func TestRequireAPIKeyAllowsValidKey(t *testing.T) {
	called := false
	handler := RequireAPIKey("secret", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		called = true
		w.WriteHeader(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodPost, "/api/data", nil)
	req.Header.Set("X-API-Key", "secret")
	recorder := httptest.NewRecorder()

	handler(recorder, req, nil)

	if !called {
		t.Fatal("expected next handler to be called")
	}

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}
}

func TestRequireAPIKeyRejectsInvalidKey(t *testing.T) {
	handler := RequireAPIKey("secret", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		t.Fatal("next handler should not be called")
	})

	req := httptest.NewRequest(http.MethodPost, "/api/data", nil)
	req = req.WithContext(helper.ContextWithRequestID(req.Context(), "req-401"))
	recorder := httptest.NewRecorder()

	handler(recorder, req, nil)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, recorder.Code)
	}

	var response web.WebResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Message != "invalid api key" {
		t.Fatalf("expected message %q, got %q", "invalid api key", response.Message)
	}

	if response.RequestID != "req-401" {
		t.Fatalf("expected request id %q, got %q", "req-401", response.RequestID)
	}
}
