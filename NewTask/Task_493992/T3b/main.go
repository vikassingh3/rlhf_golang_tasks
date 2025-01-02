package main  
import (  
    "fmt"
    "sync"
)

type AnalysisResult struct {  
    NumWords int
    NumSentences int
}  

func analyzeText(text string, resultChan chan<- AnalysisResult) {  
    // Simulate analysis work
    fmt.Printf("Analyzing: %s\n", text)

    // Calculate analysis results
    result := AnalysisResult{  
        NumWords: len(text),
        NumSentences: 1, // You can improve this sentence count logic
    }  

    // Send the result through the channel
    resultChan <- result
}  

func main() {  
    const numWorkers = 4 // Number of worker goroutines for analysis
    var wg sync.WaitGroup
    resultChan := make(chan AnalysisResult, numWorkers) // Buffered channel to hold results

    dataset := []string{  
        "This is the first chunk of text.",
        "The second chunk of text is here.",
        "Chunk number three contains more text.",
        "Finally, the last chunk to analyze.",
        // Add more chunks for testing
    }  

    // Start worker goroutines for analysis
    for i := 0; i < numWorkers; i++ {  
        wg.Add(1)
        go func() {  
            defer wg.Done()
            for text := range resultChan {  
                analyzeText(text, resultChan)
            }  
        }()
    }  

    // Send chunks to the worker goroutines
    go func() {  
        for _, text := range dataset {  
            resultChan <- text
        }  
        close(resultChan) // Close the channel to signal completion
    }()  

    // Wait for all worker goroutines to finish
    wg.Wait()  

    fmt.Println("Analysis Results:")
    // Collect and print the results from the channel
    for result := range resultChan {  
        fmt.Printf("Words - %d, Sentences - %d\n", result.NumWords, result.NumSentences)
    }  
}  