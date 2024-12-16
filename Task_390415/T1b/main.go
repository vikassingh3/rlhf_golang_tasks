package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// ProcessFile reads a file and processes its contents.
func ProcessFile(ctx context.Context, filePath string) error {
	// Simulate a long-running file read operation
	select {
	case <-time.After(3 * time.Second): // Simulate 5-second read
		fmt.Println("File reading completed.")
		return nil
	case <-ctx.Done():
		log.Printf("File processing cancelled: %v", ctx.Err())
		return ctx.Err() // Return the cancellation error
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure the context is cancelled, even on panic

	filePath := "example.txt" // Replace with your file path

	err := ProcessFile(ctx, filePath)
	if err != nil {
		fmt.Println("Error processing file:", err)
	}
}
