package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbConnectionString = "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	maxIdleConnections = 10
	maxOpenConnections = 100
)

var (
	// Database connection pool variable
	db *gorm.DB

	// WaitGroup for synchronization
	wg sync.WaitGroup
)

// User struct representing the users table
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

func main() {
	// Initialize the database connection pool
	initDB()
	defer closeDB()

	// Simulate concurrent database operations
	numOperations := 1000
	wg.Add(numOperations)

	for i := 0; i < numOperations; i++ {
		go performTransaction()
	}

	wg.Wait()
	fmt.Println("All transactions completed.")
}

func initDB() {
	var err error
	// Use the `mysql` driver from `gorm.io/driver/mysql`
	db, err = gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Set connection pool configurations
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to retrieve generic DB object: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Create the "users" table if it doesn't exist
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("Failed to migrate database schema: " + err.Error())
	}
}

func closeDB() {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func performTransaction() {
	defer wg.Done()

	// Create a context with a 30-second deadline
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Begin a new transaction
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			// Abort the transaction if it panics
			tx.Rollback()
			fmt.Println("Transaction rolled back due to panic:", r)
		}
	}()

	// Simulate a database operation
	name := fmt.Sprintf("User%d", rand.Intn(1000))
	user := User{Name: name}

	// Insert user into the database
	if err := tx.Create(&user).Error; err != nil {
		fmt.Println("Transaction failed:", err)
		tx.Rollback()
		return
	}

	// Simulate a random delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		fmt.Println("Failed to commit transaction:", err)
	}
}
