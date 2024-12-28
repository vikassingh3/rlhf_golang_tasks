package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Initialize a map with a small capacity to force frequent resizes
	const size = 1000
	m := make(map[int]int, size)

	// Measure the time taken for different operations
	var wg sync.WaitGroup
	numTransactions := 100000

	// Timing for Insertions
	fmt.Println("Timing Insertions...")
	wg.Add(1)
	start := time.Now()
	go func() {
		for i := 0; i < numTransactions; i++ {
			m[rand.Intn(size*10)] = rand.Int()
		}
		wg.Done()
	}()
	wg.Wait()
	elapsedInsertion := time.Since(start)

	// Timing for Deletions
	fmt.Println("Timing Deletions...")
	wg.Add(1)
	start = time.Now()
	go func() {
		for key := range m {
			delete(m, key)
		}
		wg.Done()
	}()
	wg.Wait()
	elapsedDeletion := time.Since(start)

	// Timing for Updates 
	fmt.Println("Timing Updates...")
	wg.Add(1)
	start = time.Now()
	go func() {
		for key := range m {
			m[key] = key * 2
		}
		wg.Done()
	}()
	wg.Wait()
	elapsedUpdate := time.Since(start)

	// Timing for Lookups
	fmt.Println("Timing Lookups...")
	wg.Add(1)
	start = time.Now()
	go func() {
		for key := range m {
			_ = m[key]
		}
		wg.Done()
	}()
	wg.Wait()

	// Print the results
	fmt.Printf("Time taken for Insertions: %s\n", elapsedInsertion)
	fmt.Printf("Time taken for Deletions: %s\n", elapsedDeletion)
	fmt.Printf("Time taken for Updates: %s\n", elapsedUpdate)
}
