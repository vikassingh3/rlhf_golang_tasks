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
	file, err := os.Open("large_data_file.bin")
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

	// Define a buffer to process the file in chunks
	chunkSize := 4096 // 4 KB chunk
	buf := make([]byte, chunkSize)

	// Sum up the values in the file using chunks
	var sum int32
	for i := 0; i < len(mmap); i += chunkSize {
		end := i + chunkSize
		if end > len(mmap) {
			end = len(mmap)
		}
		copy(buf[:], mmap[i:end])
		for j := 0; j < len(buf); j += 4 {
			if j+4 <= len(buf) {
				val := binary.LittleEndian.Uint32(buf[j : j+4])
				sum += int32(val)
			}
		}
	}

	// Print the sum
	fmt.Println("Sum:", sum)
}
