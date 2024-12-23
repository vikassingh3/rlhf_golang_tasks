package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	// Open a connection to your database
	db, err := sql.Open("your_database_driver", "your_database_connection_string")
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("error starting transaction: %v", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("error rolling back transaction: %v", err)
		}
	}()

	// Perform some operations
	err = insertData(tx)
	if err != nil {
		log.Printf("error inserting data: %v", err)
		return
	}

	err = updateData(tx)
	if err != nil {
		log.Printf("error updating data: %v", err)
		return
	}

	// Commit the transaction if all operations are successful
	if err := tx.Commit(); err != nil {
		log.Fatalf("error committing transaction: %v", err)
	}
	fmt.Println("Transaction committed successfully.")
}

func insertData(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO your_table (column1, column2) VALUES (?, ?)", "value1", "value2")
	return err
}

func updateData(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE your_table SET column3 = ? WHERE column1 = ?", "value3", "value1")
	return err
}