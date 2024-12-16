package examples

import (
	"context"
	"fmt"
	"time"
)

// CorrectPitfall1 demonstrates properly cancelling contexts.
func CorrectPitfall1() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go longRunningTask(ctx, i)
	}

	time.Sleep(2 * time.Second) // Simulate main function exiting early
	fmt.Println("Main function exiting.")
}
