package main

import (
	"log"
	"os"
	"strings"
)

// MaskSensitiveData replaces sensitive strings with '***'
func MaskSensitiveData(input string) string {
	// Define sensitive strings to mask
	sensitiveStrings := []string{
		"password",
		"token",
		"secret",
		"key",
		"cc", // credit card
	}

	for _, sensitive := range sensitiveStrings {
		// Mask if sensitive string is found
		if strings.Contains(strings.ToLower(input), strings.ToLower(sensitive)) {
			return strings.Replace(input, sensitive, "***", -1)
		}
	}

	return input
}

func main() {
	// Create a file handle
	logFile, err := os.OpenFile("./logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Set the output of the logger to the file handle
	log.SetOutput(logFile)

	// Log sensitive information after masking
	sensitiveInfo := "User entered password: mysecretpassword"
	log.Println(MaskSensitiveData(sensitiveInfo))
}