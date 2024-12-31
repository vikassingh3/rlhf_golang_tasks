package main

import (
	"fmt"
	"sync"
	"time"
)

var resultSet []int   // Shared variable
var mutex *sync.Mutex = &sync.Mutex{}

// Simulate analyzing a chunk of data and appending results to a shared set
func analyzeChunk(data []int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter

	// Simulate analysis work (could be more complex logic)
	fmt.Printf("Analyzing chunk with elements: %v\n", data)
	time.Sleep(time.Duration(len(data)) * time.Second)

	// Protect shared resource with a mutex
	mutex.Lock()
	defer mutex.Unlock()

	// Simulate some processing where we modify a shared variable
	for _, num := range data {
		resultSet = append(resultSet, num*2) // Modifying the shared variable
	}
}

func main() {
	// Create a large dataset
	dataset := []int{
		1, 2, 3, 4, 5,
		6, 7, 8, 9, 10,
		11, 12, 13, 14, 15,
		16, 17, 18, 19, 20,
	}

	// Split the dataset into smaller chunks
	chunkSize := 5
	numChunks := len(dataset) / chunkSize
	chunks := make([][]int, numChunks)
	for i := range chunks {
		start := i * chunkSize
		end := start + chunkSize
		chunks[i] = dataset[start:end]
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Launch a goroutine for each chunk
	for _, chunk := range chunks {
		wg.Add(1) // Increment the WaitGroup counter
		go analyzeChunk(chunk, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All chunks have been analyzed.")
	fmt.Printf("Processed results: %v\n", resultSet)
}