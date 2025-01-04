package main

import (
	"fmt"
	"sync"
)

type Transaction struct {
	Amount      int
	Source      string
	Destination string
}

func validateTransaction(transaction Transaction, results chan<- bool) {
	// Simulate validation process.
	if transaction.Amount <= 0 {
		results <- false
		return
	}
	if transaction.Source == transaction.Destination {
		results <- false
		return
	}
	results <- true
}

func main() {
	transactions := []Transaction{
		{Amount: 100, Source: "alice", Destination: "bob"},
		{Amount: -50, Source: "bob", Destination: "alice"}, // Invalid amount
		{Amount: 75, Source: "charles", Destination: "bob"},
		{Amount: 25, Source: "bob", Destination: "charles"},
		{Amount: 0, Source: "dawn", Destination: "eve"}, // Invalid amount
	}

	var wg sync.WaitGroup
	results := make(chan bool, len(transactions)) // Buffered channel

	// Launch validation goroutines
	for _, transaction := range transactions {
		wg.Add(1)
		go func(transaction Transaction) {
			defer wg.Done()
			validateTransaction(transaction, results)
		}(transaction)
	}

	// Close the channel once all goroutines have finished sending results
	go func() {
		wg.Wait()       // Wait for all goroutines to finish
		close(results)  // Close the results channel
	}()

	// Process the results
	valid := true
	for result := range results {
		if !result {
			valid = false
			break
		}
	}

	// Output result
	if valid {
		fmt.Println("All transactions are valid.")
	} else {
		fmt.Println("At least one transaction is invalid.")
	}
}
