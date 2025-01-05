package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SimulateTransaction simulates a database transaction
func SimulateTransaction(ctx context.Context, db *gorm.DB) (string, error) {
	// Simulate a slow transaction
	duration := time.Duration(rand.Intn(3)) * time.Second

	select {
	case <-time.After(duration):
		return "Transaction completed", nil
	case <-ctx.Done():
		return "", ctx.Err() // Return context error if timeout occurs
	}
}

func main() {
	// Initialize connection pool settings
	dsn := "host=localhost user=yourusername password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to get generic database object:", err)
		return
	}
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Starting transaction workers...")

	// Create a worker pool to simulate concurrent transactions
	const numWorkers = 10
	workerCh := make(chan struct{})
	defer close(workerCh)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				select {
				case <-workerCh:
					fmt.Printf("Worker %d stopping...\n", workerID)
					return // Stop working
				default:
				}

				// Set context timeout
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				result, err := SimulateTransaction(ctx, db)
				if err != nil {
					if err == context.DeadlineExceeded {
						fmt.Printf("Worker %d: Timeout occurred during transaction: %v\n", workerID, err)
					} else {
						fmt.Printf("Worker %d: Transaction failed: %v\n", workerID, err)
					}
				} else {
					fmt.Printf("Worker %d: Transaction result: %s\n", workerID, result)
				}

				// Balance load by sleeping briefly
				time.Sleep(100 * time.Millisecond)
			}
		}(i + 1)
	}

	// Allow workers to run for a period before stopping
	time.Sleep(10 * time.Second)
	fmt.Println("Stopping transaction workers...")
}
