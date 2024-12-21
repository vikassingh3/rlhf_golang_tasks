package main

import (
	"fmt"
	"os"
	"time"
)

// Logger defines a method for logging messages.
type Logger interface {
	Log(message string, callback func(string) error) error
}

// ConsoleLogger implements the Logger interface to log to the console.
type ConsoleLogger struct{}

// Log logs a message to the console.
func (cl *ConsoleLogger) Log(message string, callback func(string) error) error {
	_, err := fmt.Printf("[%v] %v\n", time.Now(), message)
	if err != nil {
		return err
	}
	if callback != nil {
		return callback(message)
	}
	return nil
}

// FileLogger implements the Logger interface to log to a file.
type FileLogger struct {
	file *os.File
}

// NewFileLogger creates a new FileLogger instance that logs to the specified file path.
func NewFileLogger(filePath string) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

// Log logs a message to the file.
func (fl *FileLogger) Log(message string, callback func(string) error) error {
	_, err := fmt.Fprintf(fl.file, "[%v] %v\n", time.Now(), message)
	if err != nil {
		return err
	}
	if callback != nil {
		return callback(message)
	}
	return nil
}

// Close closes the underlying file.
func (fl *FileLogger) Close() error {
	return fl.file.Close()
}

func main() {
	// Log to console
	consoleLogger := &ConsoleLogger{}
	consoleLogger.Log("This is a console log", nil)

	// Log to file
	fileLogger, err := NewFileLogger("example.log")
	if err != nil {
		fmt.Println("Error creating file logger:", err)
		return
	}
	defer fileLogger.Close()

	// Log with a callback function that transforms the message
	callback := func(message string) error {
		// Example: Capitalize the first letter of the message
		firstChar := message[0]
		if firstChar >= 'a' && firstChar <= 'z' {
			message = string(firstChar-'a'+'A') + message[1:]
		}
		// Log the transformed message to console
		fmt.Printf("Callback: [%v] Transformed: %v\n", time.Now(), message)
		return nil
	}

	err = consoleLogger.Log("This is a lowercase message", callback)
	if err != nil {
		fmt.Println("Error logging:", err)
		return
	}

	err = fileLogger.Log("This is a lowercase message", callback)
	if err != nil {
		fmt.Println("Error logging:", err)
		return
	}
}