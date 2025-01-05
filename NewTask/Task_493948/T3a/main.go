package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Mock function simulating serverless data aggregation
func aggregateData(ctx context.Context, data []int, wg *sync.WaitGroup, result chan<- int, errorCh chan<- error) {
	defer wg.Done()
	if ctx.Err() != nil {
		errorCh <- ctx.Err()
		return
	}

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
	chunkSize := (dataSize + numWorkers - 1) / numWorkers
	var wg sync.WaitGroup
	result := make(chan int, numWorkers)
	errorCh := make(chan error)

	// Divide data into chunks and process concurrently
	for i := 0; i < dataSize; i += chunkSize {
		end := i + chunkSize
		if end > dataSize {
			end = dataSize
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3) // Add timeout for resource management
		defer cancel()

		wg.Add(1)
		go aggregateData(ctx, data[i:end], &wg, result, errorCh)
	}

	wg.Wait()
	close(result)
	close(errorCh)

	// Handle errors
	for err := range errorCh {
		if err != nil {
			fmt.Printf("Encountered an error: %v\n", err)
		}
	}

	// Combine results
	finalTotal := 0
	for partialTotal := range result {
		finalTotal += partialTotal
	}

	fmt.Printf("Final Aggregated Total: %d\n", finalTotal)
}