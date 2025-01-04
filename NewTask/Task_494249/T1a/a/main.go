package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Step 2: Open a file for logging
	file, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Step 3: Create a logger with the file as output
	logger := log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Step 4: Log messages
	logger.Println("This is a log message.")
	logger.Printf("Formatted log message: %s", "Hello, world!")
	
	// Log without causing panic
	logger.Println("This is a log message, not a panic.") // Replace Panicln with Println
}
