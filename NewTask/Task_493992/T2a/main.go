package main

import (
	"fmt"
	"sync"
)

// A shared word count structure
type WordCount struct {
	Words map[string]int
	mu    sync.RWMutex
}

// Increment word count in a thread-safe manner
func (wc *WordCount) IncWordCount(word string) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	if _, ok := wc.Words[word]; !ok {
		wc.Words[word] = 0
	}
	wc.Words[word]++
}

// Calculate total word count in a thread-safe manner
func (wc *WordCount) TotalWordCount() int {
	var total int
	wc.mu.RLock()
	defer wc.mu.RUnlock()
	for _, count := range wc.Words {
		total += count
	}
	return total
}

// Function to perform a portion of the text analysis using a shared WordCount
func analyzeText(text string, wc *WordCount, wg *sync.WaitGroup) {
	defer wg.Done()

	// Split text into words for simplicity (in reality, you may need a more sophisticated text processor)
	words := make([]string, 0, 100)
	for _, w := range text {
		if w != ' ' {
			words = append(words, string(w))
		}
	}

	// Increment word counts
	for _, word := range words {
		wc.IncWordCount(word)
	}

	fmt.Printf("Analyzing: %s\n", text)
}

func main() {
	// Create a WaitGroup to track the number of goroutines
	var wg sync.WaitGroup

	// Initialize a shared WordCount structure
	wordCount := &WordCount{Words: make(map[string]int)}

	// Large dataset for text processing
	dataset := []string{
		"This is the first text",
		"Text processing is",
		"Performing amazing tasks",
		"Go applications are",
		"Scalable and fast!",
	}

	// For each text, create a new goroutine using the WaitGroup
	for _, text := range dataset {
		wg.Add(1)
		go analyzeText(text, wordCount, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the word count results
	fmt.Printf("Total word count: %d\n", wordCount.TotalWordCount())
	for word, count := range wordCount.Words {
		fmt.Printf("%s: %d\n", word, count)
	}
}