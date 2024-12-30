package main

import (
	"fmt"
	"os"
)

func main() {
	// Open a file for writing
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close() // Ensure the file is closed even if an error occurs

	// Write some data to the file
	_, err = file.WriteString("Hello, world!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Data written successfully.")
	}
}