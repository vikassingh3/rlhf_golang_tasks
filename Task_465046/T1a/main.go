package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to make an API call
func makeAPICall(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when done

	// Simulate an API call with a delay
	fmt.Println("Making request to", url)
	time.Sleep(2 * time.Second) // Simulate network delay
	fmt.Println("Request to", url, "completed")
}

func main() {
	// List of URLs to simulate API requests
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
		"https://api.example.com/data4",
		"https://api.example.com/data5",
	}

	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Create goroutines for each API call
	for _, url := range urls {
		wg.Add(1) // Increment the counter
		go makeAPICall(url, &wg)
	}

	// Wait for all API calls to complete
	wg.Wait()
	fmt.Println("All API requests completed.")
}
