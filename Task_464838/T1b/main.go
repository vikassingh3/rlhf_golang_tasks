package main

import (
	"errors"
	"fmt"
)

// Define your custom error types to represent different types of transaction validation errors
type InsufficientFundsError struct {
	AvailableBalance uint64
	RequiredBalance   uint64
}

type InvalidAddressError struct {
	Address string
}

// Implement the 'Error()' method to provide a user-friendly error message for each custom error
func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds. Available: %d, Required: %d", e.AvailableBalance, e.RequiredBalance)
}

func (e InvalidAddressError) Error() string {
	return fmt.Sprintf("invalid address: %s", e.Address)
}

// Transaction struct representing a transaction in the blockchain
type Transaction struct {
	Sender    string
	Recipient string
	Amount    uint64
}

func validateTransaction(tx Transaction, balance uint64) error {
	// Simplified validation logic
	if tx.Amount == 0 {
		return errors.New("transaction amount cannot be zero")
	}

	if balance < tx.Amount {
		return InsufficientFundsError{
			AvailableBalance: balance,
			RequiredBalance:   tx.Amount,
		}
	}

	if tx.Sender == "" || tx.Recipient == "" {
		return InvalidAddressError{
			Address: "either sender or recipient is empty",
		}
	}

	return nil
}

func main() {
	tx1 := Transaction{
		Sender:    "alice",
		Recipient: "bob",
		Amount:    100,
	}
	currentBalance := uint64(50)

	err := validateTransaction(tx1, currentBalance)
	if err != nil {
		fmt.Println("Transaction validation failed:", err)

		// Use type assertion to handle specific custom errors
		if ife, ok := err.(InvalidAddressError); ok {
			fmt.Println("Invalid Address:", ife.Address)
		} else if ife, ok := err.(InsufficientFundsError); ok {
			fmt.Println("Insufficient Funds:", ife)
		}

		return
	}

	fmt.Println("Transaction validated successfully!")
}