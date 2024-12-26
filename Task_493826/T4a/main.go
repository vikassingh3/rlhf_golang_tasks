package main

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

// Improved error handling with slice processing
func processSlice(numbers []int) error {
	logrus.WithFields(logrus.Fields{
		"length": len(numbers),
	}).Info("Processing slice")

	if numbers == nil {
		return fmt.Errorf("slice is nil")
	}

	if len(numbers) == 0 {
		return fmt.Errorf("slice is empty")
	}

	// Perform operations on the slice
	logrus.WithFields(logrus.Fields{
		"index": 0,
		"value": numbers[0],
	}).Info("Accessing element")
	fmt.Println("Accessing element at index 0:", numbers[0])
	return nil
}

// Main function demonstrating enhanced logging, error handling, and Sentry integration
func main() {
	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://YOUR_SENTRY_DSN@example.com/1", // Replace with your Sentry DSN
	})
	if err != nil {
		fmt.Println("Getting error:", err.Error())
		return
	}
	defer sentry.Flush(2) // Flush any pending events before exiting

	// Enhanced logging setup
	logrus.SetLevel(logrus.InfoLevel)

	// Simulate slice processing
	var numbers []int
	userID := 123

	// Error handling for slice processing
	err = processSlice(numbers)
	if err != nil {
		logrus.Errorf("Error processing slice: %v", err)

		// Capture the error with Sentry
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetUser(sentry.User{ID: fmt.Sprintf("%d", userID)})
			scope.SetExtra("UserID", userID)
			sentry.CaptureException(fmt.Errorf("User with ID %d encountered error: %v", userID, err))
		})
		return
	}

	logrus.Info("Processing completed successfully")
	fmt.Println("Processing completed successfully")
}
