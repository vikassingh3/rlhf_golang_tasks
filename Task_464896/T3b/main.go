package main

import (
	"encoding/json"
	"log"
	"time"
)

type Logger struct {
	logger *log.Logger
}

// NewLogger returns a new Logger instance with the provided logger.
func NewLogger(l *log.Logger) *Logger {
	return &Logger{logger: l}
}

// Info logs an informational message.
func (l *Logger) Info(message string, keyvals ...interface{}) {
    jsonData, err := json.Marshal(keyvals)
    if err != nil {
        // Handle JSON marshaling error gracefully, e.g., log it as a plain text message.
        l.logger.Printf("Error marshaling JSON: %v, %s\n", err, message)
        return
    }
	l.logger.Printf("INFO: %s %s\n", message, string(jsonData))
}

// Profile logs the execution time of an action.
func (l *Logger) Profile(name string, action func()) {
	start := time.Now()
	action()
	end := time.Now()
	duration := end.Sub(start)
	l.Info("Execution Time", "action", name, "duration", duration.String())
}

func main() {
	// You can customize the logger as needed, e.g., prefix, flags, output destination.
	logger := log.New(log.Writer(), "", log.LstdFlags|log.Lmicroseconds)
	// Create a Logger instance with the custom logger.
	myLogger := NewLogger(logger)

	// Example of logging a simple action with performance profiling
	myLogger.Profile("Sleep Action", func() {
		time.Sleep(2 * time.Second)
	})

	// You can also log regular information messages with key-value pairs.
	myLogger.Info("System Status", "load_average", 1.5, "memory_used", 350000)
} 