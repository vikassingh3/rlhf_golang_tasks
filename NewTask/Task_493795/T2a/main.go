package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
	"strings" // Add this import
)

func main() {
	// Open the file
	file, err := os.Open("./large_log_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get file info
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fi.Size()

	// Memory map the file
	mmap, err := syscall.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Munmap(mmap)

	// Initialize variables
	keyword := "error"
	count := 0
	reader := bufio.NewReader(bytes.NewReader(mmap))

	// Read the memory-mapped file line by line
	for {
		line, err := reader.ReadString('\n') // Only capture 2 return values
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Check if the line contains the keyword
		if strings.Contains(line, keyword) {
			count++
		}
	}

	// Print the result
	fmt.Printf("Occurrences of '%s': %d\n", keyword, count)
}
