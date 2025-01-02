package main

import (
	"fmt"
	"io"
	"time"
)

// DataMigrator interface for data migration
type DataMigrator interface {
	ReadData() ([]byte, error)
	WriteData(data []byte) error
}

// MigrateError represents a custom error type for migration errors
type MigrateError struct {
	Cause  error
	Action string
	Data   []byte // Add this field if you want to retain the chunk that caused the error
}

func (e *MigrateError) Error() string {
	return fmt.Sprintf("migrate error: action=%s, cause=%v", e.Action, e.Cause)
}

func (e *MigrateError) SourceError() bool {
	return e.Action == "read"
}

func (e *MigrateError) DestinationError() bool {
	return e.Action == "write"
}

// InMemoryStorage is a mock implementation of DataMigrator for in-memory data
type InMemoryStorage struct {
	data []byte
	read bool
}

func (s *InMemoryStorage) ReadData() ([]byte, error) {
	if s.read {
		return nil, io.EOF // Simulate end of data
	}
	s.read = true
	return s.data, nil
}

func (s *InMemoryStorage) WriteData(data []byte) error {
	return fmt.Errorf("in-memory storage is read-only")
}

// FileStorage is a mock implementation of DataMigrator for file-based data
type FileStorage struct {
	filename string
}

func (s *FileStorage) ReadData() ([]byte, error) {
	return nil, fmt.Errorf("file storage read not implemented")
}

func (s *FileStorage) WriteData(data []byte) error {
	fmt.Printf("Writing to file: %s\n", s.filename)
	return nil // Simulate successful write
}

func MigrateData(source DataMigrator, destination DataMigrator, chunkSize int) error {
	for {
		chunk, err := source.ReadData()
		if err != nil {
			if err == io.EOF {
				break
			}
			return &MigrateError{Cause: err, Action: "read"}
		}

		err = destination.WriteData(chunk)
		if err != nil {
			return &MigrateError{Cause: err, Action: "write", Data: chunk}
		}
	}
	return nil
}

func recoverFromDestinationError(destination DataMigrator, chunk []byte, cause error) error {
	// For this example, we recover only from transient errors
	if cause != io.ErrShortWrite {
		return cause // Do not recover from non-transient errors
	}

	// Retry after a short wait
	time.Sleep(time.Second)
	fmt.Println("Recovering from transient write error, retrying...")
	return destination.WriteData(chunk)
}

func main() {
	inMemorySource := &InMemoryStorage{data: []byte("Hello, World!")}
	fileDestination := &FileStorage{filename: "./output.txt"}

	err := MigrateData(inMemorySource, fileDestination, 1024)
	if err != nil {
		switch e := err.(type) {
		case *MigrateError:
			if e.DestinationError() {
				fmt.Println("Handling destination error...")
				newErr := recoverFromDestinationError(fileDestination, e.Data, e.Cause)
				if newErr != nil {
					fmt.Println("Failed to recover from destination error:", newErr)
				} else {
					fmt.Println("Recovery successful!")
				}
			} else {
				fmt.Println("Migration error:", err)
			}
		default:
			fmt.Println("Unexpected error:", err)
		}
	} else {
		fmt.Println("Data migration successful.")
	}
}
