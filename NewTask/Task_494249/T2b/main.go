package main

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type CustomLogFormatter struct {
	logrus.Formatter
}

func (f *CustomLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Obfuscate sensitive data by replacing it with "******"
	message := entry.Message
	message = strings.ReplaceAll(message, "your_sensitive_data", "******")

	// Format the log entry using the base formatter
	b, err := f.Formatter.Format(entry)
	if err != nil {
		return nil, err
	}

	// Append the obfuscated message to the formatted entry
	return append(b, []byte(fmt.Sprintf(" %s\n", message))...), nil
}

func main() {
	// Create a new Logrus logger
	logger := logrus.New()

	// Set the custom formatter
	logger.SetFormatter(&CustomLogFormatter{
		Formatter: &logrus.JSONFormatter{},
	})

	// Log a message with sensitive data
	logger.Println("This is a log message with sensitive data: your_sensitive_data")
}
