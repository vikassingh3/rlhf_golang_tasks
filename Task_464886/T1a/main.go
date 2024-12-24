package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processEvent(event int) {
	// Simulate some processing work
	for i := 0; i < 10; i++ {
	}
	fmt.Println("Processed event:", event)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	events := make([]int, 10) // Create a slice of 1 million events
	for i := range events {
		events[i] = rand.Intn(1000) // Random event numbers
	}

	// Start a range-based event loop
	for _, event := range events {
		processEvent(event)
	}
}