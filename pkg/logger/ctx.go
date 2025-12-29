package logger

import (
	"context"
)

type ctxKey struct{}

const defaultServiceName = "unknown"

// SetServiceName stores the service name into the context.
func SetServiceName(ctx context.Context, serviceName string) context.Context {
	return context.WithValue(ctx, ctxKey{}, serviceName)
}

// WithServiceName is an alias for SetServiceName.
func WithServiceName(ctx context.Context, serviceName string) context.Context {
	return SetServiceName(ctx, serviceName)
}

// GetServiceName retrieves the service name from the context or returns "unknown".
func GetServiceName(ctx context.Context) string {
	if ctx == nil {
		return defaultServiceName
	}
	if v, ok := ctx.Value(ctxKey{}).(string); ok && v != "" {
		return v
	}
	return defaultServiceName
}
