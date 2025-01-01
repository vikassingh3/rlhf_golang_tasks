package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	cache sync.Map
	mu    sync.Mutex // Mutex to avoid race conditions during Load/Store operations
)

// fetchData retrieves data for a given key, using the cache if possible.
func fetchData(key string) (string, bool) {
	// Check if the key exists in the cache
	if value, ok := cache.Load(key); ok {
		return value.(string), true // Return cached value
	}

	// Synchronize data fetching to avoid duplicate calls for the same key
	mu.Lock()
	defer mu.Unlock()

	// Re-check the cache after acquiring the lock
	if value, ok := cache.Load(key); ok {
		return value.(string), true
	}

	// Simulate database query or external service call
	fmt.Printf("Fetching data from DB/API for key: %s\n", key)
	time.Sleep(time.Second)

	// Store the result in the cache
	result := fmt.Sprintf("Value for %s", key)
	cache.Store(key, result)

	return result, true
}

func main() {
	keys := []string{"key1", "key2", "key3", "key1", "key2"}

	var wg sync.WaitGroup
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			data, _ := fetchData(k)
			fmt.Println("Data for", k, ":", data)
		}(key)
	}
	wg.Wait()
}
