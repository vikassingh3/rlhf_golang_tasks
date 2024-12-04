package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func someLongRunningTask(ctx context.Context, name string) {
	log.Printf("Task %s started\n", name)
	defer log.Printf("Task %s finished\n", name)

	select {
	case <-time.After(5 * time.Second):
		fmt.Printf("Task %s completed successfully\n", name)
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Printf("Task %s cancelled: %v\n", name, err)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go someLongRunningTask(ctx, "Task 1")
	go someLongRunningTask(ctx, "Task 2")

	time.Sleep(2 * time.Second)

	// Simulate cancellation of the context
	cancel()

	// Wait for all goroutines to finish
	time.Sleep(3 * time.Second)
}