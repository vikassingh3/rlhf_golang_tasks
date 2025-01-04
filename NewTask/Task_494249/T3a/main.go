package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// MaskSensitiveData masks sensitive information using regex patterns.
func MaskSensitiveData(input string) string {
	patterns := []string{
		`(?i)password:\s*\S+`,                          // Matches "password: <value>"
		`(?i)token:\s*\S+`,                             // Matches "token: <value>"
		`(?i)key:\s*\S+`,                               // Matches "key: <value>"
		`\b\d{4}-\d{4}-\d{4}-\d{4}\b`,                  // Matches credit card numbers
		`\b[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\b`, // Matches email addresses
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		input = re.ReplaceAllString(input, "***")
	}

	return input
}

// LogWithMasking logs a message after masking sensitive data.
func LogWithMasking(message string) {
	log.Println(MaskSensitiveData(message))
}

func main() {
	// Create or open log file
	logFile, err := os.OpenFile("./logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Set logger to output to file
	log.SetOutput(logFile)

	// Example log with sensitive data
	sensitiveInfo := "User entered password: mySecret123, token: abcd1234, email: user@example.com"
	LogWithMasking(sensitiveInfo)

	// Access log data for GDPR purposes
	_, err = logFile.Seek(0, 0) // Move to the start of the file
	if err != nil {
		log.Fatalf("Error seeking log file: %v", err)
	}

	scanner := bufio.NewScanner(logFile)
	var logEntry string
	for scanner.Scan() {
		logEntry += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading log file: %v", err)
	}

	// Print log entry for demonstration purposes
	log.Println("Log Entry:\n" + logEntry)

	// Handle GDPR request: user requests to view their logs
	userLogs := strings.Contains(logEntry, "user@example.com")
	if userLogs {
		log.Println("User 'user@example.com' found in logs.")
	} else {
		log.Println("User 'user@example.com' not found in logs.")
	}
}
