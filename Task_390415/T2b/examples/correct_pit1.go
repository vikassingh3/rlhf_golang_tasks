package examples

import (
	"context"
	"fmt"
	"time"
)

// CorrectPitfall2 demonstrates properly propagating contexts.
func CorrectPitfall2() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go parentTaskCorrect(ctx, i)
	}

	time.Sleep(2 * time.Second) // Simulate main function exiting early
	fmt.Println("Main function exiting.")
}

func parentTaskCorrect(ctx context.Context, id int) {
	go childTaskCorrect(ctx, id) // Passing the context
	select {
	case <-time.After(2 * time.Second):
	case <-ctx.Done():
	}
}

func childTaskCorrect(ctx context.Context, id int) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Child task", id, "is running.")
		case <-ctx.Done():
			fmt.Println("Child task", id, "cancelled.")
			return
		}
	}
}
