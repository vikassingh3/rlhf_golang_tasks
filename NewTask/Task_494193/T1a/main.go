package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateRandomNumbers(count int, rateLimit int) []int {
	var results []int
	var wg sync.WaitGroup
	wg.Add(count)

	// Create a channel to limit the rate of output
	output := make(chan int, rateLimit)

	for i := 0; i < count; i++ {
		go func(idx int) {
			// Simulate some work
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Generate a random number
			num := rand.Intn(100)
			output <- num

			// Decrement the wait group counter
			wg.Done()
		}(i)
	}

	// Collect results from the output channel
	go func() {
		for num := range output {
			results = append(results, num)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the output channel
	close(output)

	return results
}

func main() {
	// Generate 100 random numbers with a rate limit of 5 numbers per second
	numbers := generateRandomNumbers(100, 5)
	fmt.Println("Generated numbers:", numbers)
}