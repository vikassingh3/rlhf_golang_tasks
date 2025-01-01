package main

import (
	"fmt"
	"sync"
	"time"
)

func processDataConcurrently(data []int, results chan int) {
	var wg sync.WaitGroup

	for _, d := range data {
		wg.Add(1)
		go func(d int) {
			defer wg.Done() // Ensure Done is called even if there's an error
			time.Sleep(time.Millisecond * 50) // Simulate processing time
			results <- d * d
		}(d)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results) // Close channel after all data has been sent
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	results := make(chan int)

	go processDataConcurrently(data, results)

	// Read from the results channel
	for result := range results {
		fmt.Println(result)
	}
}
