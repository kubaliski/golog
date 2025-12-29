package logger

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Logger struct {
	serviceName string
}

// NewLogger is a logger constructor
func NewLogger(serviceName string) *Logger {
	log.SetFlags(0)
	return &Logger{serviceName: serviceName}
}

// write creates the log formatted and prints it
func (l *Logger) write(level, serviceName, msg string) {
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("[%s][%s][%s] %s", timestamp, level, serviceName, msg)
	log.Println(logMessage)
}

// Fatal logs a message at fatal level and exits the program
func (l *Logger) Fatal(msg string) {
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("[%s][%s][%s] %s", timestamp, "FATAL", l.serviceName, msg)
	log.Fatal(logMessage)
}

// Log writes a log using the service name from ctx if present, otherwise uses the logger's serviceName
func (l *Logger) Log(ctx context.Context, level, msg string) {
	svc := l.serviceName
	if ctx != nil {
		if s := GetServiceName(ctx); s != "" && s != defaultServiceName {
			svc = s
		}
	}
	l.write(level, svc, msg)
}

// Info is a wrapper for info level log
func (l *Logger) Info(msg string) {
	l.write("INFO", l.serviceName, msg)
}

// Error is a wrapper for error level log
func (l *Logger) Error(msg string) {
	l.write("ERROR", l.serviceName, msg)
}

// Warn is a wrapper for warn level log
func (l *Logger) Warn(msg string) {
	l.write("WARN", l.serviceName, msg)
}

// Debug is a wrapper for debug level log
func (l *Logger) Debug(msg string) {
	l.write("DEBUG", l.serviceName, msg)
}

// FromContext get de serviceName from ctx
func FromContext(ctx context.Context) *Logger {
	return NewLogger(GetServiceName(ctx))
}
