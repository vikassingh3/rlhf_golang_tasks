package main

import (
	"fmt"
	"io/ioutil"
)

// DataMigrator interface
type DataMigrator interface {
	ReadData() ([]byte, error)
	WriteData(data []byte) error
}

// InMemoryStorage implementation
type InMemoryStorage struct {
	data     []byte
	readDone bool // Track whether data has been fully read
}

func (s *InMemoryStorage) ReadData() ([]byte, error) {
	if s.readDone {
		// Simulate EOF by returning empty data after the data has been read
		return nil, nil
	}
	s.readDone = true
	return s.data, nil
}

func (s *InMemoryStorage) WriteData(data []byte) error {
	s.data = append(s.data, data...)
	return nil
}

// FileStorage implementation
type FileStorage struct {
	filename string
}

func (s *FileStorage) ReadData() ([]byte, error) {
	// Implement file reading logic
	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *FileStorage) WriteData(data []byte) error {
	// Implement file writing logic
	err := ioutil.WriteFile(s.filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// MigrateData migrates data from source to destination
func MigrateData(source DataMigrator, destination DataMigrator) error {
	for {
		chunk, err := source.ReadData()
		if err != nil {
			return err
		}
		if len(chunk) == 0 {
			// No more data to read
			break
		}
		err = destination.WriteData(chunk)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Example usage:
	inMemorySource := &InMemoryStorage{data: []byte("Hello, World!")}
	fileDestination := &FileStorage{filename: "./output.txt"}

	err := MigrateData(inMemorySource, fileDestination)
	if err != nil {
		fmt.Println("Error migrating data:", err)
	} else {
		fmt.Println("Data migration successful.")
	}

	// Verify written data (optional)
	data, err := ioutil.ReadFile("./output.txt")
	if err == nil {
		fmt.Println("Written data:", string(data))
	}
}
