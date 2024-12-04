package main

import (
	"context"
	"fmt"
	"time"
)

func processTask(ctx context.Context, task string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task", task, "timed out")
			return
		default:
			fmt.Println("Task", task, "processing...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go processTask(ctx, "Task 1")
	go processTask(ctx, "Task 2")

	time.Sleep(3 * time.Second)

	// Cancelling the context for Task 2
	go func() {
		select {
		case <-time.After(2 * time.Second):
			cancel()
		}
	}()

	// Wait for the main goroutine to finish
	time.Sleep(6 * time.Second)
}