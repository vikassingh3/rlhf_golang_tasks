package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// Open the file in read-only mode
	file, err := os.Open("./large_data_file.bin")
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

	// Initialize the sum variable
	var sum int64

	// Loop through the memory-mapped file and aggregate the sum
	for i := 0; i <= len(mmap)-8; i += 8 { // Ensure there's enough data left to read 8 bytes
		val := binary.LittleEndian.Uint64(mmap[i : i+8])
		sum += int64(val)
	}

	// Print the sum
	fmt.Println("Sum:", sum)
}
