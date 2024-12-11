package main

import (
	"errors"
	"fmt"
	"os"
)

func processFile(filename string) error {
	filename = "Task_422784/T2a/data.txt"
	if filename == "" {
		return errors.New("filename is required")
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Proceed with file processing
	return nil
}

func main() {
	err := processFile("")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = processFile("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File processing completed successfully")
}

// Code response one
// func main() {
// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Proceed with file operations
// }

// func someOperation() error {
// 	return errors.New("some operation failed")
// }

// func anotherOperation() error {
// 	err := someOperation()
// 	if err != nil {
// 		return fmt.Errorf("another operation failed: %w", err) // Properly wrap the error
// 	}
// 	return nil
// }

// func main() {
// 	err := anotherOperation()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

// func randomNumber(min, max int) (int, error) {
// 	if min > max {
// 		return 0, errors.New("min must be less than or equal to max")
// 	}
// 	return rand.Intn(max-min+1) + min, nil
// }

// func main() {
// 	num, err := randomNumber(10, 5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Random number:", num)
// }
