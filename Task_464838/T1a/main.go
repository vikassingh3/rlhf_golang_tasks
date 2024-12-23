package main

import (
	"fmt"
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

// Function to validate a transaction
func ValidateTransaction(from string, to string, amount float64, balance float64) error {
	if amount <= 0 {
		return &ValidationError{Message: "Invalid amount: amount must be greater than zero"}
	}

	if amount > balance {
		return &InsufficientFundsError{
			From:    from,
			To:      to,
			Amount:  amount,
			Message: "Insufficient funds",
		}
	}

	return nil
}

// Function to process a transaction
func ProcessTransaction(from string, to string, amount float64, balance float64) {
	err := ValidateTransaction(from, to, amount, balance)
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction successful.")
	}
}

func main() {
	balance := 100.0

	ProcessTransaction("Alice", "Bob", 50.0, balance) // Successful transaction
	ProcessTransaction("Alice", "Bob", 150.0, balance) // Insufficient funds
	ProcessTransaction("Alice", "Bob", 0.0, balance)   // Invalid amount
}