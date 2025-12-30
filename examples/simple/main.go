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
	// existing non-formatted logs (keep for compatibility)
	lg.Info("starting application")
	lg.Debug("debug details")

	// new formatted logs to demonstrate the `*f` methods
	lg.Infof("starting application version %s", "1.0.0")
	lg.Debugf("debug details: enabled=%v", true)

	// log with an overridden service name in context
	ctxSub := logger.SetServiceName(ctx, "SubService")
	// original non-formatted log
	lg.Log(ctxSub, "WARN", "a warning from subservice")
	// formatted variant
	lg.Logf(ctxSub, "WARN", "a warning from %s", "subservice")

	// explicit logger with its own service name
	other := logger.NewLogger("ExplicitService")
	// original non-formatted log
	other.Error("an explicit service error")
	// formatted variant
	other.Errorf("an explicit service error: code=%d", 500)

	// example of fatal log (uncomment to test)
	// other.Fatal("an unrecoverable error, exiting")

	// small sleep to ensure logs appear in simple examples
	time.Sleep(10 * time.Millisecond)
}
