package main

import (
	"fmt"
	"strings"
	"sync"
)

// Replace this with your desired typographic analysis logic
func typographicAnalysis(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Your typographic analysis code here
	// For simplicity, let's just count the words in the text for this example
	words := strings.Fields(text)
	fmt.Printf("Word Count for chunk: %q - %d words\n", text, len(words))
}

func main() {
	data := []string{
		"The quick brown fox jumps over the lazy dog",
		"Another sentence for testing purpose",
		"More data goes here",
		// Add more sentences for your dataset
	}

	const chunkSize = 2 // Size of each chunk to process concurrently

	var wg sync.WaitGroup

	// Split the dataset into smaller chunks and process them concurrently
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		wg.Add(1)

		// Combine the chunk into a single string
		combinedText := strings.Join(chunk, " ")

		// Start a new goroutine for each chunk
		go typographicAnalysis(combinedText, &wg)
	}

	// Wait for all goroutines to finish processing
	wg.Wait()

	fmt.Println("Typographic analysis completed.")
}
