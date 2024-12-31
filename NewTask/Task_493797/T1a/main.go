package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func lockFile(file *os.File) error {
	// Acquire an exclusive lock
	return unix.Flock(int(file.Fd()), unix.LOCK_EX)
}

func unlockFile(file *os.File) error {
	// Release the lock
	return unix.Flock(int(file.Fd()), unix.LOCK_UN)
}

func editFile(filename string, content string) error {
	// Open the file for reading and writing, creating it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Acquire a mandatory lock
	if err := lockFile(file); err != nil {
		return fmt.Errorf("error acquiring lock: %v", err)
	}
	defer unlockFile(file)

	// Write the new content to the file
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(content); err != nil {
		return err
	}
	writer.Flush()

	return nil
}

func main() {
	filename := "example.txt"
	content := "Hello, World!\n"

	if err := editFile(filename, content); err != nil {
		log.Fatal(err)
	}

	fmt.Println("File edited successfully.")
}
