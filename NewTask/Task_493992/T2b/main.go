package main  
import (  
    "fmt"
    "sync"
)

type AnalysisResult struct {
    // Define the structure of analysis results here
    NumWords int
    NumSentences int
    // Add more analysis metrics as needed
}

func analyzeText(text string, wg *sync.WaitGroup, results *sync.Map) {
    defer wg.Done()

    // Simulate analysis work
    fmt.Printf("Analyzing: %s\n", text)

    // Calculate analysis results for this chunk
    result := &AnalysisResult{
        NumWords: len(text),
        NumSentences: 1, // You can improve this sentence count logic
    }

    // Store the result using the sync.Map
    key := text[:4] // Use a unique key based on the text chunk or any other appropriate criteria
    results.Store(key, result)
}

func main() {  
    var wg sync.WaitGroup
    var results sync.Map

    dataset := []string{  
        "This is the first chunk of text.",
        "The second chunk of text is here.",
        "Chunk number three contains more text.",
        "Finally, the last chunk to analyze.",
    }

    for _, text := range dataset {  
        wg.Add(1)
        go analyzeText(text, &wg, &results)
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