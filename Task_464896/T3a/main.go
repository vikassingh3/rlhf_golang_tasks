package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type StructuredLogger struct {
	logger zerolog.Logger
}

// NewStructuredLogger initializes a new structured logger.
func NewStructuredLogger() *StructuredLogger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &StructuredLogger{logger: logger}
}

// Log logs a message with the specified level and optional fields.
func (l *StructuredLogger) Log(message string, level zerolog.Level, fields map[string]interface{}) {
	event := l.logger.WithLevel(level)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// LogAction logs a message with the specified level and measures its execution time.
func (l *StructuredLogger) LogAction(name string, action func(), level zerolog.Level) {
	start := time.Now()
	action()
	duration := time.Since(start)
	l.Log(fmt.Sprintf("Action '%s' executed in %s", name, duration), level, nil)
}

func exampleWork() {
	time.Sleep(2 * time.Second) // Simulate some work
}

func main() {
	// Create a structured logger instance.
	logger := NewStructuredLogger()

	// Example of logging a simple action
	logger.LogAction("Example Action", exampleWork, zerolog.InfoLevel)

	// Add custom metadata to a log message
	customMetadata := map[string]interface{}{
		"application": "Sample App",
		"environment": "development",
		"timestamp":   time.Now(),
	}
	logger.Log("This is a structured log message.", zerolog.InfoLevel, customMetadata)

	// Log additional system information
	systemInfo := map[string]interface{}{
		"pid":  os.Getpid(),
		"host": os.Getenv("HOSTNAME"),
		"ip":   strings.TrimSpace(os.Getenv("IPADDRESS")),
	}
	logger.Log("System information logged.", zerolog.InfoLevel, systemInfo)
}
