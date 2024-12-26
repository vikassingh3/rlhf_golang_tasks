package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

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

	// Simulate a slice-related error with context
	var users []string
	userID := 123

	if len(users) > 0 {
		fmt.Println("User found:", users[0])
	} else {
		// Capture the error with context
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetUser(sentry.User{ID: fmt.Sprintf("%d", userID)})
			scope.SetExtra("UserID", userID)
			sentry.CaptureException(fmt.Errorf("User with ID %d not found", userID))
		})
		fmt.Println("User not found:", userID)
	}
}
