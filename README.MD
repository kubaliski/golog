# golog

A simple logging package for Go that provides structured log output with service name context propagation.

## Disclaimer

This is a personal utility package created to standardize logging across my own projects according to my preferences. It is not intended for production use and should not be considered a professional-grade logging solution. Use at your own discretion.

## Installation

```bash
go get github.com/kubaliski/golog
```

## Features

- Context-based service name propagation
- Structured log format with timestamp, level, and service name
- Simple API with common log levels (INFO, ERROR, WARN, DEBUG)

## Log Format

All logs follow this structure:

```
[timestamp][LEVEL][service-name] message
```

Example output:

```
[2025-01-15T10:30:00Z][INFO][my-service] Application started
[2025-01-15T10:30:01Z][ERROR][my-service] Connection failed
```

## Usage

### Basic Usage

```go
package main

import (
    "github.com/kubaliski/golog/pkg/logger"
)

func main() {
    log := logger.NewLogger("my-service")

    log.Info("Application started")
    log.Debug("Loading configuration")
    log.Warn("Configuration file not found, using defaults")
    log.Error("Failed to connect to database")
}
```

### Using Context for Service Name

```go
package main

import (
    "context"

    "github.com/kubaliski/golog/pkg/logger"
)

func main() {
    ctx := context.Background()
    ctx = logger.SetServiceName(ctx, "my-service")

    log := logger.FromContext(ctx)
    log.Info("Application started")
}
```

### Propagating Service Name Across Functions

```go
func main() {
    ctx := logger.SetServiceName(context.Background(), "api-gateway")
    handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
    log := logger.FromContext(ctx)
    log.Info("Processing request")
}
```

### Using Log with Context Override

```go
func main() {
    log := logger.NewLogger("default-service")

    ctx := logger.SetServiceName(context.Background(), "override-service")
    log.Log(ctx, "INFO", "This will use override-service as the service name")
}
```

## API Reference

### Logger

| Method                                        | Description                                          |
| --------------------------------------------- | ---------------------------------------------------- |
| `NewLogger(serviceName string)`               | Creates a new logger with the given service name     |
| `FromContext(ctx context.Context)`            | Creates a logger using the service name from context |
| `Info(msg string)`                            | Logs a message at INFO level                         |
| `Error(msg string)`                           | Logs a message at ERROR level                        |
| `Warn(msg string)`                            | Logs a message at WARN level                         |
| `Debug(msg string)`                           | Logs a message at DEBUG level                        |
| `Log(ctx context.Context, level, msg string)` | Logs with context-based service name override        |

### Context Functions

| Function                     | Description                                                            |
| ---------------------------- | ---------------------------------------------------------------------- |
| `SetServiceName(ctx, name)`  | Stores the service name in context                                     |
| `WithServiceName(ctx, name)` | Alias for SetServiceName                                               |
| `GetServiceName(ctx)`        | Retrieves the service name from context (returns "unknown" if not set) |

## Running Tests

```bash
go test ./... -v
```

## License

MIT
