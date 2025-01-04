package main

import (
	"log"
	"os"
	"regexp"
)

// MaskSensitiveData uses regex to mask sensitive patterns
func MaskSensitiveData(input string) string {
	// Define regular expressions for sensitive patterns
	regexes := []*regexp.Regexp{
		regexp.MustCompile(`\b\d{4}-\d{4}-\d{4}-\d{4}\b`), // Credit card pattern
		regexp.MustCompile(`\b\d{9}\b`),                 // SSN pattern
	}

	for _, regex := range regexes {
		// Mask if a sensitive pattern is found
		input = regex.ReplaceAllString(input, "***")
	}

	return input
}

func main() {
	// Create a file handle
	logFile, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Set the output of the logger to the file handle
	log.SetOutput(logFile)

	// Log sensitive information after masking
	sensitiveInfo := "Credit card number: 1234-5678-9012-3456"
	log.Println(MaskSensitiveData(sensitiveInfo))
}