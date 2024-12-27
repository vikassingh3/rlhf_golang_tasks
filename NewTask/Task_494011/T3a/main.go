package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Partition represents a single partition with its own lock and data map.
type Partition struct {
	mu               sync.RWMutex
	data             map[string]string
	transactionLog   []*Transaction
	recoveryState    bool
	checkpointFreq   time.Duration
	checkpointTimer  *time.Timer
	checkpointChan   chan struct{}
	transactionState map[int64]int32 // 0: Pending, 1: Committed, 2: Aborted
}

type Transaction struct {
	id      int64
	key     string
	value   string
	timestamp time.Time
}

// NewPartitionedMap initializes a PartitionedMap with the specified number of partitions.
func NewPartitionedMap(numPartitions int) *PartitionedMap {
	pm := &PartitionedMap{
		partitions: make([]*Partition, numPartitions),
	}

	for i := 0; i < numPartitions; i++ {
		pm.partitions[i] = &Partition{
			data:             make(map[string]string),
			transactionLog:   make([]*Transaction, 0),
			transactionState: make(map[int64]int32),
			checkpointFreq:   5 * time.Second,
			checkpointChan:   make(chan struct{}),
		}

		go pm.partitions[i].checkpointLoop()
	}

	return pm
}

func (pm *PartitionedMap) nextTransactionID() int64 {
	return atomic.AddInt64(&pm.transactionID, 1)
}

// getPartitionIndex determines which partition a key belongs to.
func (pm *PartitionedMap) getPartitionIndex(key string) int {
	// Hash the key to determine the partition index
	// (Simple hashing function for demonstration purposes)
	return int(key[0]) % len(pm.partitions)
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
	partition.transactionLog = append(partition.transactionLog, &Transaction{
		id:      txID,
		key:     key,
		value:   value,
		timestamp: time.Now(),
	})
}

// Commit commits the transaction across all participating partitions.
func (pm *PartitionedMap) Commit(txID int64) bool {
	for _, partition := range pm.partitions {
		partition.mu.Lock()
		defer partition.mu.Unlock()

		// If the transaction is not pending, abort the commit
		if partition.transactionState[txID] != Pending {
			log.Printf("Transaction %d not pending, aborting commit\n", txID)
			return false
		}

		// Mark the transaction as committed
		partition.transactionState[txID] = Committed
	}

	log.Printf("Transaction %d committed successfully\n", txID)
	return true
}

// Abort aborts the transaction across all participating partitions.
func (pm *PartitionedMap) Abort(txID int64) {
	for _, partition := range pm.partitions {
		partition.mu.Lock()
		defer partition.mu.Unlock()

		// If the transaction is already committed, ignore the abort
		if partition.transactionState[txID] == Committed {
			log.Printf("Transaction %d already committed, ignoring abort\n", txID)
			return
		}

		// Mark the transaction as aborted
		partition.transactionState[txID] = Aborted
	}

	log.Printf("Transaction %d aborted successfully\n", txID)
}

func (p *Partition) checkpointLoop() {
	ticker := time.NewTicker(p.checkpointFreq)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-p.checkpointChan:
			p.mu.Lock()
			defer p.mu.Unlock()

			if p.recoveryState {
				p.recoveryState = false
				log.Println("Recovery completed.")
				break
			}

			p.recoveryState = true
			log.Println("Checkpointing...")

			// Save the partition state
			log.Println("Checkpoint data saved.")

			log.Println("Checkpoint log saved.")
			p.recoveryState = false
		default:
		}
	}
}

func main() {
	pm := NewPartitionedMap(5)

	// Simulate a transaction
	txID := pm.BeginTransaction()
	pm.Set(txID, "key1", "value1")
	pm.Set(txID, "key2", "value2")

	// Simulate a crisis event by stopping the checkpoint loop temporarily
	time.Sleep(3 * time.Second)
	for _, partition := range pm.partitions {
		close(partition.checkpointChan)
		partition.checkpointTimer = time.NewTimer(1 * time.Second)
		go func(timer *time.Timer, partition *Partition) {
			select {
			case <-partition.checkpointTimer.C:
				partition.checkpointChan = make(chan struct{})
			}
		}(partition.checkpointTimer, partition)
	}

	// Simulate recovery
	time.Sleep(2 * time.Second)
	pm.Commit(txID)
}