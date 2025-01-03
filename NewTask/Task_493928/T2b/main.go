package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type AgriculturalData struct {
	SensorID    int
	Timestamp   time.Time
	Temperature float64
	Humidity    float64
}

func main() {
	// Generate some sample agricultural data
	data := generateSampleData(10)

	// Process the data concurrently using multiple goroutines
	var wg sync.WaitGroup
	numWorkers := 4
	dataCh := make(chan AgriculturalData, numWorkers)

	// Start worker goroutines
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, dataCh, &wg)
	}

	// Feed the data to the worker goroutines
	for _, d := range data {
		dataCh <- d
	}
	close(dataCh) // Close the channel after all data is sent

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("Data processing completed.")
}

func worker(id int, dataCh <-chan AgriculturalData, wg *sync.WaitGroup) {
	defer wg.Done() // Signal the WaitGroup that this worker is done

	for data := range dataCh { // This loop will exit once the channel is closed and all data is processed
		// Start timing the data processing for latency measurement
		start := time.Now()

		// Simulate processing with a potential bottleneck (e.g., database or network call)
		processData(data)

		// Measure the latency of the processing (simulating network/database delay)
		latency := time.Since(start)

		// Print out the processing time (latency)
		fmt.Printf("Worker %d: Processed data in %s\n", id, latency)
	}
}

func processData(data AgriculturalData) {
	// Simulate a processing task (e.g., a complex database query or I/O operation)
	// This is where real bottlenecks might occur in real-world scenarios.
	// In this case, we're just performing some calculations.
	data.Temperature += 10
	data.Humidity += 20

	// Simulate external API call or database call latency by introducing a delay
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond) // Simulating a realistic latency
}

func generateSampleData(count int) []AgriculturalData {
	data := make([]AgriculturalData, count)
	for i := 0; i < count; i++ {
		data[i] = AgriculturalData{
			SensorID:    rand.Intn(100),
			Timestamp:   time.Now(),
			Temperature: rand.Float64() * 50,
			Humidity:    rand.Float64() * 100,
		}
	}
	return data
}
