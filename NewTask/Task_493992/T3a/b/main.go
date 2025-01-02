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

func analyzeText(text string, wg *sync.WaitGroup, results chan *AnalysisResult) {
	defer wg.Done()

	// Simulate analysis work
	time.Sleep(time.Duration(100*time.Millisecond)) // Simulate analysis time
	fmt.Printf("Analyzing: %s\n", text)

	// Calculate analysis results for this chunk
	result := &AnalysisResult{
		NumWords: len(text),
		NumSentences: 1, // You can improve this sentence count logic
	}

	results <- result // Send the result through the channel
}

func main() {
	var wg sync.WaitGroup
	var results = make(chan *AnalysisResult)

	dataset := []string{
		"This is the first chunk of text.",
		"The second chunk of text is here.",
		"Chunk number three contains more text.",
		"Finally, the last chunk to analyze.",
	}

	for _, text := range dataset {
		wg.Add(1)
		go analyzeText(text, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results) // Signal that all results have been sent
	}()

	fmt.Println("Analysis Results:")
	for result := range results {
		fmt.Printf("Chunk: Words - %d, Sentences - %d\n", result.NumWords, result.NumSentences)
	}
}