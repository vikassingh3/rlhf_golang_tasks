package main

import (
	"fmt"
	"sync"
)

// Function to perform typographic analysis on a given text chunk.
func analyzeText(text string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done.

	// Simulate analysis work
	fmt.Printf("Analyzing: %s\n", text)
	// sleep(2) // Uncomment to simulate analysis time
}

func main() {
	// Create a WaitGroup to track the number of goroutines
	var wg sync.WaitGroup

	// Large dataset - we'll split it into smaller chunks for analysis
	dataset := []string{
		"This is the first chunk of text.",
		"The second chunk of text is here.",
		"Chunk number three contains more text.",
		"Finally, the last chunk to analyze.",
	}

	// For each chunk, create a new goroutine that uses the WaitGroup
	for _, text := range dataset {
		wg.Add(1) // Increment the WaitGroup count
		go analyzeText(text, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All analysis tasks have completed.")
}