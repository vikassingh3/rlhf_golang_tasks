package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// FileContext holds metadata related to a file operation.
type FileContext struct {
	Context    context.Context
	FileSize   int64
	Operation  string // "read" or "write"
	Status     string
	Duration    time.Duration
	Error      error
}

func readFileWithContext(ctx context.Context, filename string) (*FileContext, error) {
	// Create a child context with a deadline to handle timeouts.
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	startTime := time.Now()

	// Open the file for reading.
	file, err := os.Open(filename)
	if err != nil {
		return &FileContext{Context: ctx, Error: err}, err
	}
	defer file.Close()

	// Read the file content.
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return &FileContext{Context: ctx, Error: err}, err
	}

	fmt.Println(data)

	// Retrieve file information.
	fileInfo, err := file.Stat()
	if err != nil {
		return &FileContext{Context: ctx, Error: err}, err
	}

	// Calculate the duration of the operation.
	duration := time.Since(startTime)

	// Create the FileContext struct with metadata.
	return &FileContext{
		Context:    ctx,
		FileSize:   fileInfo.Size(),
		Operation:  "read",
		Status:     "success",
		Duration:    duration,
		Error:      nil,
	}, nil
}

func writeFileWithContext(ctx context.Context, filename string, data []byte) (*FileContext, error) {
	// Create a child context with a deadline to handle timeouts.
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	startTime := time.Now()

	// Open the file for writing.
	file, err := os.Create(filename)
	if err != nil {
		return &FileContext{Context: ctx, Error: err}, err
	}
	defer file.Close()

	// Write the file content.
	_, err = file.Write(data)
	if err != nil {
		return &FileContext{Context: ctx, Error: err}, err
	}

	// Retrieve file information.
	fileInfo, err := file.Stat()
	if err != nil {
		return &FileContext{Context: ctx, Error: err}, err
	}

	// Calculate the duration of the operation.
	duration := time.Since(startTime)

	// Create the FileContext struct with metadata.
	return &FileContext{
		Context:    ctx,
		FileSize:   fileInfo.Size(),
		Operation:  "write",
		Status:     "success",
		Duration:    duration,
		Error:      nil,
	}, nil
}

func main() {
	// Create a background context.
	ctx := context.Background()

	// Perform a file read operation.
	readCtx, err := readFileWithContext(ctx, "example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Read Operation Metadata:\n")
	fmt.Printf("File Size: %d bytes\n", readCtx.FileSize)
	fmt.Printf("Operation Status: %s\n", readCtx.Status)
	fmt.Printf("Duration: %s\n", readCtx.Duration)

	// Perform a file write operation.
	writeData := []byte("Hello, world!")
	writeCtx, err := writeFileWithContext(ctx, "example_output.txt", writeData)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("\nWrite Operation Metadata:\n")
	fmt.Printf("File Size: %d bytes\n", writeCtx.FileSize)
	fmt.Printf("Operation Status: %s\n", writeCtx.Status)
	fmt.Printf("Duration: %s\n", writeCtx.Duration)
}