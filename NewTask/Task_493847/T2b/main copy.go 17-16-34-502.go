package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite3 driver
)

// Transaction struct defines a transaction record
type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     string
}

// Transactions slice to manage multiple transactions
type Transactions []Transaction

// InitializeDatabase initializes the database connection and creates the transactions table
func (ts *Transactions) InitializeDatabase() *sql.DB {
	// Open or create the SQLite database file
	db, err := sql.Open("sqlite3", "./expenses.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the transactions table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		date TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// AddTransaction to the database using SQLite
func (ts *Transactions) AddTransaction(amount float64, category string) {
	db := ts.InitializeDatabase()
	defer db.Close()

	date := time.Now().Format("2006-01-02")
	insertSQL := "INSERT INTO transactions (amount, category, date) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(amount, category, date)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction added successfully.")
}

// LoadTransactionsFromDatabase loads all transactions from the SQLite database into the slice
func (ts *Transactions) LoadTransactionsFromDatabase() {
	db := ts.InitializeDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT id, amount, category, date FROM transactions")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.Amount, &t.Category, &t.Date); err != nil {
			log.Fatal(err)
		}
		*ts = append(*ts, t)
	}
}

func main() {
	transactions := Transactions{}

	// Load transactions from the database on app start
	transactions.LoadTransactionsFromDatabase()
	fmt.Println("Transactions loaded from database:")

	for _, t := range transactions {
		fmt.Printf("ID: %d, Amount: $%.2f, Category: %s, Date: %s\n", t.ID, t.Amount, t.Category, t.Date)
	}

	// Add a new transaction
	transactions.AddTransaction(300.0, "Clothing")
}
