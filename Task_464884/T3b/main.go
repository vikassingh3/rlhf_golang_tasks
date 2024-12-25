package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// FileOperation represents a single file operation with metadata.
type FileOperation struct {
	Path     string
	Size     int64
	Status   string // "success" or "error"
	Error    error
	Started  time.Time
	Finished time.Time
}

// FileMonitor monitors file operations and tracks their status.
type FileMonitor struct {
	Operations []FileOperation
	mu         sync.Mutex
}

// NewFileMonitor creates a new FileMonitor.
func NewFileMonitor() *FileMonitor {
	return &FileMonitor{
		Operations: []FileOperation{},
	}
}

// Add adds a new file operation to the monitor.
func (m *FileMonitor) Add(op FileOperation) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Operations = append(m.Operations, op)
}

// GetOperations returns a copy of the current operations list.
func (m *FileMonitor) GetOperations() []FileOperation {
	m.mu.Lock()
	defer m.mu.Unlock()
	ops := make([]FileOperation, len(m.Operations))
	copy(ops, m.Operations)
	return ops
}

// copyFile copies the source file to the destination file with context monitoring.
func copyFile(ctx context.Context, source, destination string, monitor *FileMonitor) error {
	// Create a new FileOperation for this operation.
	op := FileOperation{
		Path:    destination,
		Started: time.Now(),
	}
	defer func() {
		op.Finished = time.Now()
		monitor.Add(op)
	}()

	// Check if the context has been canceled before starting the operation.
	select {
	case <-ctx.Done():
		op.Status = "error"
		op.Error = ctx.Err()
		return op.Error
	default:
	}

	// Open the source file for reading.
	sourceFile, err := os.Open(source)
	if err != nil {
		op.Status = "error"
		op.Error = err
		return err
	}
	defer sourceFile.Close()

	// Stat the source file to get its size.
	fi, err := sourceFile.Stat()
	if err != nil {
		op.Status = "error"
		op.Error = err
		return err
	}
	op.Size = fi.Size()

	// Create the destination file for writing.
	destinationFile, err := os.Create(destination)
	if err != nil {
		op.Status = "error"
		op.Error = err
		return err
	}
	defer destinationFile.Close()

	// Copy the file content using io.Copy.
	n, err := io.Copy(destinationFile, sourceFile)
	if err != nil {
		op.Status = "error"
		op.Error = err
		return err
	}

	// Check if the number of bytes copied matches the source file size.
	if n != op.Size {
		op.Status = "error"
		op.Error = fmt.Errorf("copied %d bytes, expected %d bytes", n, op.Size)
		return op.Error
	}

	// Set the operation status to "success" if no errors occurred.
	op.Status = "success"
	return nil
}

func main() {
	// Example usage of the copyFile function.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	source := "source.txt"
	destination := filepath.Join("dest", "copied.txt")

	// Create a FileMonitor instance.
	monitor := NewFileMonitor()

	// Ensure the destination directory exists.
	if err := os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		fmt.Println("Error creating destination directory:", err)
		return
	}

	// Perform the file copy operation.
	err := copyFile(ctx, source, destination, monitor)
	if err != nil {
		fmt.Println("Error during file copy:", err)
	} else {
		fmt.Println("File copied successfully.")
	}

	// Print the file operation logs.
	for _, op := range monitor.GetOperations() {
		fmt.Printf("Path: %s, Size: %d, Status: %s, Error: %v, Duration: %v\n",
			op.Path, op.Size, op.Status, op.Error, op.Finished.Sub(op.Started))
	}
}
