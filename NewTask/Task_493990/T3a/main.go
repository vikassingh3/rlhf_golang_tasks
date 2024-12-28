package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// CSV record structure with default values for missing fields
type Record struct {
	Name    string  `csv:"Name"`
	Age     int     `csv:"Age"`
	Email   string  `csv:"Email"`
	Balance float64 `csv:"Balance"`
}

func main() {
	filePath := "data.csv"

	// Determine the number of concurrent workers
	workers := runtime.NumCPU()
	maxRecordBuffer := 1000 // Capability to buffer a certain number of records

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute) // Timeout to prevent infinite waits
	defer cancel()

	recordChannel := make(chan Record, maxRecordBuffer)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for record := range recordChannel {
				processRecord(record)
			}
		}()
	}

	// Read CSV records and send them to the channel
	go func() {
		defer close(recordChannel)

		// Skip header line
		_, err := reader.Read()
		if err != nil {
			log.Fatalf("Error skipping header: %v", err)
		}

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error reading CSV record: %v", err)
				continue
			}

			// Parse record fields
			var r Record
			if len(record) >= 1 {
				r.Name = record[0]
			}
			if len(record) >= 2 {
				age, err := strconv.Atoi(record[1])
				if err != nil {
					r.Age = 0
					log.Printf("Invalid age format: %v, setting default value 0", record[1])
				} else {
					r.Age = age
				}
			}
			if len(record) >= 3 {
				r.Email = record[2]
			}
			if len(record) >= 4 {
				balance, err := strconv.ParseFloat(record[3], 64)
				if err != nil {
					r.Balance = 0.0
					log.Printf("Invalid balance format: %v, setting default value 0.0", record[3])
				} else {
					r.Balance = balance
				}
			}

			select {
			case recordChannel <- r:
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
	log.Println("All records processed successfully.")
}

func processRecord(record Record) {
	fmt.Printf("Processed Record: %+v\n", record)
}
