package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Define a struct to represent the patient data
type Patient struct {
	PatientID   int    `db:"patient_id"`
	Name        string `db:"name"`
	Age         int    `db:"age"`
	Gender      string `db:"gender"`
	Diagnosis   string `db:"diagnosis"`
	Medication string `db:"medication"`
}

func main() {
	// Open the MySQL database connection
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/healthcare_db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT patient_id, name, age, gender, diagnosis, medication FROM patients"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	// Iterate through the results using a range loop
	var patients []Patient
	for rows.Next() {
		var patient Patient
		if err := rows.Scan(&patient.PatientID, &patient.Name, &patient.Age, &patient.Gender, &patient.Diagnosis, &patient.Medication); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		patients = append(patients, patient)
	}

	// Print the patient data
	fmt.Println("Patient Data:")
	for _, patient := range patients {
		fmt.Printf("Patient ID: %d, Name: %s, Age: %d, Gender: %s, Diagnosis: %s, Medication: %s\n",
			patient.PatientID, patient.Name, patient.Age, patient.Gender, patient.Diagnosis, patient.Medication)
	}
} 