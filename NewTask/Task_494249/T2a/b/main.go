package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// MaskSensitiveData replaces sensitive strings with '***'
func MaskSensitiveData(input string) string {
	sensitiveStrings := []string{
		"password",
		"token",
		"secret",
		"key",
		"cc", // credit card
	}

	for _, sensitive := range sensitiveStrings {
		if strings.Contains(strings.ToLower(input), strings.ToLower(sensitive)) {
			return strings.Replace(input, sensitive, "***", -1)
		}
	}

	return input
}

// LogMasked logs a message after masking sensitive data
func LogMasked(format string, args ...interface{}) {
	maskedMessage := MaskSensitiveData(fmt.Sprintf(format, args...))
	log.Println(maskedMessage)
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

	// Log sensitive information using the custom logging function
	sensitiveInfo := "User entered password: mysecretpassword"
	LogMasked(sensitiveInfo)
}