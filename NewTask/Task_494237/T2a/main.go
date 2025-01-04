package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import PostgreSQL driver
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

	// Variables to keep track of total age and count for filtered patients
	totalAge := 0
	count := 0
	filteredPatients := []int{} // To store IDs of patients that meet the age criteria

	// Iterate through query results using a range loop
	for rows.Next() {
		var patientID int
		var name string
		var age int

		// Scan row data into variables
		if err := rows.Scan(&patientID, &name, &age); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		// Filter patients older than 40
		if age > 40 {
			filteredPatients = append(filteredPatients, patientID)
			totalAge += age
			count++

			fmt.Printf("Patient ID: %d, Name: %s, Age: %d (Matches criteria)\n", patientID, name, age)
		} else {
			fmt.Printf("Patient ID: %d, Name: %s, Age: %d (Does not match criteria)\n", patientID, name, age)
		}
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v", err)
	}

	// Calculate and print the average age of filtered patients
	if count > 0 {
		averageAge := float64(totalAge) / float64(count)
		fmt.Printf("Average age of patients older than 40: %.2f\n", averageAge)
	} else {
		fmt.Println("No patients older than 40 found.")
	}

	fmt.Println("Query completed successfully.")
}