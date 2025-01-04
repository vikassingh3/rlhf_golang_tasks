package main

import (
	"fmt"
	"sync"
)

type BlockchainLedger struct {
	mutex  sync.RWMutex
	blocks map[int]string // map of block indices to block content
}

func NewBlockchainLedger() *BlockchainLedger {
	return &BlockchainLedger{
		blocks: make(map[int]string),
	}
}

func (ledger *BlockchainLedger) AddBlock(blockIndex int, blockContent string) {
	ledger.mutex.Lock()
	defer ledger.mutex.Unlock()

	ledger.blocks[blockIndex] = blockContent
}

func (ledger *BlockchainLedger) GetBlock(blockIndex int) (string, bool) {
	ledger.mutex.RLock()
	defer ledger.mutex.RUnlock()

	blockContent, exists := ledger.blocks[blockIndex]
	return blockContent, exists
}

func main() {
	ledger := NewBlockchainLedger()

	// Simulate adding blocks concurrently
	var wg sync.WaitGroup
	numBlocks := 10000

	for i := 0; i < numBlocks; i++ {
		wg.Add(1)
		go func(blockIndex int) {
			blockContent := fmt.Sprintf("Block %d content", blockIndex)
			ledger.AddBlock(blockIndex, blockContent)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Simulate fetching blocks concurrently
	wg.Add(numBlocks)
	for i := 0; i < numBlocks; i++ {
		go func(blockIndex int) {
			blockContent, exists := ledger.GetBlock(blockIndex)
			if exists {
				fmt.Printf("Fetched block %d: %s\n", blockIndex, blockContent)
			} else {
				fmt.Printf("Block %d not found\n", blockIndex)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("All blocks added and fetched.")
}
