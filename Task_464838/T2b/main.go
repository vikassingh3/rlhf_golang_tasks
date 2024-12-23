package main

import (
	"crypto/sha256"
	"fmt"
	"time"
	"errors"
)

// Define custom error types
type InsufficientFundsError struct {
	From   string
	To     string
	Amount float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: from %s to %s amount %.2f", e.From, e.To, e.Amount)
}

// Block struct
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Hash         string
	PrevHash     string
}

// Transaction struct
type Transaction struct {
	From    string
	To      string
	Amount  float64
}

// Blockchain slice
type Blockchain []Block

// ValidateTransaction function
func ValidateTransaction(from, to string, amount, balance float64) error {
	if amount > balance {
		return &InsufficientFundsError{
			From:   from,
			To:     to,
			Amount: amount,
		}
	}
	if amount <= 0 {
		return errors.New("transaction amount must be greater than zero")
	}
	return nil
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(transactions []Transaction) error {
	// Validate transactions before adding them to the block
	for _, tx := range transactions {
		if err := ValidateTransaction(tx.From, tx.To, tx.Amount, 100.0); err != nil { // Assuming 100.0 as balance for simplicity
			return err
		}
	}

	// Create a new block
	prevBlock := (*bc)[len(*bc)-1]
	newBlock := Block{
		Index:        len(*bc),
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Hash:         "",
	}

	// Calculate the hash for the new block
	newBlock.Hash = calculateHash(newBlock)

	// Add the block to the blockchain
	*bc = append(*bc, newBlock)
	return nil
}

// calculateHash generates a SHA-256 hash for a block
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Transactions, block.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return fmt.Sprintf("%x", hash)
}

// NewBlockchain function to initialize the blockchain
func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: nil,
		Hash:         "GenesisHash",
		PrevHash:     "",
	}
	return &Blockchain{genesisBlock}
}

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Example transactions
	tx1 := Transaction{From: "Alice", To: "Bob", Amount: 50.0}
	tx2 := Transaction{From: "Bob", To: "Charlie", Amount: 25.0}
	tx3 := Transaction{From: "Alice", To: "Bob", Amount: 150.0} // Invalid amount

	// Validate and add transactions to the blockchain
	if err := blockchain.AddBlock([]Transaction{tx1, tx2}); err != nil {
		fmt.Println("Error adding block:", err)
	}

	if err := blockchain.AddBlock([]Transaction{tx3}); err != nil {
		// Handle the InsufficientFundsError
		if ife, ok := err.(*InsufficientFundsError); ok {
			fmt.Println("Transaction failed due to insufficient funds:")
			fmt.Printf("From: %s, To: %s, Amount: %.2f\n", ife.From, ife.To, ife.Amount)
		} else {
			fmt.Println("Error adding block:", err)
		}
	}

	// Print the blockchain
	for _, block := range *blockchain {
		fmt.Printf("Index: %d, Timestamp: %s, Hash: %s, PrevHash: %s\n", block.Index, block.Timestamp, block.Hash, block.PrevHash)
	}
}
