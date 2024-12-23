package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Define custom error types
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type InsufficientFundsError struct {
	From    string
	To      string
	Amount  float64
	Message string
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, fmt.Sprintf("From: %s, To: %s, Amount: %.2f", e.From, e.To, e.Amount))
}

type DoubleSpendError struct {
	TransactionID string
	Message        string
}

func (e *DoubleSpendError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.TransactionID)
}

// Define a simple account structure
type Account struct {
	ID    string
	Balance float64
}

// Define a transaction structure
type Transaction struct {
	ID         string
	From       string
	To         string
	Amount     float64
	Timestamp  time.Time
	Signature  []byte // For simplicity, we'll just use a random byte slice
}

// Simulate a blockchain transaction pool
var transactionPool = []*Transaction{}

// Simulate a list of accounts
var accounts = []*Account{
	&Account{ID: "Alice", Balance: 100.0},
	&Account{ID: "Bob", Balance: 50.0},
}

// Function to generate a random transaction ID
func generateTransactionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Function to validate a transaction
func ValidateTransaction(tx *Transaction) error {
	for _, t := range transactionPool {
		if t.From == tx.From && t.Amount == tx.Amount {
			return &DoubleSpendError{TransactionID: tx.ID, Message: "Double spend detected"}
		}
	}

	for _, account := range accounts {
		if account.ID == tx.From {
			if tx.Amount <= 0 {
				return &ValidationError{Message: "Invalid amount: amount must be greater than zero"}
			}

			if tx.Amount > account.Balance {
				return &InsufficientFundsError{
					From:    tx.From,
					To:      tx.To,
					Amount:  tx.Amount,
					Message: "Insufficient funds",
				}
			}

			break
		}
	}

	return nil
}

// Function to process a transaction
func ProcessTransaction(tx *Transaction) {
	log.Println("Attempting to process transaction:", tx.ID)
	err := ValidateTransaction(tx)
	if err != nil {
		log.Printf("Transaction failed: %v", err)
	} else {
		log.Println("Transaction successful:", tx.ID)
		transactionPool = append(transactionPool, tx)

		// For simplicity, let's assume the transaction updates the accounts immediately
		for _, account := range accounts {
			if account.ID == tx.From {
				account.Balance -= tx.Amount
			}
			if account.ID == tx.To {
				account.Balance += tx.Amount
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a new transaction
	tx := &Transaction{
		ID:         generateTransactionID(),
		From:       "Alice",
		To:         "Bob",
		Amount:     50.0,
		Timestamp:  time.Now(),
		Signature:  make([]byte, 32), // Random signature
	}

	// Attempt to process the transaction
	ProcessTransaction(tx)

	// Create a double-spend transaction
	doubleSpendTx := &Transaction{
		ID:         generateTransactionID(),
		From:       "Alice",
		To:         "Charlie",
		Amount:     50.0,
		Timestamp:  time.Now(),
		Signature:  make([]byte, 32), // Random signature
	}

	// Attempt to process the double-spend transaction
	ProcessTransaction(doubleSpendTx)
}