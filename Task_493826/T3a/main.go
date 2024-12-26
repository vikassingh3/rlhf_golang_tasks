package main

import (
	"fmt"
	"log"
	"github.com/getsentry/sentry-go"
)

func main() {
	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://YOUR_SENTRY_DSN@example.com/1",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sentry.Flush(2)

	var numbers []int

	// Attempting to access a nil slice
	if numbers != nil {
		log.Printf("Accessing element at index 0: %v", numbers[0])
	} else {
		log.Println("Slice is nil")
		sentry.CaptureMessage("Slice is nil")
	}

	// Example of capturing an error and sending it to Sentry
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			log.Printf("Error: %v", r)
			sentry.CaptureException(fmt.Errorf("panic: %v", r))
		}
	}()
}