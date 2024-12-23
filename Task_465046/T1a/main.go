package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to make an API call
func makeAPICall(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter

	// Simulate API request
	fmt.Println("Making request to", url)

	// You would normally make an HTTP request here
	// For demonstration, we will sleep for a random time
	select {
	case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
		fmt.Println("Request to", url, "completed")
	default:
	}
}

func main() {
	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Define some URLs for API requests
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
		"https://api.example.com/data4",
		"https://api.example.com/data5",
	}

	// Create a goroutine for each URL
	for _, url := range urls {
		wg.Add(1) // Increment the WaitGroup counter
		go makeAPICall(url, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All API requests completed.")
}
