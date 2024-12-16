package examples

import (
	"context"
	"fmt"
	"time"
)

// Pitfall3 demonstrates holding contexts too long.
func Pitfall3() {
	for i := 0; i < 1000; i++ {
		ctx := getContext() // Context is pooled and not cancelled
		go longRunningTask(ctx, i)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Main function exiting.")
}

func getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}
