package main

import (
    "fmt"
    "sync"
)

// FashionTrend represents a fashion trend for a specific season.
type FashionTrend struct {
    Season string
    Trend   string
}

// analyzeSeason analyzes fashion trends for a given season and updates the result channel.
func analyzeSeason(season string, trends []FashionTrend, result chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    // Your season-specific analysis logic goes here
    // For example, you could count the frequency of different trends.
    // For demonstration purposes, we'll just count the trends.
    trendCount := len(trends)
    fmt.Printf("%s: Analyzed %d trends\n", season, trendCount)

    // Send the result back to the main goroutine
    result <- fmt.Sprintf("%s: %d trends", season, trendCount)
}

func main() {
    trendsData := []FashionTrend{
        {"Spring", "Pastel Colors"},
        {"Spring", "Printed Shirts"},
        {"Summer", "Lightweight Fabrics"},
        {"Summer", "Breathing Skirts"},
        {"Fall", "Burgundy Colors"},
        {"Fall", "Knitted Sweaters"},
        {"Winter", "Cozy Sweaters"},
        {"Winter", "Black Boots"},
    }

    // Create a map to group trends by season
    trendsBySeason := make(map[string][]FashionTrend)
    for _, trend := range trendsData {
        trendsBySeason[trend.Season] = append(trendsBySeason[trend.Season], trend)
    }

    // Create a channel to receive the results from the goroutines
    resultChannel := make(chan string)

    // Create a WaitGroup to wait for all goroutines to complete
    var wg sync.WaitGroup

    // Launch a goroutine for each season to analyze trends
    for season, trends := range trendsBySeason {
        wg.Add(1)
        go analyzeSeason(season, trends, resultChannel, &wg)
    }

    // Start a goroutine to collect the results from the channels
    go func() {
        for result := range resultChannel {
            fmt.Println("Result:", result)
        }
    }()

    // Wait for all goroutines to complete their work
    wg.Wait()

    // Close the result channel to signal the collector goroutine to stop
    close(resultChannel)

    fmt.Println("All analysis complete.")
}