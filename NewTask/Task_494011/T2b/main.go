package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Partition represents a single partition with its own lock and data map.
type Partition struct {
	mu               sync.RWMutex
	data             map[string]string
	transactionState map[int64]int32 // 0: Pending, 1: Committed, 2: Aborted
}

const (
	Pending   = 0
	Committed = 1
	Aborted   = 2
)

// PartitionedMap manages multiple partitions with 2PC support.
type PartitionedMap struct {
	partitions    []*Partition
	numPartitions int
	transactionID int64 // Atomic counter for transaction IDs
}

// NewPartitionedMap initializes a PartitionedMap with the specified number of partitions.
func NewPartitionedMap(numPartitions int) *PartitionedMap {
	pm := &PartitionedMap{
		partitions:    make([]*Partition, numPartitions),
		numPartitions: numPartitions,
		transactionID: 0,
	}
	for i := 0; i < numPartitions; i++ {
		pm.partitions[i] = &Partition{
			data:             make(map[string]string),
			transactionState: make(map[int64]int32),
		}
	}
	return pm
}

// nextTransactionID generates a unique transaction ID.
func (pm *PartitionedMap) nextTransactionID() int64 {
	return atomic.AddInt64(&pm.transactionID, 1)
}

// getPartitionIndex determines which partition a key belongs to.
func (pm *PartitionedMap) getPartitionIndex(key string) int {
	// Hash the key to determine the partition index
	// (Simple hashing function for demonstration purposes)
	return int(key[0]) % pm.numPartitions
}

// BeginTransaction starts a new transaction.
func (pm *PartitionedMap) BeginTransaction() int64 {
	return pm.nextTransactionID()
}

// Set adds a key-value pair to the appropriate partition during a transaction.
func (pm *PartitionedMap) Set(txID int64, key, value string) {
	index := pm.getPartitionIndex(key)
	partition := pm.partitions[index]
	partition.mu.Lock()
	defer partition.mu.Unlock()

	// If the transaction is already aborted, ignore the set operation
	if partition.transactionState[txID] == Aborted {
		return
	}

	// Mark the transaction as pending
	partition.transactionState[txID] = Pending
	partition.data[key] = value
}

// Commit commits the transaction across all participating partitions.
func (pm *PartitionedMap) Commit(txID int64) bool {
	// Prepare phase: Check if all partitions can commit the transaction
	for _, partition := range pm.partitions {
		partition.mu.Lock()
		// If the transaction is not pending, abort the commit
		if partition.transactionState[txID] != Pending {
			// Abort phase: If any partition cannot commit, we abort the transaction
			partition.transactionState[txID] = Aborted
			partition.mu.Unlock()
			return false
		}
		// Unlock after checking the transaction state
		partition.mu.Unlock()
	}

	// Commit phase: Mark all partitions as committed
	for _, partition := range pm.partitions {
		partition.mu.Lock()
		partition.transactionState[txID] = Committed
		partition.mu.Unlock()
	}

	// Return true indicating successful commit
	return true
}

// Abort aborts the transaction across all participating partitions.
func (pm *PartitionedMap) Abort(txID int64) {
	// Abort phase: Mark all partitions as aborted
	for _, partition := range pm.partitions {
		partition.mu.Lock()
		partition.transactionState[txID] = Aborted
		partition.mu.Unlock()
	}
}

func main() {
	pm := NewPartitionedMap(5)

	// Start a transaction
	txID := pm.BeginTransaction()

	// Perform some operations
	pm.Set(txID, "key1", "value1")
	pm.Set(txID, "key2", "value2")

	// Commit the transaction
	if pm.Commit(txID) {
		fmt.Println("Transaction committed successfully")
	} else {
		fmt.Println("Transaction failed during commit")
	}

	// Try aborting a transaction
	txID2 := pm.BeginTransaction()
	pm.Set(txID2, "key3", "value3")
	pm.Abort(txID2)
}
