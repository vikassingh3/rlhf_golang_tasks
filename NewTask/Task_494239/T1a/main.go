package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a slice of numbers
	numbers := []int{1, 2, 3, 4, 5}

	// Create a WaitGroup to wait for all goroutines to finish
	wg := &sync.WaitGroup{}

	// Increment the WaitGroup counter for each goroutine
	wg.Add(2)

	// Process the entire slice
	go func() {
		processSlice(numbers[:])
		wg.Done() // Decrement the counter when this goroutine finishes
	}()

	// Process elements of the slice using a channel
	go func() {
		processElements(numbers)
		wg.Done() // Decrement the counter when this goroutine finishes
	}()

	// Wait for all goroutines to complete
	wg.Wait()
}

// Process the entire slice
func processSlice(numbers []int) {
	for _, num := range numbers {
		fmt.Println("Processing slice:", num)
	}
}

// Process elements of the slice using a channel
func processElements(numbers []int) {
	elementChannel := make(chan int, len(numbers)) // Buffered channel to avoid deadlock

	// Send each element of the slice to the channel
	go func() {
		for _, num := range numbers {
			elementChannel <- num
		}
		close(elementChannel) // Close the channel when done sending
	}()

	// Process each element from the channel
	for element := range elementChannel {
		fmt.Println("Processing element:", element)
	}
}
