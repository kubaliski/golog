package logger

import (
	"bytes"
	"context"
	"log"
	"strings"
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

func captureOutput(f func()) string {
	var buf bytes.Buffer
	orig := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(orig)
	f()
	return buf.String()
}

func TestInfofLogger(t *testing.T) {
	ctx = SetServiceName(context.Background(), "InfofService")
	logger := FromContext(ctx)

	out := captureOutput(func() {
		logger.Infof("Hello %s %d", "world", 1)
	})

	if !strings.Contains(out, "Hello world 1") {
		t.Fatalf("expected formatted message in output, got: %q", out)
	}
	if !strings.Contains(out, "INFO") {
		t.Fatalf("expected level INFO in output, got: %q", out)
	}
	if !strings.Contains(out, "InfofService") {
		t.Fatalf("expected service name InfofService in output, got: %q", out)
	}
}

func TestErrorfWarnfDebugfLogger(t *testing.T) {
	ctx = SetServiceName(context.Background(), "FmtService")
	logger := FromContext(ctx)

	outErr := captureOutput(func() {
		logger.Errorf("error: %v", "oops")
	})
	if !strings.Contains(outErr, "error: oops") || !strings.Contains(outErr, "ERROR") {
		t.Fatalf("unexpected Errorf output: %q", outErr)
	}

	outWarn := captureOutput(func() {
		logger.Warnf("warn: %s", "careful")
	})
	if !strings.Contains(outWarn, "warn: careful") || !strings.Contains(outWarn, "WARN") {
		t.Fatalf("unexpected Warnf output: %q", outWarn)
	}

	outDebug := captureOutput(func() {
		logger.Debugf("debug: %d", 42)
	})
	if !strings.Contains(outDebug, "debug: 42") || !strings.Contains(outDebug, "DEBUG") {
		t.Fatalf("unexpected Debugf output: %q", outDebug)
	}
}

func TestLogfUsesCtxServiceName(t *testing.T) {
	// logger has default service name but ctx should override
	logger := NewLogger("DefaultService")
	ctx := SetServiceName(context.Background(), "CtxService")

	out := captureOutput(func() {
		logger.Logf(ctx, "INFO", "msg %s", "ok")
	})

	if !strings.Contains(out, "msg ok") {
		t.Fatalf("expected formatted message in output, got: %q", out)
	}
	if !strings.Contains(out, "CtxService") {
		t.Fatalf("expected ctx service name in output, got: %q", out)
	}
}
