package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Block struct represents a block in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Transactions []string
	PrevHash     string
	Hash         string
}

// Blockchain struct to encapsulate thread safety
type Blockchain struct {
	blocks []Block
	mutex  sync.Mutex
}

func (bc *Blockchain) AddBlock(transactions []string) error {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	latestBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := Block{
		Index:        latestBlock.Index + 1,
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Transactions: transactions,
		PrevHash:     latestBlock.Hash,
		Hash:         fmt.Sprintf("%d", rand.Int()), // Simulated hash
	}
	bc.blocks = append(bc.blocks, newBlock)

	if newBlock.Index == 5 {
		return errors.New("simulated error at block 5")
	}

	return nil
}

func main() {
	// Initialize blockchain with a genesis block
	bc := &Blockchain{
		blocks: []Block{
			{Index: 0, Timestamp: time.Now().Format("2006-01-02 15:04:05"), Transactions: []string{"Genesis transaction"}, PrevHash: "", Hash: "0"},
		},
	}

	// Simulate concurrent block addition
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errc := make(chan error, 10) // Buffered channel to hold errors
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := bc.AddBlock([]string{fmt.Sprintf("Transaction %d", i)})
			if err != nil {
				select {
				case <-ctx.Done():
					// Context cancelled, ignore error
					return
				case errc <- err:
					// Send error to error channel
					cancel() // Cancel context on error
				}
			}
		}(i)
	}

	wg.Wait()

	// Check for any errors that occurred during processing
	select {
	case err := <-errc:
		fmt.Println("Error:", err)
	default:
		fmt.Println("All transactions successfully added to the blockchain.")
	}

	// Print the blockchain
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	for _, block := range bc.blocks {
		fmt.Printf("Block %d: %+v\n", block.Index, block)
	}
}
 