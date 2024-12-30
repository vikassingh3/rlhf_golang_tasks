package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Example DB connection
	dsn := "user:password@tcp(localhost:3306)/database" // Update with your actual database credentials
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verify the connection
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}

	// User input
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	// Prepared statement
	stmt, err := db.Prepare("SELECT * FROM users WHERE name = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute with parameter
	rows, err := stmt.Query(name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Process rows
	for rows.Next() {
		var userName, userEmail string
		err = rows.Scan(&userName, &userEmail)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s, Email: %s\n", userName, userEmail)
	}

	// Handle no rows returned
	if rows.Err() != nil {
		fmt.Println("No users found.")
	}
}
