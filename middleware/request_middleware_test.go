package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api-go-test/helper"
)

func TestRequestIDUsesExistingHeader(t *testing.T) {
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(helper.RequestIDFromContext(r.Context())))
	}))

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	req.Header.Set("X-Request-ID", "existing-id")
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if got := recorder.Header().Get("X-Request-ID"); got != "existing-id" {
		t.Fatalf("expected response header request id %q, got %q", "existing-id", got)
	}

	if got := strings.TrimSpace(recorder.Body.String()); got != "existing-id" {
		t.Fatalf("expected request id in context %q, got %q", "existing-id", got)
	}
}

func TestRequestIDGeneratesHeaderWhenMissing(t *testing.T) {
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(helper.RequestIDFromContext(r.Context())))
	}))

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	requestID := recorder.Header().Get("X-Request-ID")
	if len(requestID) != 24 {
		t.Fatalf("expected generated request id length 24, got %d", len(requestID))
	}

	if got := strings.TrimSpace(recorder.Body.String()); got != requestID {
		t.Fatalf("expected handler to receive request id %q, got %q", requestID, got)
	}
}

func TestLoggerWritesRequestLog(t *testing.T) {
	var buffer bytes.Buffer
	originalWriter := log.Writer()
	log.SetOutput(&buffer)
	defer log.SetOutput(originalWriter)

	handler := RequestID(Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = io.WriteString(w, "created")
	})))

	req := httptest.NewRequest(http.MethodGet, "/api/data", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	logOutput := buffer.String()
	if !strings.Contains(logOutput, "method=GET") {
		t.Fatalf("expected log output to contain method, got %q", logOutput)
	}

	if !strings.Contains(logOutput, "path=/api/data") {
		t.Fatalf("expected log output to contain path, got %q", logOutput)
	}

	if !strings.Contains(logOutput, "status=201") {
		t.Fatalf("expected log output to contain status, got %q", logOutput)
	}
}
