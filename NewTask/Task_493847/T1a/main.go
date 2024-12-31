package main

import (
	"fmt"
	"time"
)

// Transaction struct defines a transaction record
type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
}

// Transactions is a slice of Transaction records
type Transactions []Transaction

// AddTransaction appends a new transaction to the slice
func (ts *Transactions) AddTransaction(amount float64, category string) {
	date := time.Now()
	newTransaction := Transaction{
		ID:       len(*ts) + 1,  // Assign a new ID based on the current length of the slice
		Amount:   amount,
		Category: category,
		Date:     date,
	}
	*ts = append(*ts, newTransaction)
}

// UpdateTransaction updates an existing transaction by ID
func (ts *Transactions) UpdateTransaction(id int, amount float64, category string) {
	for i, t := range *ts {
		if t.ID == id {
			(*ts)[i].Amount = amount
			(*ts)[i].Category = category
			return
		}
	}
	fmt.Printf("Transaction with ID %d not found.\n", id)
}

// RetrieveTransactionByID retrieves a transaction by its ID
func (ts *Transactions) RetrieveTransactionByID(id int) (Transaction, bool) {
	for _, t := range *ts {
		if t.ID == id {
			return t, true
		}
	}
	return Transaction{}, false
}

// ListTransactions lists all transactions in the slice
func (ts *Transactions) ListTransactions() {
	if len(*ts) == 0 {
		fmt.Println("No transactions found.")
		return
	}
	for _, t := range *ts {
		fmt.Printf("ID: %d, Amount: $%.2f, Category: %s, Date: %s\n", t.ID, t.Amount, t.Category, t.Date.Format("2006-01-02"))
	}
}

func main() {
	transactions := Transactions{}

	// Add new transactions
	transactions.AddTransaction(100.0, "Food")
	transactions.AddTransaction(50.0, "Transport")

	// List all transactions
	fmt.Println("List of transactions:")
	transactions.ListTransactions()

	// Update a transaction
	transactions.UpdateTransaction(1, 150.0, "Groceries")

	// List all transactions again to show update
	fmt.Println("\nList of transactions after update:")
	transactions.ListTransactions()

	// Retrieve a transaction by ID
	retrievedTransaction, found := transactions.RetrieveTransactionByID(2)
	if found {
		fmt.Printf("\nRetrieved Transaction: ID: %d, Amount: $%.2f, Category: %s, Date: %s\n", retrievedTransaction.ID, retrievedTransaction.Amount, retrievedTransaction.Category, retrievedTransaction.Date.Format("2006-01-02"))
	} else {
		fmt.Println("\nTransaction not found.")
	}
}