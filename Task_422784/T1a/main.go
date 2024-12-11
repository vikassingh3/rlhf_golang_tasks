package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return data, nil
}

func processFile(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		log.Printf("error processing file: %v", err)
		return err
	}

	fmt.Printf("File content of %s:\n%s\n", filename, string(data))
	return nil
}

func main() {

	filename := "./T1a/sample.txt"
	content := "This is a sample file content."
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalf("failed to create sample file: %v", err)
	}
	defer os.Remove(filename)

	// Process the file
	if err := processFile(filename); err != nil {
		log.Fatalf("failed to process file: %v", err)
	} else {
		fmt.Println("File processed successfully.")
	}
}
