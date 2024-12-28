package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CSVRecord struct {
	Name     string
	Age      int
	Email    string `json:"email,omitempty"`
	PhoneNum string `json:"phone_num,omitempty"`
}

func parseCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	for _, record := range records {
		// Ensure the record has the required number of fields
		if len(record) < 4 {
			log.Printf("Skipping record due to missing fields: %v\n", record)
			continue
		}

		// Parse the record into the CSVRecord struct
		var csvRecord CSVRecord
		csvRecord.Name = record[0]
		csvRecord.Age, _ = strconv.Atoi(record[1])
		csvRecord.Email = record[2]
		csvRecord.PhoneNum = record[3]

		fmt.Println("CSV Parsing Result:")
		fmt.Printf("Name: %s, Age: %d, Email: %s, Phone Num: %s\n",
			csvRecord.Name, csvRecord.Age, csvRecord.Email, csvRecord.PhoneNum)
	}
}

func main() {
	parseCSV("./people.csv")
}
