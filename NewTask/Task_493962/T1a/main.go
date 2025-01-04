package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to simulate validating a blockchain transaction
func validateTransaction(tx int, wg *sync.WaitGroup) {
	// Simulate some time-consuming transaction validation logic
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if tx%2 == 0 {
		fmt.Printf("Transaction %d is valid.\n", tx)
	} else {
		fmt.Printf("Transaction %d is invalid.\n", tx)
	}
	wg.Done()
}

func main() {
	// Initialize a WaitGroup
	var wg sync.WaitGroup
	// Number of transactions to validate
	numTransactions := 10

	// Start validating transactions concurrently
	for i := 0; i < numTransactions; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go validateTransaction(i, &wg)
	}

	// Wait for all transactions to be validated
	wg.Wait()

	// Proceed with further operations once all validations are complete
	fmt.Println("All transactions have been validated.")
}