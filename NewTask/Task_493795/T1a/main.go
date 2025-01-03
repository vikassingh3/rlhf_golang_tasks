package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// Open the file for reading
	file, err := os.Open("./example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Get the file's size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	// Map the file into memory
	data, err := syscall.Mmap(int(file.Fd()), 0, int(fileInfo.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println("Error mapping file:", err)
		return
	}
	defer syscall.Munmap(data) // Unmap the file from memory

	// Convert the slice of bytes to a string
	fileContent := string(data) // Direct conversion of the byte slice to a string

	// Print the file content
	fmt.Println("File content:", fileContent)
}
