package main

import (
	"fmt"
	"sync"
	"time"
)

type AnalysisResult struct {
	NumWords int
	NumSentences int
	// Additional analysis metrics
}

func analyzeText(text string, wg *sync.WaitGroup, results *sync.Map, mutex *sync.Mutex) {
	defer wg.Done()

	// Simulate analysis work
	time.Sleep(time.Duration(100*time.Millisecond)) // Simulate analysis time
	fmt.Printf("Analyzing: %s\n", text)

	// Calculate analysis results for this chunk
	result := &AnalysisResult{
		NumWords: len(text),
		NumSentences: 1, // You can improve this sentence count logic
	}

	// Use a mutex to ensure exclusive access when updating the shared counter
	key := text[:4] // Use a unique key based on the text chunk
	mutex.Lock()
	results.Store(key, result)
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var results sync.Map
	var mutex sync.Mutex

	dataset := []string{
		"This is the first chunk of text.",
		"The second chunk of text is here.",
		"Chunk number three contains more text.",
		"Finally, the last chunk to analyze.",
	}

	for _, text := range dataset {
		wg.Add(1)
		go analyzeText(text, &wg, &results, &mutex)
	}

	wg.Wait()

	fmt.Println("Analysis Results:")
	// Iterate through the sync.Map to print the results
	results.Range(func(key, value interface{}) bool {
		result := value.(*AnalysisResult)
		fmt.Printf("Chunk %s: Words - %d, Sentences - %d\n", key, result.NumWords, result.NumSentences)
		return true // Keep iterating through the map
	})
}