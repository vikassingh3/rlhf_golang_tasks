package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type APICallResult struct {
	URL   string
	Error error
}

// Function to make an API call
func makeAPICall(url string, wg *sync.WaitGroup, resultChan chan<- APICallResult) {
	defer wg.Done()

	// Simulate an API call with a delay
	fmt.Println("Making request to", url)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Simulate random delay
	// fmt.Println("Request to", url, "completed")

	// Simulate an error for demonstration purposes
	if rand.Intn(2) == 0 {
		resultChan <- APICallResult{URL: url, Error: fmt.Errorf("random error for %s", url)}
	} else {
		resultChan <- APICallResult{URL: url}
	}
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
	resultChan := make(chan APICallResult)
	errorCount := 0

	// Create goroutines for each API call
	for _, url := range urls {
		wg.Add(1)
		go makeAPICall(url, &wg, resultChan)
	}

	// Go routine to collect results and handle errors
	go func() {
		defer close(resultChan)
		for result := range resultChan {
			if result.Error != nil {
				fmt.Printf("Error from %s: %v\n", result.URL, result.Error)
				errorCount++
			}
		}
	}()

	// Wait for all API calls to complete
	wg.Wait()
	fmt.Println("All API requests completed.")

	if errorCount > 0 {
		fmt.Printf("Found %d errors during API calls.\n", errorCount)
	} else {
		fmt.Println("All API calls succeeded.")
	}
} 