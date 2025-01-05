package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate a data aggregation task
func aggregateData(id int, wg *sync.WaitGroup, data *[]int) {
	defer wg.Done() // Ensure we signal the WaitGroup when done
	time.Sleep(time.Duration(id) * 250 * time.Millisecond) // Simulate work
	fmt.Printf("Aggregating data for task %d...\n", id)
	*data = append(*data, id*10)
}

func main() {
	var wg sync.WaitGroup
	data := make([]int, 0)

	// Create and start multiple data aggregation tasks
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go aggregateData(i, &wg, &data) // Start a goroutine
	}

	fmt.Println("Main thread waiting for all tasks to complete...")
	wg.Wait() // Main thread waits for all goroutines to signal completion

	fmt.Printf("All tasks completed. Aggregated data: %v\n", data)
}