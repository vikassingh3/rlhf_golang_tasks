package main  
import (  
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Optimize processEvent by doing a bigger chunk of work per event 
func processEvent(events []int) {  
    for i, event := range events {
        // Simulate some processing work by calculating the sum of even numbers from 0 to 10000
        total := 0
        for j := 0; j < 10000; j++ {
            if j%2 == 0 {
                total += j
            }
        }
        events[i] = event * total
    }
    fmt.Println("Processed events...")
}  

func main() {  
    // Set GOMAXPROCS to number of CPUs for better performance
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    const numEvents = 1000000
    const batchSize = 10000
    
    // Start a timer to measure total processing time
    start := time.Now()
    
    // Use a pool of goroutines to process events in batches
    var wg sync.WaitGroup
    wg.Add(numEvents / batchSize)
    eventChannel := make(chan []int, 100) // Buffered channel to reduce blocking
    
    go func() {
        for i := 0; i < numEvents; i += batchSize {
            // Create batches of events
            end := i + batchSize - 1
            if end >= numEvents {
                end = numEvents - 1
            }
            eventBatch := make([]int, end-i+1)
            for j := i; j <= end; j++ {
                eventBatch[j-i] = j // Simulate random event data
            }
            eventChannel <- eventBatch
        }
        close(eventChannel)
    }()

    // Processor goroutines
    for i := 0; i < 10; i++ {
        go func() {
            for events := range eventChannel {
                processEvent(events)
                wg.Done()
            }
        }()
    }

    wg.Wait()

    elapsed := time.Since(start)
    fmt.Printf("Took %s to process %d events.\n", elapsed, numEvents)
} 
