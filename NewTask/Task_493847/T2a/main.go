package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
}

func initDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./expenses.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddTransaction(db *sql.DB, amount float64, category string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO transactions (amount, category) VALUES (?, ?);")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(amount, category)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ListTransactions(db *sql.DB) ([]Transaction, error) {
	rows, err := db.Query("SELECT id, amount, category, date FROM transactions;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		err = rows.Scan(&t.ID, &t.Amount, &t.Category, &t.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func main() {
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Add new transaction
	id, err := AddTransaction(db, 100.0, "Groceries")
	if err != nil {
		log.Fatalf("Error adding transaction: %v", err)
	}
	fmt.Printf("Transaction added with ID: %d\n", id)

	// List transactions
	transactions, err := ListTransactions(db)
	if err != nil {
		log.Fatalf("Error listing transactions: %v", err)
	}
	fmt.Println("Transactions:")
	for _, t := range transactions {
		fmt.Printf("ID: %d, Amount: %.2f, Category: %s, Date: %s\n", t.ID, t.Amount, t.Category, t.Date.Format("2006-01-02"))
	}
}
