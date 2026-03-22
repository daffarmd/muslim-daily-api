package helper

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func NewRequestID() string {
	buffer := make([]byte, 12)
	if _, err := rand.Read(buffer); err != nil {
		return "unknown-request-id"
	}

	return hex.EncodeToString(buffer)
}

func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func RequestIDFromContext(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDKey).(string)
	return requestID
}
