package helper

import "testing"

func TestNewRequestID(t *testing.T) {
	requestID := NewRequestID()

	if len(requestID) != 24 {
		t.Fatalf("expected request id length to be 24, got %d", len(requestID))
	}
}

func TestContextWithRequestID(t *testing.T) {
	ctx := ContextWithRequestID(t.Context(), "req-123")

	if got := RequestIDFromContext(ctx); got != "req-123" {
		t.Fatalf("expected request id %q, got %q", "req-123", got)
	}
}
