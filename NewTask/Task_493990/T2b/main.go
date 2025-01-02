package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
)

// CSV record structure with default values for missing fields
type Record struct {
    Name    string  `csv:"Name"`
    Age     int     `csv:"Age"`
    Email   string  `csv:"Email"`
    Balance float64 `csv:"Balance"`
}

func main() {
    filePath := "./data.csv"
    records, err := parseCSV(filePath)
    if err != nil {
        log.Fatalf("Error parsing CSV file: %v", err)
    }

    fmt.Println("CSV Parsing Result:")
    for _, record := range records {
        fmt.Println(record)
    }
}

func parseCSV(filePath string) ([]Record, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("error opening CSV file: %v", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records := []Record{}
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("error reading CSV record: %v", err)
        }

        var r Record
        // Parse each field in the CSV record

        // Name (required field)
        if len(record) >= 1 {
            r.Name = record[0]
        } else {
            log.Printf("Missing Name field, skipping record")
            continue // Skip the record if Name is missing
        }

        // Age (optional, default 0 if missing or invalid)
        if len(record) >= 2 {
            age, err := strconv.Atoi(record[1])
            if err != nil {
                r.Age = 0
                log.Printf("Invalid age format: %v, setting default value 0", record[1])
            } else {
                r.Age = age
            }
        } else {
            r.Age = 0
        }

        // Email (optional, default empty if missing)
        if len(record) >= 3 {
            r.Email = record[2]
        } else {
            r.Email = "" // Setting default empty value for missing Email
        }

        // Balance (optional, default to 0.0 if missing or invalid)
        if len(record) >= 4 {
            balance, err := strconv.ParseFloat(record[3], 64)
            if err != nil {
                r.Balance = 0.0
                log.Printf("Invalid balance format: %v, setting default value 0.0", record[3])
            } else {
                r.Balance = balance
            }
        } else {
            r.Balance = 0.0 // Setting default balance value if missing
        }

        records = append(records, r)
    }

    return records, nil
}
