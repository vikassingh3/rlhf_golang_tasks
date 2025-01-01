package main

import (
	"fmt"
	"sync"
	"time"
)
var (
	// Cache data
	cache sync.Map
)

// Retrieve data from a remote service (Simulate using an API call)
func retrieveDataFromRemoteService(key string) string {
	time.Sleep(200 * time.Millisecond) // Simulate processing time
	return fmt.Sprintf("Value for Key %s", key)
}
//Get data from cache with fallback to retrieve it from remote service
func getDataFromCacheOrService(key string) string {
	// Check if data is already in the cache
	if val, ok := cache.Load(key); ok {
		return val.(string)
	}

	// If data is not in the cache, fetch it from the remote service
	val := retrieveDataFromRemoteService(key)

	// Store the retrieved data in the cache
	cache.Store(key, val)
	return val
}

func main() {
	keys := []string{"key1", "key2", "key3", "key1", "key4", "key3"}

	// Simulate real-time requests
	for _, key := range keys {
		start := time.Now()
		data := getDataFromCacheOrService(key)
		fmt.Println(data)
		end := time.Now()
		latency := end.Sub(start)
		fmt.Printf("Retrieved data for key: %s in %s \n", key, latency)
	}
}