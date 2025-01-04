package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	workDuration := time.Duration(rand.Intn(5)+1) * time.Second
	fmt.Printf("Worker %d started working for %s\n", id, workDuration)
	time.Sleep(workDuration)
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	var wg sync.WaitGroup
	numWorkers := 5

	// Launch initial workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, &wg)
	}

	// Wait for initial workers to complete
	wg.Wait()
	fmt.Println("All initial workers completed.")

	// Add more workers dynamically
	additionalWorkers := 3
	for i := numWorkers; i < numWorkers+additionalWorkers; i++ {
		wg.Add(1)
		go worker(i+1, &wg)
	}

	// Wait for dynamically added workers to complete
	wg.Wait()
	fmt.Println("All workers (including dynamic ones) completed.")
}
