package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	// Connection string to the PostgreSQL database
	connString := "host=localhost dbname=postgres password=root sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		// Handle connection failure
		fmt.Printf("Error opening connection: %v\n", err)
		return
	}
	defer db.Close() // Ensure the database connection is closed

	// Query to execute
	query := "SELECT 1"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		// Handle query execution error
		fmt.Printf("Error executing query: %v\n", err)
		return
	}
	defer rows.Close() // Ensure the rows are closed

	// Iterate over the rows and scan the result
	var result int
	if rows.Next() { // Call Next() to move to the first row
		if err := rows.Scan(&result); err != nil {
			// Handle scanning error
			fmt.Printf("Error scanning result: %v\n", err)
			return
		}
	} else {
		// No rows returned
		fmt.Println("No rows found.")
		return
	}

	// Output the result
	fmt.Println("Query executed successfully:", result)
}
