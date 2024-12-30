package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func processData(ctx context.Context, data io.Reader) error {
	// Open a file for writing
	file, err := os.Create("./example.txt")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Limit the processing time to 5 seconds
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Process the data and write to the file
	_, err = io.Copy(file, data)
	if err != nil {
		return fmt.Errorf("error processing data: %w", err)
	}

	// Check if the context is canceled
	select {
	case <-ctx.Done():
		return fmt.Errorf("context canceled: %w", ctx.Err())
	default:
	}

	return nil
}

func main() {
	// Simulate large data source
	data := make([]byte, 100) // 100 MB of data


    fmt.Println(data)

	// Create a context
	ctx := context.Background()

	// Process the data
	err := processData(ctx, os.Stdin)
	if err != nil {
		fmt.Println("Error processing data:", err)
	} else {
		fmt.Println("Data processed successfully.")
	}
}