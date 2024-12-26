package main

import (
	"database/sql"
	"fmt"
	"strings"
    _ "github.com/lib/pq"
)

func main() {
	// Example DB connection
	db, err := sql.Open("postgres", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
}