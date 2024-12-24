package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processEvent(event int, done chan<- bool) {
	// Simulate some processing work using a more efficient algorithm
	// For example, using math operations instead of a loop
	for i := 0; i < event*1000; i++ {
	}

	fmt.Println("Processed event:", event)
	done <- true // Signal completion
}

func main() {
	rand.Seed(time.Now().UnixNano())
	events := make([]int, 1000000)
	for i := range events {
		events[i] = rand.Intn(1000)
	}

	// Create a buffered channel to allow worker goroutines to complete before blocking
	done := make(chan bool, 10000)

	// Start worker goroutines
	numWorkers := 4 // You can adjust the number of workers based on your system's capabilities
	for i := 0; i < numWorkers; i++ {
		go func() {
			for event := range events {
				processEvent(event, done)
			}
		}()
	}

	// Wait for all workers to finish processing all events
	for range events {
		<-done
	}

	fmt.Println("All events processed.")
}
                    