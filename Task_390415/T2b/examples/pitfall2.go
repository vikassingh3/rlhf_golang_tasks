package examples

import (
	"context"
	"fmt"
	"time"
)

// Pitfall2 demonstrates not propagating contexts.
func Pitfall2() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go parentTask(ctx, i)
	}

	time.Sleep(2 * time.Second) // Simulate main function exiting early
	fmt.Println("Main function exiting.")
}

func parentTask(ctx context.Context, id int) {
	go childTask(id) // Not passing the context
	select {
	case <-time.After(2 * time.Second):
	case <-ctx.Done():
	}
}

func childTask(id int) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Child task", id, "is running.")
		}
	}
}
