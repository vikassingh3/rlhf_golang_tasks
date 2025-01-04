package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate fetching data from the URL
	time.Sleep(2 * time.Second)
	fmt.Println("Fetched:", url)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://example.com/1",
		"https://example.com/2",
		"https://example.com/3",
		"https://example.com/4",
		"https://example.com/5",
	}

	// Add the number of URLs to the wait group
	wg.Add(len(urls))

	for _, url := range urls {
		go fetchURL(url, &wg)
	}

	// Wait for all fetches to complete
	wg.Wait()

	fmt.Println("All data fetched. Processing...")
	// Process the fetched data
}