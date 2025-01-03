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

	// Sum up the values in the file
	var sum int32
	for i := 0; i < len(mmap); i += 4 {
		val := binary.LittleEndian.Uint32(mmap[i : i+4])
		sum += int32(val)
	}

	// Print the sum
	fmt.Println("Sum:", sum)
}
