package main

import (
	"fmt"
)

// DataMigrator defines the methods required for data migration.
type DataMigrator interface {
	ReadData() ([]byte, error)
	WriteData([]byte) error
}

// MigrateData migrates data from a source to a destination.
func MigrateData(source DataMigrator, destination DataMigrator) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from a panic:", r)
		}
	}()

	for {
		chunk, err := source.ReadData()
		if err != nil {
			return fmt.Errorf("error reading data: %w", err)
		}
		if len(chunk) == 0 {
			break
		}

		err = destination.WriteData(chunk)
		if err != nil {
			return fmt.Errorf("error writing data: %w", err)
		}
	}
	return nil
}

// InMemoryStorage is an in-memory implementation of DataMigrator.
type InMemoryStorage struct {
	data []byte
	err  error
}

// ReadData reads data from the in-memory storage.
func (s *InMemoryStorage) ReadData() ([]byte, error) {
	if s.err != nil {
		return nil, s.err
	}
	// Simulate reading all data at once
	chunk := s.data
	s.data = nil
	return chunk, nil
}

// WriteData writes data to the in-memory storage.
func (s *InMemoryStorage) WriteData(data []byte) error {
	s.data = append(s.data, data...)
	// Simulate an error when exceeding a capacity
	if len(s.data) > 10 {
		s.err = fmt.Errorf("in-memory storage exceeded capacity")
		return s.err
	}
	return nil
}

func main() {
	// Source with initial data
	source := &InMemoryStorage{
		data: []byte("Hello, Go!"),
	}

	// Destination with no data initially
	destination := &InMemoryStorage{}

	// Perform the migration
	err := MigrateData(source, destination)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration successful. Destination data:", string(destination.data))
	}
}
