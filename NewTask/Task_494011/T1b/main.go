package main

import (
	"fmt"
	"sync"
)

type partitionedMap struct {
	partitions []map[string]string
	lock       sync.RWMutex
}

func newPartitionedMap(numPartitions int) *partitionedMap {
	pm := &partitionedMap{
		partitions: make([]map[string]string, numPartitions),
	}
	for i := range pm.partitions {
		pm.partitions[i] = make(map[string]string)
	}
	return pm
}

func (pm *partitionedMap) partitionIndex(key string) int {
	// Simple hash function using modulo
	return int(hash(key)) & (len(pm.partitions) - 1)
}

func (pm *partitionedMap) Get(key string) (string, bool) {
	pm.lock.RLock()
	defer pm.lock.RUnlock()
	index := pm.partitionIndex(key)
	value, ok := pm.partitions[index][key]
	return value, ok
}

func (pm *partitionedMap) Set(key, value string) {
	pm.lock.Lock()
	defer pm.lock.Unlock()
	index := pm.partitionIndex(key)
	pm.partitions[index][key] = value
}

func (pm *partitionedMap) Delete(key string) {
	pm.lock.Lock()
	defer pm.lock.Unlock()
	index := pm.partitionIndex(key)
	delete(pm.partitions[index], key)
}

func hash(s string) uint32 {
	// Simple hashing function for demonstration purposes
	var h uint32
	for _, b := range s {
		h = 31*h + uint32(b)
	}
	return h
}

func main() {
	pm := newPartitionedMap(10) // Create a partitioned map with 10 partitions
	var wg sync.WaitGroup

	// Increment WaitGroup counter for each goroutine
	wg.Add(2)

	go func() {
		defer wg.Done() // Decrement counter when done
		for i := 0; i < 1000; i++ {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			pm.Set(key, value)
		}
	}()

	go func() {
		defer wg.Done() // Decrement counter when done
		for i := 0; i < 1000; i++ {
			key := fmt.Sprintf("key%d", i)
			value, _ := pm.Get(key)
			fmt.Println(key, value)
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()
}
