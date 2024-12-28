package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)  
const (  
    numWorkers = 4 // Number of worker goroutines for concurrency  
    bufferSize  = 100 // Buffer size for channel communication  
)  
// Record struct with default values for missing fields  
type Record struct {  
    Name    string `csv:"Name"`  
    Age     int `csv:"Age"`  
    Email   string `csv:"Email"`  
    Balance float64 `csv:"Balance"`  
}  
func main() {  
    filePath := "./data.csv"  
    start := time.Now()  
    records, err := parseCSVConcurrently(filePath)  
    if err != nil {  
        log.Fatalf("Error parsing CSV file: %v", err)  
    }  
    elapsed := time.Since(start)  
    fmt.Println("CSV Parsing Result:")  
    fmt.Println("Total Records:", len(records))  
    fmt.Println("Parsing Time:", elapsed)  
}  
func parseCSVConcurrently(filePath string) ([]Record, error) {  
    file, err := os.Open(filePath)  
    if err != nil {  
        return nil, fmt.Errorf("error opening CSV file: %v", err)  
    }  
    defer file.Close()  
    reader := csv.NewReader(bufio.NewReader(file))  
    recordCh := make(chan Record, bufferSize)  
    var wg sync.WaitGroup  
    wg.Add(numWorkers)  
    // Launch worker goroutines to parse records concurrently  
    for i := 0; i < numWorkers; i++ {  
        go func() {  
            defer wg.Done()  
            for record := range recordCh {  
                parseRecord(&record)  
            }  
        }()  
    }  
    go func() {  
        // Close the record channel once all records have been sent  
        defer close(recordCh)  
        for {  
            _, err := reader.Read()  
            if err == io.EOF {  
                break  
            }  
            if err != nil {  
                log.Printf("Error reading CSV record: %v", err)  
                continue  
            }  
            var r Record  
            // Send the CSV record to the worker goroutines for parsing  
            recordCh <- r  
        }  
    }()  
    // Wait for all worker goroutines to finish parsing  
    wg.Wait()  
    var records []Record  
    // Collect all the parsed records from the worker goroutines  
    for record := range recordCh {  
        records = append(records, record)  
    }  
    return records, nil  
}  
// parseRecord parses a single CSV record and handles errors  
func parseRecord(record *Record) {  
    // Implement your parsing logic here  
    // For demonstration purposes, we'll just set default values  
    record.Name = "Default Name"  
    record.Age = 0  
    record.Email = "default@example.com"  
    record.Balance = 0.0  
}  
