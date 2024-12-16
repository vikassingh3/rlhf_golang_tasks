package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func cancelContextExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensures that the context is cancelled

	// Simulate some operation
	go func() {
		time.Sleep(2 * time.Second)
		cancel() // Cancelling the context after work is done
	}()

	// Simulate work that checks the context
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("Completed work!")
	}
}

func doWork(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("Work cancelled:", ctx.Err())
		return
	default:
		fmt.Println("Doing work...")
	}
}

func wrongContextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Ensure we cancel the context

	// Pass the context to doWork
	doWork(ctx)

	time.Sleep(2 * time.Second) // Wait to observe the output
}

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping due to cancellation.")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func longRunningTaskExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the context is cancelled

	longRunningTask(ctx) // This will check for cancellation
}

func processItem(ctx context.Context, item string) {
	select {
	case <-ctx.Done():
		fmt.Println("Cancelled processing:", item)
		return
	default:
		fmt.Printf("Processing item: %s\n", item)
	}
}

func contextInLoopsExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	items := []string{"item1", "item2", "item3"}
	for _, item := range items {
		processItem(ctx, item) // Pass the same context
	}
}

func processRequest(ctx context.Context) error {
	time.Sleep(2 * time.Second) // Simulate some work
	return ctx.Err()           // Return the context error if any
}

func handleErrorsExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := processRequest(ctx); err != nil {
		fmt.Println("Error processing request:", err)
	}
}

func main() {
	fmt.Println("Starting cancelContextExample...")
	cancelContextExample()

	fmt.Println("\nStarting wrongContextExample...")
	wrongContextExample()

	fmt.Println("\nStarting longRunningTaskExample...")
	longRunningTaskExample()

	fmt.Println("\nStarting contextInLoopsExample...")
	contextInLoopsExample()

	fmt.Println("\nStarting handleErrorsExample...")
	handleErrorsExample()
}
