package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	// Database connection string
	connString := "host=localhost user=your_username dbname=healthcare password=your_password sslmode=disable"

	// Connect to the database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Query to retrieve patient data
	query := "SELECT patient_id, name, age FROM patients"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	// Iterate through the result set using a range loop
	for rows.Next() {
		var patientID int
		var name string
		var age int

		// Scan row data into variables
		if err := rows.Scan(&patientID, &name, &age); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		// Print each record
		fmt.Printf("Patient ID: %d, Name: %s, Age: %d\n", patientID, name, age)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v", err)
	}

	fmt.Println("Query completed successfully.")
}
