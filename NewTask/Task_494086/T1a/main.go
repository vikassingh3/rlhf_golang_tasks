package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate processing a request
func processRequest(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work (e.g., database query, API call)
	fmt.Printf("Processing request ID: %d...\n", id)
	// Sleep for a random amount of time
	// In a real application, this would be replaced with actual work
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("Request ID: %d processed.\n", id)
}

func main() {
	var wg sync.WaitGroup
	numRequests := 5

	// Start processing requests in separate goroutines
	for i := 1; i <= numRequests; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go processRequest(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All requests have been processed.")
}