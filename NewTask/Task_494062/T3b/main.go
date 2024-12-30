package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Global variable to store the database connection (reused)
var dbConn *databaseConnection
var dbOnce sync.Once

// Simulate a database connection
type databaseConnection struct {
}

func (db *databaseConnection) Query(query string) ([]byte, error) {
	// Simulate database query
	time.Sleep(1 * time.Second)
	return []byte("Query result"), nil
}

// Acquire database connection lazily
func getDatabaseConnection() *databaseConnection {
	dbOnce.Do(func() {
		dbConn = &databaseConnection{}
	})
	return dbConn
}

func main() {
	// Process multiple network requests
	for i := 0; i < 5; i++ {
		processNetworkRequest()
	}
}

func processNetworkRequest() {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Acquire database connection lazily
	db := getDatabaseConnection()

	// Simulate network request
	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Println("Error making network request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body and query database
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	queryResult, err := db.Query("SELECT * FROM table")
	if err != nil {
		log.Println("Error executing database query:", err)
		return
	}

	// Perform processing using the acquired resources
	fmt.Println("Response body:", string(body))
	fmt.Println("Query result:", string(queryResult))
}