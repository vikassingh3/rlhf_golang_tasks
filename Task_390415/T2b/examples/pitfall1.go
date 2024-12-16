package examples

import (
	"context"
	"fmt"
	"time"
)

// Pitfall1 demonstrates not cancelling contexts properly.
func Pitfall1() {
	ctx := context.Background()

	for i := 0; i < 5; i++ {
		go longRunningTask(ctx, i)
	}

	time.Sleep(2 * time.Second) // Simulate main function exiting early
	fmt.Println("Main function exiting.")
}

func longRunningTask(ctx context.Context, id int) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Task", id, "is running.")
		case <-ctx.Done():
			fmt.Println("Task", id, "cancelled.")
			return
		}
	}
}
