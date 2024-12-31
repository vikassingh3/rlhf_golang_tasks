package main

import (
	"fmt"
	"sync"
)

type fashionTrend struct {
	Color  string
	Style  string
	Count int
}

var fashionTrends []fashionTrend
var mutex sync.Mutex

// Function to count trending colors
func countTrendingColors(data []fashionTrend, wg *sync.WaitGroup) {
	defer wg.Done()

	colorCounts := make(map[string]int)

	// Process each trend to count colors
	for _, trend := range data {
		colorCounts[trend.Color]++
	}

	// Protect shared resource with a mutex
	mutex.Lock()
	defer mutex.Unlock()

	// Add color counts to the global results
	for color, count := range colorCounts {
		fashionTrends = append(fashionTrends, fashionTrend{
			Color:  color,
			Style:  "All Styles",
			Count:  count,
		})
	}
}

// Function to count popular styles
func countPopularStyles(data []fashionTrend, wg *sync.WaitGroup) {
	defer wg.Done()

	styleCounts := make(map[string]int)

	// Process each trend to count styles
	for _, trend := range data {
		styleCounts[trend.Style]++
	}

	// Protect shared resource with a mutex
	mutex.Lock()
	defer mutex.Unlock()

	// Add style counts to the global results
	for style, count := range styleCounts {
		fashionTrends = append(fashionTrends, fashionTrend{
			Color:  "All Colors",
			Style:  style,
			Count:  count,
		})
	}
}

func main() {
	// Create a large dataset of fashion trends
	fashionTrendsData := []fashionTrend{
		{Color: "Red", Style: "Casual", Count: 10},
		{Color: "Blue", Style: "Formal", Count: 15},
		{Color: "Red", Style: "Formal", Count: 20},
		{Color: "Green", Style: "Casual", Count: 25},
		{Color: "Blue", Style: "Casual", Count: 30},
		{Color: "Black", Style: "Formal", Count: 35},
		{Color: "White", Style: "Casual", Count: 40},
		{Color: "Yellow", Style: "Formal", Count: 45},
		// Add more data as needed
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Launch goroutines for processing different metrics
	wg.Add(1)
	go countTrendingColors(fashionTrendsData, &wg)

	wg.Add(1)
	go countPopularStyles(fashionTrendsData, &wg)

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the results
	fmt.Println("Fashion Trends Analysis Complete:")
	for _, trend := range fashionTrends {
		fmt.Printf("Color: %v, Style: %v, Count: %v\n", trend.Color, trend.Style, trend.Count)
	}
}