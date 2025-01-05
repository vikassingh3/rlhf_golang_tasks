package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Mock function simulating serverless data aggregation
func aggregateData(data []int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	total := 0
	for _, value := range data {
		total += value
	}
	result <- total
}

func main() {
	dataSize := 100000
	data := make([]int, dataSize)
	for i := range data {
		data[i] = rand.Intn(100)
	}

	numWorkers := 4
	chunkSize := (dataSize + numWorkers - 1) / numWorkers // Ensures all data is processed
	var wg sync.WaitGroup
	result := make(chan int, numWorkers)

	// Divide data into chunks and process concurrently
	for i := 0; i < dataSize; i += chunkSize {
		end := i + chunkSize
		if end > dataSize {
			end = dataSize
		}
		wg.Add(1)
		go aggregateData(data[i:end], &wg, result)
	}

	wg.Wait()
	close(result)

	// Combine results
	finalTotal := 0
	for partialTotal := range result {
		finalTotal += partialTotal
	}

	fmt.Printf("Final Aggregated Total: %d\n", finalTotal)
}
