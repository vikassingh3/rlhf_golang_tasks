
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

	// Synchronize access to the map
	var mu sync.Mutex
	var wg sync.WaitGroup
	numTransactions := 100000

	// Timing for Insertions
	fmt.Println("Timing Insertions...")
	wg.Add(1)
	start := time.Now()
	go func() {
		defer wg.Done()
		for i := 0; i < numTransactions; i++ {
			mu.Lock()
			m[rand.Intn(size*10)] = rand.Int()
			mu.Unlock()
		}
	}()
	wg.Wait()
	elapsedInsertion := time.Since(start)
	fmt.Printf("Time taken for Insertions: %s\n", elapsedInsertion)

	// Timing for Deletions
	fmt.Println("Timing Deletions...")
	wg.Add(1)
	start = time.Now()
	go func() {
		defer wg.Done()
		for key := range m {
			mu.Lock()
			delete(m, key)
			mu.Unlock()
		}
	}()
	wg.Wait()
	elapsedDeletion := time.Since(start)
	fmt.Printf("Time taken for Deletions: %s\n", elapsedDeletion)

	// Additional tests for Updates and Lookups can be added as needed.
}



