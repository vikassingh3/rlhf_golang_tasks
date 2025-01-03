package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// Open the file in read-only mode
	file, err := os.Open("large_data_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get the file size
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

	// Process the data as needed
	for i := 0; i < len(mmap); i++ {
		// Example of reading the byte value
		fmt.Printf("Byte %d: %v\n", i, mmap[i])
	}
}
