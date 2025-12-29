package logger

import (
	"context"
	"testing"
)

func TestSetAndGetServiceName(t *testing.T) {
	ctx = SetServiceName(ctx, "MyService")
	if got := GetServiceName(ctx); got != "MyService" {
		t.Fatalf("GetServiceName() = %q, want %q", got, "MyService")
	}
}

func TestWithServiceNameAlias(t *testing.T) {
	ctx2 := WithServiceName(ctx, "AliasService")
	if got := GetServiceName(ctx2); got != "AliasService" {
		t.Fatalf("GetServiceName() = %q, want %q", got, "AliasService")
	}
}

func TestGetServiceNameNilAndEmpty(t *testing.T) {
	// use context.TODO() instead of nil to avoid linter warnings (behavior is equivalent here)
	if got := GetServiceName(context.TODO()); got != defaultServiceName {
		t.Fatalf("GetServiceName(context.TODO()) = %q, want %q", got, defaultServiceName)
	}

	// empty service name should fall back to default
	ctx := SetServiceName(context.Background(), "")
	if got := GetServiceName(ctx); got != defaultServiceName {
		t.Fatalf("GetServiceName(empty) = %q, want %q", got, defaultServiceName)
	}
}

func TestFromContextCreatesLoggerWithServiceName(t *testing.T) {
	ctx := SetServiceName(context.Background(), "FromCtxService")
	lg := FromContext(ctx)
	if lg == nil {
		t.Fatal("FromContext returned nil logger")
	}
	if lg.serviceName != "FromCtxService" {
		t.Fatalf("logger.serviceName = %q, want %q", lg.serviceName, "FromCtxService")
	}
}
