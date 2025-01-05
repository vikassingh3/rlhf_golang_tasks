package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Custom error types
var (
	ErrDatabaseTimeout = errors.New("database timeout error")
	ErrConnectionError = errors.New("connection error")
)


func queryDatabase(ctx context.Context, db *gorm.DB) (string, error) {
	// Simulate a database query
	var result string
	select {
	case <-ctx.Done():
		return "", ErrDatabaseTimeout
	default:
		time.Sleep(5 * time.Second) // Simulate a slow query
		result = "Data retrieved"
	}
	return result, nil
}


func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dsn := "host=localhost user=postgres password=root dbname=teastall port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}


	// defer db.Close()

	result, err := queryDatabase(ctx, db)
	if err != nil {
		switch err {
		case ErrDatabaseTimeout:
			fmt.Println("Timeout error:", err)
			// Take specific actions for timeout errors, such as retrying the operation or logging
		case ErrConnectionError:
			fmt.Println("Connection error:", err)
			// Take specific actions for connection errors, such as reconnecting or logging
		default:
			fmt.Println("An unknown error occurred:", err)
		}
	} else {
		fmt.Println("Result:", result)
	}
}