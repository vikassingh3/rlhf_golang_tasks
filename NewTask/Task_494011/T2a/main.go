package main

import (
	"fmt"
	"sync"
)

// Partition represents a single partition with its own lock and data map.
type Partition struct {
	mu   sync.RWMutex  // Change to RWMutex
	data map[string]string
}

// PartitionedMap manages multiple partitions for improved concurrency.
type PartitionedMap struct {
	partitions []*Partition
	numPartitions int
}

// NewPartitionedMap initializes a PartitionedMap with the specified number of partitions.
func NewPartitionedMap(numPartitions int) *PartitionedMap {
	partitions := make([]*Partition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		partitions[i] = &Partition{
			data: make(map[string]string),
		}
	}
	return &PartitionedMap{
		partitions: partitions,
		numPartitions: numPartitions,
	}
}

// getPartitionIndex determines which partition a key belongs to.
func (pm *PartitionedMap) getPartitionIndex(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % pm.numPartitions
}

// setWithLocks sets a key-value pair while acquiring locks on affected partitions.
func (pm *PartitionedMap) setWithLocks(key, value string) {
	partitionIndex := pm.getPartitionIndex(key)
	partition := pm.partitions[partitionIndex]

	partition.mu.Lock()
	defer partition.mu.Unlock()

	partition.data[key] = value
}

// getWithLocks retrieves a value for a given key while acquiring locks on affected partitions.
func (pm *PartitionedMap) getWithLocks(key string) (string, bool) {
	partitionIndex := pm.getPartitionIndex(key)
	partition := pm.partitions[partitionIndex]

	partition.mu.RLock()  // Use RLock for read access
	defer partition.mu.RUnlock()  // Use RUnlock for read access

	value, exists := partition.data[key]
	return value, exists
}

func main() {
	pm := NewPartitionedMap(5)

	// Simulate writes with locking
	pm.setWithLocks("key1", "value1")
	pm.setWithLocks("key2", "value2")

	// Simulate reads with locking
	if val, ok := pm.getWithLocks("key1"); ok {
		fmt.Println("Got key1:", val)
	} else {
		fmt.Println("Key1 not found")
	}

	if val, ok := pm.getWithLocks("key4"); ok {
		fmt.Println("Got key4:", val)
	} else {
		fmt.Println("Key4 not found")
	}
}
