package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// Worker function to process files
func worker(filePaths <-chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark the worker as done when it exits
	for filePath := range filePaths {
		processFile(filePath)
	}
}

// Process each file
func processFile(filePath string) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// Perform any necessary processing on the file content
	fmt.Printf("Processing file: %s\n", filePath)
	// For example, you can count the number of lines in the file
	numLines := 0
	for _, line := range string(content) {
		if line == '\n' {
			numLines++
		}
	}
	fmt.Printf("File %s contains %d lines.\n", filePath, numLines)
}

func main() {
	// Number of worker goroutines to use
	numWorkers := 4

	// Create a channel to pass file paths to worker goroutines
	filePaths := make(chan string)

	// Create a wait group to track the number of pending jobs
	var wg sync.WaitGroup

	// Ensure the directory exists
	os.MkdirAll("./example_files", 0755)

	// Launch worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(filePaths, &wg)
	}

	// Read file paths from the directory
	dirPath := "./example_files"
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking directory: ", err)
			return nil
		}
		if !info.IsDir() {
			filePaths <- path // Send the file path to the worker goroutines
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking directory: ", err)
	}

	// Close the file path channel to signal the end of work
	close(filePaths)

	// Wait for all worker goroutines to finish
	wg.Wait()

	fmt.Println("All files processed.")
}
