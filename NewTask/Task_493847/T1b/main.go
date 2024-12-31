
package main  
import (  
    "fmt"
    "time"
)
type Transaction struct {  
    Date        time.Time
    Amount      float64
    Description string
}
// define a slice of transactions to store transactions
var transactions []Transaction
func addTransaction(amount float64, description string) {
    newTransaction := Transaction{
        Date:        time.Now(),
        Amount:      amount,
        Description: description,
    }
    transactions = append(transactions, newTransaction)
}
func updateTransaction(index int, amount float64, description string) {  
    if index >= 0 && index < len(transactions) {
        transactions[index].Amount = amount
        transactions[index].Description = description
    } else {
        fmt.Println("Invalid transaction index.")
    }
}
func retrieveTransaction(index int) Transaction {  
    if index >= 0 && index < len(transactions) {
        return transactions[index]
    }
    fmt.Println("Invalid transaction index.")
    return Transaction{} // return an empty transaction
}
func main() {
    // Add some transactions
    addTransaction(100.0, "Groceries")
    addTransaction(200.0, "Transportation")
    addTransaction(50.0, "Bills")
    
    // Display the current transactions
    fmt.Println("Current Transactions:")
    for i, t := range transactions {
        fmt.Printf("%d: %s (%.2f) - %s\n", i, t.Date.Format("2006-01-02"), t.Amount, t.Description)
    }

    // Update the second transaction
    updateTransaction(1, 150.0, "Updated Transportation")

    // Retrieve and display the third transaction
    t := retrieveTransaction(2)
    fmt.Printf("\nTransaction 3 details:\n%s (%.2f) - %s\n", t.Date.Format("2006-01-02"), t.Amount, t.Description)
    
    // Attempt to retrieve an invalid transaction
    retrieveTransaction(-1)  
}