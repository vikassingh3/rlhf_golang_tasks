package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Mock data source for demonstration
func fetchData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(100)
	}
	return data
}

func aggregateDataSequential(data []int) int {
	total := 0
	for _, value := range data {
		total += value
	}
	return total
}

func aggregateDataParallel(data []int, numWorkers int) int {
	var wg sync.WaitGroup
	chunkSize := len(data) / numWorkers
	result := make(chan int, numWorkers) // Buffered to avoid blocking

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			partialTotal := aggregateDataSequential(chunk)
			result <- partialTotal
		}(data[i:end])
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(result)

	// Aggregate results from all workers
	total := 0
	for partialTotal := range result {
		total += partialTotal
	}

	return total
}

func main() {
	dataSize := 1000000 // Adjust this to test different data sizes
	numWorkers := 4     // Number of workers for parallel processing

	// Sequential Approach
	data := fetchData(dataSize)
	start := time.Now()
	sequentialResult := aggregateDataSequential(data)
	sequentialDuration := time.Since(start)

	// Parallel Approach
	start = time.Now()
	parallelResult := aggregateDataParallel(data, numWorkers)
	parallelDuration := time.Since(start)

	// Validate the results
	if sequentialResult != parallelResult {
		fmt.Println("Error: Results are different!")
		return
	}

	fmt.Printf("Data Size: %d\n", dataSize)
	fmt.Printf("Sequential Aggregation Time: %v\n", sequentialDuration)
	fmt.Printf("Parallel Aggregation Time: %v\n", parallelDuration)
	fmt.Printf("Speedup: %.2fx\n", sequentialDuration.Seconds()/parallelDuration.Seconds())
}
