package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection details.
	connString := "host=localhost user=your_username dbname=healthcare password=your_password sslmode=disable"

	// Connect to the database.
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Query to retrieve all patient data.
	query := "SELECT patient_id, name, age, height, weight FROM patients"

	// Define a slice to store patient data.
	type Patient struct {
		ID       int
		Name     string
		Age      int
		Height   float32
		Weight   float32
		BMI      float32
		IsAdult  bool
		IsValid  bool
	}
	var patients []Patient

	// Execute the query and scan the results into the patients slice.
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var p Patient
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Height, &p.Weight); err != nil {
			log.Printf("Error scanning row: %v. Skipping patient ID %d.\n", err, p.ID)
			continue // Skip this row if scanning fails
		}

		// Check for missing or inconsistent data
		if p.Age < 0 || p.Height <= 0 || p.Weight <= 0 {
			log.Printf("Inconsistent data for patient ID %d: age=%d, height=%.2f, weight=%.2f. Marking as invalid.\n", p.ID, p.Age, p.Height, p.Weight)
			p.IsValid = false
		} else {
			p.IsValid = true
			p.BMI = p.Weight / (p.Height * p.Height)
			p.IsAdult = p.Age >= 18
		}

		patients = append(patients, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v", err)
	}

	// Filter and print patients who are adults, have a valid BMI, and BMI > 25.
	fmt.Println("Adults with valid BMI > 25:")
	for _, p := range patients {
		if p.IsAdult && p.IsValid && p.BMI > 25 {
			fmt.Printf("Patient ID: %d, Name: %s, BMI: %.2f\n", p.ID, p.Name, p.BMI)
		}
	}

	// Calculate the average age of valid patients.
	totalValidAge := 0
	validPatientCount := 0
	for _, p := range patients {
		if p.IsValid {
			totalValidAge += p.Age
			validPatientCount++
		}
	}
	if validPatientCount > 0 {
		averageValidAge := float32(totalValidAge) / float32(validPatientCount)
		fmt.Printf("Average Age (valid patients): %.2f\n", averageValidAge)
	} else {
		fmt.Println("No valid patients found to calculate average age.")
	}

	// Calculate the maximum age of valid patients.
	maxValidAge := 0
	for _, p := range patients {
		if p.IsValid && p.Age > maxValidAge {
			maxValidAge = p.Age
		}
	}
	if maxValidAge > 0 {
		fmt.Printf("Maximum Age (valid patients): %d\n", maxValidAge)
	} else {
		fmt.Println("No valid patients found to calculate maximum age.")
	}
}