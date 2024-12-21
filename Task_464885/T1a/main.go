package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	// Define the data retention period in hours
	retentionPeriod = 24
)

func main() {
	// Define the file path
	filePath := "example.txt"

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If the file doesn't exist, create it
		createFile(filePath)
	}

	// Process the file
	processFile(filePath)

	// Implement data retention
	enforceDataRetention(filePath)
}

// createFile creates a new file if it doesn't exist
func createFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()
}

// processFile reads the file, processes its content, and appends new data
func processFile(filePath string) {
	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println("Current file content:")
	fmt.Printf("%s\n", string(content))

	// Append new data to the file
	newData := "New data added at " + time.Now().Format(time.RFC3339) + "\n"
	err = ioutil.WriteFile(filePath, []byte(string(content)+newData), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Println("New data added to the file.")
}

// enforceDataRetention deletes the file if it is older than the retention period
func enforceDataRetention(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}

	// Calculate the time the file was created
	fileModTime := fileInfo.ModTime()

	// Calculate the current time
	currentTime := time.Now()

	// Calculate the age of the file
	age := currentTime.Sub(fileModTime)

	// Check if the file is older than the retention period
	if age.Hours() > float64(retentionPeriod) {
		err = os.Remove(filePath)
		if err != nil {
			log.Fatalf("Error deleting file: %v", err)
		}
		fmt.Println("File has been deleted due to age.")
	} else {
		fmt.Println("File is within retention period.")
	}
}