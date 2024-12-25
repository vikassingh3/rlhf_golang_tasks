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

	// Simulate slice-related error
	var numbers []int
	if len(numbers) > 0 && numbers[0] == 0 { // Check if slice is not empty before accessing elements
		fmt.Println("Numbers are not zero")
	} else {
		fmt.Println("Slice is empty or does not start with zero")
	}

	// Example of capturing an error and sending it to Sentry
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			sentry.CaptureException(fmt.Errorf("panic: %v", r))
		}
	}()
}
