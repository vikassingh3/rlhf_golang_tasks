package main

import (
	"fmt"
	"sync"
	"time"
)

func processRequest(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work being done by the request
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Request %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	numRequests := 5

	// Add the number of concurrent requests to the wait group
	wg.Add(numRequests)

	// Start processing the requests concurrently
	for i := 1; i <= numRequests; i++ {
		go processRequest(i, &wg)
	}

	// Wait for all the concurrent requests to complete
	wg.Wait()

	fmt.Println("All requests completed.")
}