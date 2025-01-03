package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add 3 goroutines to the WaitGroup.
	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	// Wait for all goroutines to finish.
	wg.Wait()

	fmt.Println("All workers completed!")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine exits.

	fmt.Printf("Worker %d is starting...\n", id)

	// Simulate work by sleeping for a random duration.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	fmt.Printf("Worker %d has completed.\n", id)
}