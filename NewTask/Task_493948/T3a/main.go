package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

// Mock data source for demonstration
func fetchData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(100)
	}
	return data
}

func aggregateData(ctx context.Context, data []int, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		// Canceled before completion
		return
	default:
		total := 0
		for _, value := range data {
			total += value
		}
		result <- total
	}
}

func main() {
	dataSize := 1000000
	numWorkers := 4

	// Fetch data
	data := fetchData(dataSize)

	// Create a context for managing task lifecycles
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	result := make(chan int, numWorkers)

	// Start worker tasks
	for i := 0; i < len(data); i += dataSize / numWorkers {
		end := i + dataSize/numWorkers
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go aggregateData(ctx, data[i:end], &wg, result)
	}

	// Wait for all tasks to complete
	wg.Wait()
	close(result)

	// Aggregate results from all workers
	total := 0
	for partialTotal := range result {
		total += partialTotal
	}

	fmt.Printf("Data Size: %d\n", dataSize)
	fmt.Printf("Total Aggregated Result: %d\n", total)
}
