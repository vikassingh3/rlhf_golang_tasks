package examples

import (
	"context"
	"fmt"
	"time"
)

// CorrectPitfall3 demonstrates avoiding holding contexts too long.
func CorrectPitfall3() {
	for i := 0; i < 1000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		go longRunningTask(ctx, i)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Main function exiting.")
}
