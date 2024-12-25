package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// Define a custom error type
type FileError struct {
	Message string
}

func (e *FileError) Error() string {
	return e.Message
}

// Function to read a file
func ReadFile(filename string) (*[]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, &FileError{Message: fmt.Sprintf("Error reading file: %v", err)}
	}
	return &data, nil
}

// Function to process file data
func ProcessFileData(data []byte) error {
	// Simulate some processing logic
	// For this example, let's assume there's a failure if the data length is zero
	if len(data) == 0 {
		return errors.New("File data is empty")
	}
	// Add more processing logic here
	return nil
}

func main() {
	filename := "example.txt"

	// Read the file
	data, err := ReadFile(filename)
	if err != nil {
		// Handle the error from ReadFile
		fmt.Println("Error reading file:", err)
		return
	}

	// Process the file data
	err = ProcessFileData(*data)
	if err != nil {
		// Handle the error from ProcessFileData
		fmt.Println("Error processing file data:", err)
		return
	}

	// If no errors occurred, process the data
	fmt.Println("File data processed successfully!")
}