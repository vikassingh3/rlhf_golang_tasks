package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

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

	// Query to retrieve all patient data along with their medication information.
	query := `
	SELECT p.patient_id, p.name, p.age, m.medication_name, m.dose
	FROM patients p
	LEFT JOIN medications m ON p.patient_id = m.patient_id;
	`

	// Define a slice to store patient data with medication information.
	type PatientMedication struct {
		ID            int
		Name          string
		Age           int
		Medication    string
		Dose          string
		HasMedication bool // To indicate if a patient has medication information
	}
	var patients []PatientMedication

	// Execute the query and scan the results into the patients slice.
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	// Loop through each row of result and append patient data.
	for rows.Next() {
		var p PatientMedication
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Medication, &p.Dose); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		// Check if medication information is present
		if p.Medication != "" && p.Dose != "" {
			p.HasMedication = true
		}
		patients = append(patients, p)
	}

	// Check for any row iteration errors
	if err = rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v", err)
	}

	// Process patients with medication information
	fmt.Println("Patients with medication:")
	for _, p := range patients {
		if p.HasMedication {
			fmt.Printf("Patient ID: %d, Name: %s, Age: %d, Medication: %s, Dose: %s\n", p.ID, p.Name, p.Age, p.Medication, p.Dose)
		}
	}

	// Process patients without medication information
	fmt.Println("\nPatients without medication:")
	for _, p := range patients {
		if !p.HasMedication {
			fmt.Printf("Patient ID: %d, Name: %s, Age: %d (No medication information)\n", p.ID, p.Name, p.Age)
		}
	}

	// Calculate statistics on patients with medication (excluding those without)
	var medicationCount int
	var totalDose float32
	for _, p := range patients {
		if p.HasMedication {
			medicationCount++
			// Convert dose to float32 and add to total
			doseFloat, err := ParseDose(p.Dose)
			if err != nil {
				log.Printf("Error parsing dose for patient %d: %v", p.ID, err)
				continue
			}
			totalDose += doseFloat
		}
	}

	// Handle edge case: No medication information found
	if medicationCount > 0 {
		averageDose := totalDose / float32(medicationCount)
		fmt.Printf("\nAverage Dose among patients with medication: %.2f\n", averageDose)
	} else {
		fmt.Println("\nNo patients with medication found to calculate average dose.")
	}

	// Handle edge case: ParseDose error
	fmt.Println("\nPatients with invalid dose format:")
	for _, p := range patients {
		if _, err := ParseDose(p.Dose); err != nil {
			fmt.Printf("Patient ID: %d, Name: %s, Dose: %s (Invalid format)\n", p.ID, p.Name, p.Dose)
		}
	}
}

// ParseDose converts the dose string to a float32. If the conversion fails, it returns an error.
func ParseDose(dose string) (float32, error) {
	doseFloat, err := strconv.ParseFloat(dose, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid dose format: %v", err)
	}
	return float32(doseFloat), nil
}
