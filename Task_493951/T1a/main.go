package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a DataPoint struct to represent the data we're retrieving.
type DataPoint struct {
	Source    string
	Statistic string
	Value     float64
}

// Simulate retrieving data from an API.
func retrieveDataFromAPI(wg *sync.WaitGroup, source string, delay time.Duration) []DataPoint {
	defer wg.Done()

	time.Sleep(delay) // Simulate some processing time.

	// Simulate returned data.
	dataPoints := []DataPoint{
		{source, "Goals Scored", 32.5},
		{source, "Assists", 20.0},
	}

	return dataPoints
}

func main() {
	var wg sync.WaitGroup
	var allData []DataPoint

	// Simulate multiple data sources.
	sources := []string{"API1", "API2", "API3"}

	// Launch concurrent data retrieval tasks.
	for _, source := range sources {
		wg.Add(1) // Increment the WaitGroup counter.
		go retrieveDataFromAPI(&wg, source, time.Duration(rand.Intn(3)+1)*time.Second)
	}

	// Wait for all data retrieval tasks to complete.
	wg.Wait()

	fmt.Println("Aggregated Data:")
	for _, dataPoint := range allData {
		fmt.Printf("Source: %s, Statistic: %s, Value: %.2f\n", dataPoint.Source, dataPoint.Statistic, dataPoint.Value)
	}
}