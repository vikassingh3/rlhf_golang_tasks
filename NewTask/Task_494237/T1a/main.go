package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	// Connect to the database
	connString := "host=localhost user=your_username dbname=healthcare password=your_password sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Ensure the database connection is closed

	// Prepare a SQL query
	query := "SELECT patient_id, name, age FROM patients"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // Ensure the result set is closed

	// Iterate through the result set using a range loop
	for rows.Next() {
		var patientID int
		var name string
		var age int

		// Scan each row into variables
		err = rows.Scan(&patientID, &name, &age)
		if err != nil {
			log.Fatal(err)
		}

		// Print the result
		fmt.Printf("Patient ID: %d, Name: %s, Age: %d\n", patientID, name, age)
	}

	// Check for any errors during row iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query completed successfully.")
}