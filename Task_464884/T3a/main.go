package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// FileMetadata represents the metadata related to a file operation.
type FileMetadata struct {
	FileSize   int64
	Operation  string // "read" or "write"
	Status     string
	Duration    time.Duration
	Error      error
}

// WithFileMetadata sets the given FileMetadata in the provided context.
func WithFileMetadata(ctx context.Context, metadata *FileMetadata) context.Context {
	return context.WithValue(ctx, "FileMetadata", metadata)
}

// GetFileMetadata retrieves the FileMetadata from the provided context.
func GetFileMetadata(ctx context.Context) (*FileMetadata, bool) {
	md, ok := ctx.Value("FileMetadata").(*FileMetadata)
	return md, ok
}

func readFileWithContext(ctx context.Context, filename string) (*FileMetadata, error) {
	// Create a child context with a deadline to handle timeouts.
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	startTime := time.Now()

	// Open the file for reading.
	file, err := os.Open(filename)
	if err != nil {
		return &FileMetadata{
			Operation: "read",
			Status:    "failed",
			Error:      err,
		}, err
	}
	defer file.Close()

	// Read the file content.
	data, err := ioutil.ReadAll(file)
	fmt.Println(data)
	if err != nil {
		return &FileMetadata{
			Operation: "read",
			Status:    "failed",
			Error:      err,
		}, err
	}

	// Retrieve file information.
	fileInfo, err := file.Stat()
	if err != nil {
		return &FileMetadata{
			Operation: "read",
			Status:    "failed",
			Error:      err,
		}, err
	}

	// Calculate the duration of the operation.
	duration := time.Since(startTime)

	// Create the FileMetadata struct with successful metadata.
	return &FileMetadata{
		FileSize:   fileInfo.Size(),
		Operation:  "read",
		Status:    "success",
		Duration:    duration,
		Error:      nil,
	}, nil
}

func writeFileWithContext(ctx context.Context, filename string, data []byte) (*FileMetadata, error) {
	// Create a child context with a deadline to handle timeouts.
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	startTime := time.Now()

	// Open the file for writing.
	file, err := os.Create(filename)
	if err != nil {
		return &FileMetadata{
			Operation: "write",
			Status:    "failed",
			Error:      err,
		}, err
	}
	defer file.Close()

	// Write the file content.
	_, err = file.Write(data)
	if err != nil {
		return &FileMetadata{
			Operation: "write",
			Status:    "failed",
			Error:      err,
		}, err
	}

	// Retrieve file information.
	fileInfo, err := file.Stat()
	if err != nil {
		return &FileMetadata{
			Operation: "write",
			Status:    "failed",
			Error:      err,
		}, err
	}

	// Calculate the duration of the operation.
	duration := time.Since(startTime)

	// Create the FileMetadata struct with successful metadata.
	return &FileMetadata{
		FileSize:   fileInfo.Size(),
		Operation:  "write",
		Status:    "success",
		Duration:    duration,
		Error:      nil,
	}, nil
}

func main() {
	// Create a background context.
	ctx := context.Background()

	// Read file and monitor its operation metadata.
	readMetadata, err := readFileWithContext(ctx, "example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Read Operation Metadata:\n")
	fmt.Printf("File Size: %d bytes\n", readMetadata.FileSize)
	fmt.Printf("Operation Status: %s\n", readMetadata.Status)
	fmt.Printf("Duration: %s\n", readMetadata.Duration)

	// Write file and monitor its operation metadata.
	writeData := []byte("Hello, world!")
	writeMetadata, err := writeFileWithContext(ctx, "example_output.txt", writeData)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("\nWrite Operation Metadata:\n")
	fmt.Printf("File Size: %d bytes\n", writeMetadata.FileSize)
	fmt.Printf("Operation Status: %s\n", writeMetadata.Status)
	fmt.Printf("Duration: %s\n", writeMetadata.Duration)
}