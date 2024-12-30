package main

import (
	"fmt"
	"io"
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
	// Write data to a temporary file first
	tmpFile, err := ioutil.TempFile("", "migration")
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(data)
	if err != nil {
		return err
	}

	// Rename the temporary file to the destination file
	err = os.Rename(tmpFile.Name(), s.filename)
	if err != nil {
		return err
	}

	return nil
}

func MigrateData(source DataMigrator, destination DataMigrator) error {
	for {
		chunk, err := source.ReadData()
		if err != nil {
			// If the error is io.EOF, we've reached the end of the data
			if err == io.EOF {
				break
			}
			// Handle other errors
			return fmt.Errorf("error reading data: %v", err)
		}

		if len(chunk) == 0 {
			break
		}

		err = destination.WriteData(chunk)
		if err != nil {
			// Handle write errors
			return fmt.Errorf("error writing data: %v", err)
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
		// You might want to log the error or perform some cleanup actions here
		return
	}

	fmt.Println("Data migration successful.")
}