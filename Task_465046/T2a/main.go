package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to make an API call
func makeAPICall(url string, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done() // Decrement the WaitGroup counter when done

	// Simulate an API call with a delay
	fmt.Println("Making request to", url)
	time.Sleep(2 * time.Second) // Simulate network delay

	// Introduce a random error for demonstration
	if time.Now().Second()%3 == 0 {
		errors <- fmt.Errorf("Error fetching data from %s", url)
	}

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

	// Channel to collect errors
	errors := make(chan error, len(urls))

	// Create goroutines for each API call
	for _, url := range urls {
		wg.Add(1) // Increment the counter
		go makeAPICall(url, &wg, errors)
	}

	// Wait for all API calls to complete
	wg.Wait()

	// Close the error channel
	close(errors)

	// Collect and print errors
	fmt.Println("Collecting errors:")
	for err := range errors {
		fmt.Println(err)
	}

	fmt.Println("All API requests completed.")
}