package main

import (
	"sync"
)

type Partition struct {
	data map[int]string
	mu   sync.Mutex
}

func (p *Partition) Set(key int, value string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data[key] = value
}

func (p *Partition) Get(key int) (string, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	val, ok := p.data[key]
	return val, ok
}

type PartitionedMap struct {
	partitions    []*Partition
	numPartitions int
}

func NewPartitionedMap(numPartitions int) *PartitionedMap {
	pm := &PartitionedMap{
		partitions:    make([]*Partition, numPartitions),
		numPartitions: numPartitions,
	}

	// Initialize each partition
	for i := 0; i < numPartitions; i++ {
		pm.partitions[i] = &Partition{
			data: make(map[int]string),
		}
	}

	return pm
}

func (pm *PartitionedMap) Set(key int, value string) {
	idx := key % pm.numPartitions
	pm.partitions[idx].Set(key, value)
}

func (pm *PartitionedMap) Get(key int) (string, bool) {
	idx := key % pm.numPartitions
	return pm.partitions[idx].Get(key)
}

func main() {
	pm := NewPartitionedMap(10) // Create a partitioned map with 10 partitions

	// Set some key-value pairs
	pm.Set(1, "one")
	pm.Set(2, "two")
	pm.Set(15, "fifteen")

	// Get values
	val, ok := pm.Get(1)
	if ok {
		println("Got key 1:", val) // Output: Got key 1: one
	} else {
		println("Key 1 not found")
	}

	val, ok = pm.Get(20)
	if ok {
		println("Got key 20:", val)
	} else {
		println("Key 20 not found") // Output: Key 20 not found
	}
}
