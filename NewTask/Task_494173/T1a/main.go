package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	// Connection string to the PostgreSQL database
	connString := "host=localhost dbname=teastall password=root sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Printf("Error opening connection: %v\n", err)
		return
	}
	defer db.Close() // Ensure the database connection is closed

	// Ping the database to ensure the connection is working
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		return
	}

	fmt.Println("Database connection established successfully!")
}