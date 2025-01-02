package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type DataMigrator interface {
	ReadData() ([]byte, error)
	WriteData(data []byte) error
}

type InMemoryStorage struct {
	data []byte
}

func (s *InMemoryStorage) ReadData() ([]byte, error) {
	return s.data, nil
}

func (s *InMemoryStorage) WriteData(data []byte) error {
	s.data = append(s.data, data...)
	return nil
}

type FileStorage struct {
	filename string
}

func (s *FileStorage) ReadData() ([]byte, error) {
	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *FileStorage) WriteData(data []byte) error {
	err := ioutil.WriteFile(s.filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func MigrateData(source DataMigrator, destination DataMigrator, chunkSize int) error {
	data, err := source.ReadData()
	if err != nil {
		return fmt.Errorf("error reading data from source: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			// Rollback changes (e.g., delete temporary files)
			if fs, ok := destination.(*FileStorage); ok {
				os.Remove(fs.filename)
			}
			panic(r)
		}
	}()

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data) // Adjust the slice to not exceed the data length
		}
		chunk := data[i:end]
		err := destination.WriteData(chunk)
		if err != nil {
			return fmt.Errorf("error writing data to destination: %w", err)
		}
	}

	return nil
}

func main() {
	// Example usage:
	inMemorySource := &InMemoryStorage{data: []byte("Hello, World!")}
	fileDestination := &FileStorage{filename: "output.txt"}

	chunkSize := 1024 // in bytes
	err := MigrateData(inMemorySource, fileDestination, chunkSize)
	if err != nil {
		fmt.Println("Error migrating data:", err)
	} else {
		fmt.Println("Data migration successful.")
	}

	// Verify the written file (optional)
	data, err := ioutil.ReadFile("output.txt")
	if err == nil {
		fmt.Println("Written data:", string(data))
	}
}
