package logger

import (
	"context"
	"testing"
)

var (
	ctx = context.Background()
)

func TestLogger(t *testing.T) {
	ctx = SetServiceName(ctx, "TestService")

	logger := FromContext(ctx)

	tests := []struct {
		level   string
		message string
	}{
		{"INFO", "This is an info message"},
		{"ERROR", "This is an error message"},
		{"WARN", "This is a warn message"},
		{"DEBUG", "This is a debug message"},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			logger.Log(ctx, tt.level, tt.message)
			// capture output and assert if needed
		})
	}
}

func TestInfoLogger(t *testing.T) {
	ctx = SetServiceName(ctx, "InfoService")

	logger := FromContext(ctx)

	logger.Info("This is a info log")
}

func TestErrorLogger(t *testing.T) {
	ctx = SetServiceName(ctx, "ErrorService")

	logger := FromContext(ctx)

	logger.Error("This is an erro log")
}

func TestWarnLogger(t *testing.T) {
	ctx = SetServiceName(ctx, "Warn Service")

	logger := FromContext(ctx)

	logger.Warn("This is a warn log")
}

func TestDebugLogger(t *testing.T) {
	ctx = SetServiceName(ctx, "Debug Service")

	logger := FromContext(ctx)

	logger.Debug("This is a debug log")
}

func TestSetServiceName(t *testing.T) {
	serviceName := "MyService"
	ctx = SetServiceName(ctx, serviceName)

	if got := GetServiceName(ctx); got != serviceName {
		t.Errorf("GetServiceName() = %v, want %v", got, serviceName)
	}
}
