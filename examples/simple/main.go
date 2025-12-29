package main

import (
	"context"
	"time"

	"github.com/kubaliski/golog/pkg/logger"
)

func main() {
	ctx := context.Background()
	ctx = logger.SetServiceName(ctx, "ExampleService")

	// create logger from context
	lg := logger.FromContext(ctx)
	lg.Info("starting application")
	lg.Debug("debug details")

	// log with an overridden service name in context
	ctxSub := logger.SetServiceName(ctx, "SubService")
	lg.Log(ctxSub, "WARN", "a warning from subservice")

	// explicit logger with its own service name
	other := logger.NewLogger("ExplicitService")
	other.Error("an explicit service error")

	// ejemplo de uso de Fatal
	// descomenta la siguiente línea para probar el cierre del programa con Fatal
	// other.Fatal("an unrecoverable error, exiting")

	// small sleep to ensure logs appear in simple examples
	time.Sleep(10 * time.Millisecond)
}
