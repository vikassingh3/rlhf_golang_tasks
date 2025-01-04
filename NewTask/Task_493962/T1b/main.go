package main  
import (  
  "fmt"
  "sync"
)

type Transaction struct {  
  Amount int
  Source string
  Destination string
}  

func validateTransaction(transaction Transaction, wg *sync.WaitGroup) {  
  defer wg.Done()  

  // Simulate validation process.
  // In a real implementation, you would use cryptographic 
  // functions to verify the transaction's validity.

  if transaction.Amount <= 0 {  
    fmt.Printf("Invalid transaction: Amount must be greater than 0. Transaction: %+v\n", transaction)
    return
  }  
  if transaction.Source == transaction.Destination {  
    fmt.Printf("Invalid transaction: Source and destination cannot be the same. Transaction: %+v\n", transaction)  
    return
  }  

  fmt.Printf("Transaction is valid: %+v\n", transaction)
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
  wg.Add(len(transactions))  

  for _, transaction := range transactions {  
    go validateTransaction(transaction, &wg)
  }  

  // Wait for all transactions to be validated
  wg.Wait()

  fmt.Println("All transactions have been validated.")
} 