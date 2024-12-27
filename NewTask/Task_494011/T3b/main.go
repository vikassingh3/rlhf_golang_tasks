package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// VersionedPartition represents a single partition with versioned data.
type VersionedPartition struct {
	mu   sync.RWMutex
	data map[string]*dataWithVersion
}

type dataWithVersion struct {
	value  string
	version int64
}

func (vp *VersionedPartition) set(key, value string) {
	vp.mu.Lock()
	defer vp.mu.Unlock()
	// If the key doesn't exist, initialize it with a version 0
	if _, exists := vp.data[key]; !exists {
		vp.data[key] = &dataWithVersion{value: value, version: 0}
	}
	// Increment the version for existing keys
	newVersion := atomic.AddInt64(&vp.data[key].version, 1)
	vp.data[key] = &dataWithVersion{value, newVersion}
}

func (vp *VersionedPartition) get(key string) (string, int64, bool) {
	vp.mu.RLock()
	defer vp.mu.RUnlock()
	dv, ok := vp.data[key]
	if !ok {
		return "", 0, false
	}
	return dv.value, dv.version, true
}

// PartitionedMapWithVersioning manages multiple partitions with versioning support.
type PartitionedMapWithVersioning struct {
	partitions    []*VersionedPartition
	numPartitions int
}

func NewPartitionedMapWithVersioning(numPartitions int) *PartitionedMapWithVersioning {
	pm := &PartitionedMapWithVersioning{
		partitions:   make([]*VersionedPartition, numPartitions),
		numPartitions: numPartitions,
	}
	for i := 0; i < numPartitions; i++ {
		pm.partitions[i] = &VersionedPartition{
			data: make(map[string]*dataWithVersion),
		}
	}
	return pm
}

func (pm *PartitionedMapWithVersioning) getPartitionIndex(key string) int {
	return int(key[0]) % pm.numPartitions
}

func (pm *PartitionedMapWithVersioning) Set(key, value string) {
	index := pm.getPartitionIndex(key)
	partition := pm.partitions[index]
	partition.set(key, value)
}

func main() {
	// Example usage
	pm := NewPartitionedMapWithVersioning(3)

	pm.Set("key1", "value1")
	pm.Set("key1", "value2")
	value, version, exists := pm.partitions[pm.getPartitionIndex("key1")].get("key1")
	if exists {
		fmt.Printf("Key: key1, Value: %s, Version: %d\n", value, version)
	}
}
