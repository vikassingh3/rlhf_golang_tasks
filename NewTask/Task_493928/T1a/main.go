package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Fetch data from the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data from %s: %v\n", url, err)
		return
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("Error closing response body for %s: %v\n", url, err)
		}
	}()

	// Read the response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading data from %s: %v\n", url, err)
		return
	}

	// Print a truncated version of the data
	fmt.Printf("Fetched data from %s: %s...\n", url, string(data[:100]))
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://example.com/data1",
		"https://example.com/data2",
		"https://example.com/data3",
	}

	// Limit the number of concurrent goroutines
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent)

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer func() { <-sem }()
			sem <- struct{}{}
			fetchData(u, &wg)
		}(url)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All data fetched.")
}
