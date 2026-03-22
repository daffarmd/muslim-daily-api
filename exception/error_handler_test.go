package exception

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"api-go-test/helper"
	"api-go-test/model/web"

	"github.com/go-playground/validator"
)

func TestErrorHandlerNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/data/99", nil)
	req = req.WithContext(helper.ContextWithRequestID(req.Context(), "req-not-found"))
	recorder := httptest.NewRecorder()

	ErrorHandler(recorder, req, NewNotFoundError("data is not found"))

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}

	var response web.WebResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Message != "resource not found" {
		t.Fatalf("expected message %q, got %q", "resource not found", response.Message)
	}
}

func TestErrorHandlerValidationError(t *testing.T) {
	validate := validator.New()
	err := validate.Struct(struct {
		Name string `validate:"required,min=2"`
	}{})
	if err == nil {
		t.Fatal("expected validation error")
	}

	req := httptest.NewRequest(http.MethodPost, "/api/data", nil)
	req = req.WithContext(helper.ContextWithRequestID(req.Context(), "req-bad-request"))
	recorder := httptest.NewRecorder()

	ErrorHandler(recorder, req, err)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	var response web.WebResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Message != "validation failed" {
		t.Fatalf("expected message %q, got %q", "validation failed", response.Message)
	}

	errorsMap, ok := response.Errors.(map[string]any)
	if !ok {
		t.Fatalf("expected errors map, got %#v", response.Errors)
	}

	if errorsMap["name"] != "field is required" {
		t.Fatalf("expected name validation message, got %#v", errorsMap["name"])
	}
}

func TestErrorHandlerBadRequestError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/prayer-times", nil)
	req = req.WithContext(helper.ContextWithRequestID(req.Context(), "req-bad-request-custom"))
	recorder := httptest.NewRecorder()

	ErrorHandler(recorder, req, NewBadRequestError("validation failed", map[string]string{
		"date": "must use format YYYY-MM-DD",
	}))

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	var response web.WebResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	errorsMap, ok := response.Errors.(map[string]any)
	if !ok {
		t.Fatalf("expected errors map, got %#v", response.Errors)
	}

	if errorsMap["date"] != "must use format YYYY-MM-DD" {
		t.Fatalf("unexpected date validation message: %#v", errorsMap["date"])
	}
}

func TestErrorHandlerInternalServerError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/data", nil)
	req = req.WithContext(helper.ContextWithRequestID(req.Context(), "req-internal"))
	recorder := httptest.NewRecorder()

	ErrorHandler(recorder, req, errors.New("boom"))

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, recorder.Code)
	}

	var response web.WebResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Message != "internal server error" {
		t.Fatalf("expected message %q, got %q", "internal server error", response.Message)
	}
}

func TestNotFoundHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	req = req.WithContext(helper.ContextWithRequestID(req.Context(), "req-404"))
	recorder := httptest.NewRecorder()

	NotFoundHandler(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}

func TestMethodNotAllowedHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPatch, "/api/data", nil)
	req = req.WithContext(helper.ContextWithRequestID(req.Context(), "req-405"))
	recorder := httptest.NewRecorder()

	MethodNotAllowedHandler(recorder, req)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}
