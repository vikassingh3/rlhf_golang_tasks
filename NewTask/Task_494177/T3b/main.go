package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn               = "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	maxIdleConnections = 10
	maxOpenConnections = 100
)

var (
	db *gorm.DB
	wg sync.WaitGroup
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func main() {
	// Initialize the database
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer closeDB()

	// Perform transactions concurrently
	numOperations := 100
	wg.Add(numOperations)
	for i := 0; i < numOperations; i++ {
		go performTransaction()
	}
	wg.Wait()
	fmt.Println("All transactions completed.")
}

func initDB() error {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Configure the connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Automatically migrate the User schema
	if err := db.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("failed to migrate schema: %w", err)
	}

	return nil
}

func closeDB() {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func performTransaction() {
	defer wg.Done()

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Begin a transaction
	tx := db.WithContext(ctx).Begin()
	if err := tx.Create(&User{Name: "John Doe"}).Error; err != nil {
		tx.Rollback()
		log.Printf("Transaction failed: %v", err)
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return
	}

	log.Println("Transaction succeeded")
}
