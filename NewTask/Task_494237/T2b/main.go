
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
		IsAdult bool
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
			log.Fatalf("Error scanning row: %v", err)
		}
		patients = append(patients, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v", err)
	}

	// Calculate BMI for each patient and categorize them as adults or children.
	for i := range patients {
		patients[i].BMI = patients[i].Weight / (patients[i].Height * patients[i].Height)
		patients[i].IsAdult = patients[i].Age >= 18
	}

	// Filter and print patients who are adults and have a BMI above 25.
	fmt.Println("Adults with BMI > 25:")
	for _, p := range patients {
		if p.IsAdult && p.BMI > 25 {
			fmt.Printf("Patient ID: %d, Name: %s, BMI: %.2f\n", p.ID, p.Name, p.BMI)
		}
	}

	// Calculate the average age of patients.
	totalAge := 0
	for _, p := range patients {
		totalAge += p.Age
	}
	averageAge := float32(totalAge) / float32(len(patients))
	fmt.Printf("Average Age: %.2f\n", averageAge)

	// Calculate the maximum age of patients.
	maxAge := 0
	for _, p := range patients {
		if p.Age > maxAge {
			maxAge = p.Age
		}
	}
	fmt.Printf("Maximum Age: %d\n", maxAge)
} 
