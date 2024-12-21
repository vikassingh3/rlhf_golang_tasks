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
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	// Perform any necessary processing on the file content
	fmt.Printf("Processing file: %s\n", filePath)
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

	// Create a wait group to track the number of workers
	var wg sync.WaitGroup

	// Ensure the directory exists
	dirPath := "./example_files"
	os.MkdirAll(dirPath, 0755)

	// Add some example files (optional setup step for demonstration)
	createExampleFiles(dirPath)

	// Launch worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(filePaths, &wg)
	}

	// Read file paths from the directory and send them to the workers
	fileWalkWg := sync.WaitGroup{} // Additional WaitGroup for directory traversal
	fileWalkWg.Add(1)

	go func() {
		defer fileWalkWg.Done()
		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error walking directory %s: %v\n", path, err)
				return nil
			}
			if !info.IsDir() {
				filePaths <- path // Send the file path to the worker goroutines
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error walking directory: %v\n", err)
		}
	}()

	// Wait for the file walking to complete, then close the channel
	fileWalkWg.Wait()
	close(filePaths)

	// Wait for all worker goroutines to finish
	wg.Wait()

	fmt.Println("All files processed.")
}

// createExampleFiles creates dummy files for demonstration purposes
func createExampleFiles(dirPath string) {
	files := []struct {
		name    string
		content string
	}{
		{"example1.txt", "Line 1\nLine 2\nLine 3\n"},
		{"example2.txt", "Line A\nLine B\nLine C\n"},
		{"example3.txt", "Line X\nLine Y\nLine Z\n"},
	}

	for _, file := range files {
		err := os.WriteFile(filepath.Join(dirPath, file.name), []byte(file.content), 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", file.name, err)
		}
	}
}
