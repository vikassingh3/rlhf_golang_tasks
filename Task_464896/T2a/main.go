package main

import (
	"fmt"
	"time"
)

type Logger struct {
	logCallback func(string)
}

// SetLogCallback sets the callback function for logging.
func (l *Logger) SetLogCallback(callback func(string)) {
	l.logCallback = callback
}

// Log logs a message using the set callback.
func (l *Logger) Log(message string) {
	if l.logCallback != nil {
		l.logCallback(message)
	}
}

// LogAction logs a message and measures its execution time using the set callback.
func (l *Logger) LogAction(name string, action func()) {
	start := time.Now()
	action()
	end := time.Now()
	duration := end.Sub(start)
	l.Log(fmt.Sprintf("Action '%s' executed in %s", name, duration))
}

func consoleLogger(message string) {
	fmt.Println("Console Log:", message)
}

func main() {
	// Create a Logger instance.
	logger := &Logger{}
	// Set the console logger as the callback.
	logger.SetLogCallback(consoleLogger)

	// Example of logging a simple action
	logger.LogAction("Example Action", func() {
		time.Sleep(2 * time.Second) // Simulate some work
	})

	// You can still log normal messages
	logger.Log("This is a regular log message.")
}