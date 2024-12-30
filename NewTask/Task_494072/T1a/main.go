package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Data structure for a record
type Record struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Created time.Time `db:"created"`
}

// Function to connect to the source SQLite database
func connectSource() (*sql.DB, error) {
	return sql.Open("sqlite3", "./source.db")
}

// Function to connect to the destination SQLite database
func connectDestination() (*sql.DB, error) {
	return sql.Open("sqlite3", "./destination.db")
}

// Function to read records in chunks
func readRecordsInChunks(db *sql.DB, chunkSize int) []Record {
	records := make([]Record, 0, chunkSize)
	rows, err := db.Query("SELECT * FROM records")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.ID, &record.Name, &record.Created); err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
		// If chunk size is reached, yield the slice and reset it
		if len(records) == chunkSize {
			yield(records)
			records = make([]Record, 0, chunkSize)
		}
	}

	// Yield the remaining records if any
	if len(records) > 0 {
		yield(records)
	}

	return nil
}

// Function to yield a slice of records
func yield(records []Record) {
	fmt.Println("Migrating", len(records), "records")
	// Implement logic to insert records into the destination database here
}

func main() {
	sourceDB, err := connectSource()
	if err != nil {
		log.Fatal(err)
	}
	defer sourceDB.Close()

	destinationDB, err := connectDestination()
	if err != nil {
		log.Fatal(err)
	}
	defer destinationDB.Close()

	chunkSize := 1000 // Define the size of each data chunk
	readRecordsInChunks(sourceDB, chunkSize)

	fmt.Println("Migration completed!")
}