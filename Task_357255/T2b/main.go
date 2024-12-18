package main

import (
	"fmt"
	"os"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	apiKey := os.Getenv("API_KEY")
	if databaseURL == "" || apiKey == "" {
		fmt.Println("Error: Missing environment variables")
		os.Exit(1)
	}

	fmt.Printf("Running with DATABASE_URL: %s and API_KEY: %s\n", databaseURL, apiKey)
}
