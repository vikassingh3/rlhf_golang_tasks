package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Simulate a database operation that may timeout
func SimulateDBTimeout(ctx context.Context) (string, error) {
	// Simulate random time to complete the operation
	sleepDuration := time.Duration(rand.Intn(5)) * time.Second

	// Create a new channel to receive result or error
	resultCh := make(chan string)
	errorCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errorCh)

		// Simulate some work
		time.Sleep(sleepDuration)

		// Check if the context has been canceled (due to timeout)
		select {
		case <-ctx.Done():
			errorCh <- ctx.Err()
			return
		default:
		}

		// Simulate returning a result
		resultCh <- "Result from Database"
	}()

	// Wait for the result or error
	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errorCh:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	// Set the timeout duration
	timeoutDuration := 2 * time.Second

	// Create a context with a deadline
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	// Call the SimulateDBTimeout function with the context
	result, err := SimulateDBTimeout(ctx)

	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("Timeout occurred during database operation.")
			// Handle timeout scenario here
			// You can retry the operation, log the error, or take other appropriate actions
		} else {
			fmt.Println("Error occurred during database operation:", err)
		}
		return
	}

	fmt.Println("Result:", result)
}

