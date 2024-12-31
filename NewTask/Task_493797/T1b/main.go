package main

import (
	"fmt"
	"os"

	"github.com/gofrs/flock"
)

func main() {
	// Define the file path to work on
	filePath := "example.txt"

	// Create a new flock.Flock instance for the file
	lock := flock.New(filePath + ".lock") // Use a separate lock file

	// Attempt to acquire an exclusive lock on the file
	locked, err := lock.TryLock()
	if err != nil {
		fmt.Printf("Error acquiring lock: %v\n", err)
		return
	}

	// Ensure the lock is released when we're done
	defer lock.Unlock()

	if !locked {
		fmt.Println("Could not acquire lock, another process might be using the file.")
		return
	}

	// Open the file for reading and writing, creating it if necessary
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	// Read the current content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Modify the content (e.g., append some data)
	modifiedContent := append(content, []byte(" Modified by User1")...)

	// Write the modified content back to the file
	err = os.WriteFile(filePath, modifiedContent, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("File modified successfully by User1.")
}
