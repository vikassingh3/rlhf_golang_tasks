package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	// Simulate CPU-intensive work
	for i := 0; i < 10; i++ {
		_ = i * i
	}
	fmt.Printf("Worker %d completed.\n", id)
}

func main() {
	var wg sync.WaitGroup
	const numWorkers = 5
	const numTasks = 10

	// Channel to limit the number of goroutines
	c := make(chan struct{}, numWorkers)

	// Start tasks
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()

			// Acquire a slot in the channel
			c <- struct{}{}
			defer func() { <-c }() // Release the slot in the channel

			worker(taskID)
		}(i)
	}

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All tasks completed. Application exited.")
}
