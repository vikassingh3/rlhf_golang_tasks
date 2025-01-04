package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"testing"
	"time"
)

// Define a struct to represent different logging strategies
type LoggingStrategy struct {
	name   string
	logger *log.Logger
}

// Create an example of a multiWriter to handle writing to multiple outputs (e.g., stdout and a file)
var multiWriter = io.MultiWriter(os.Stdout, os.Stderr)

var loggingStrategies = []LoggingStrategy{
	{
		name:   "Standard Log",
		logger: log.New(multiWriter, "", log.LstdFlags),
	},
	// You can add other logging strategies here if needed
}

func BenchmarkLogging(b *testing.B) {
	// Run benchmarks for each logging strategy
	for _, strategy := range loggingStrategies {
		strategy.logger.SetOutput(multiWriter) // This can be removed if multiWriter is already set during logger creation
		b.Run(strategy.name, func(b *testing.B) {
			startTime := time.Now()
			for i := 0; i < b.N; i++ {
				strategy.logger.Println("This is an informational message")
				strategy.logger.Printf("Error: %v\n", "An error occurred!")
				strategy.logger.Println("Warning: Something went wrong")
				strategy.logger.Println("Debug: Debug information goes here")
			}
			endTime := time.Now()

			// Calculate latency and throughput
			latency := endTime.Sub(startTime) / time.Duration(b.N)
			throughput := float64(b.N) / float64(endTime.Sub(startTime))

			// Measure resource usage
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			memoryAlloc := m.Alloc
			cpuUsage := float64(runtime.NumGoroutine())

			// Print the metrics
			fmt.Printf("Strategy: %s, Latency: %s, Throughput: %.2f messages/s, Memory Alloc: %d bytes, CPU Usage: %.2f goroutines\n",
				strategy.name, latency, throughput, memoryAlloc, cpuUsage)
		})
	}
}

func main() {
	// Example of running the benchmark
	result := testing.Benchmark(BenchmarkLogging)
	fmt.Println(result)
}
