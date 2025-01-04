package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sort"
	"testing"
	"time"
)

// Define a struct for the logging strategy with necessary performance metrics
type LoggingStrategy struct {
	name       string
	logger     *log.Logger
	latency    time.Duration
	throughput float64
	memoryAlloc uint64
	cpuUsage    float64
}

// Example of a multiWriter to handle writing to multiple outputs (stdout and stderr)
var multiWriter = io.MultiWriter(os.Stdout, os.Stderr)

// Create a list of logging strategies
var loggingStrategies = []LoggingStrategy{
	{
		name:   "Standard Log",
		logger: log.New(multiWriter, "", log.LstdFlags),
	},
	// You can add more logging strategies as needed
}

// BenchmarkLogging runs the benchmark and collects performance metrics
func BenchmarkLogging(b *testing.B) {
	// Example of logging with metrics collection
	for _, strategy := range loggingStrategies {
		strategy.logger.SetOutput(multiWriter)
		b.Run(strategy.name, func(b *testing.B) {
			startTime := time.Now()

			// Benchmark log generation in the loop
			for i := 0; i < b.N; i++ {
				strategy.logger.Println("This is an informational message")
				strategy.logger.Printf("Error: %v\n", "An error occurred!")
				strategy.logger.Println("Warning: Something went wrong")
				strategy.logger.Println("Debug: Debug information goes here")
			}

			endTime := time.Now()

			// Calculate latency and throughput
			strategy.latency = endTime.Sub(startTime) / time.Duration(b.N)
			strategy.throughput = float64(b.N) / float64(endTime.Sub(startTime))

			// Measure resource usage
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			strategy.memoryAlloc = m.Alloc
			strategy.cpuUsage = float64(runtime.NumGoroutine())
		})
	}
}

// BenchmarkResult stores the benchmark results for analysis
type BenchmarkResult struct {
	Strategy   string
	Latency    time.Duration
	Throughput float64
	MemoryAlloc uint64
	CPUUsage    float64
}

// Collect and store benchmark results
func collectResults() []BenchmarkResult {
	results := make([]BenchmarkResult, 0, len(loggingStrategies))

	for _, logger := range loggingStrategies {
		result := BenchmarkResult{
			Strategy:    logger.name,
			Latency:     logger.latency,
			Throughput:  logger.throughput,
			MemoryAlloc: logger.memoryAlloc,
			CPUUsage:    logger.cpuUsage,
		}
		results = append(results, result)
	}

	return results
}

// Sort the results by throughput
func sortResultsByThroughput(results []BenchmarkResult) {
	// Sorting in descending order of throughput
	sort.Slice(results, func(i, j int) bool {
		return results[i].Throughput > results[j].Throughput
	})
}

// Analyze the benchmark results to identify bottlenecks
func analyzeResults() {
	results := collectResults()
	sortResultsByThroughput(results)

	// Print sorted results
	fmt.Printf("\nSorted Benchmark Results by Throughput:\n")
	for _, result := range results {
		fmt.Printf("Strategy: %s, Latency: %s, Throughput: %.2f messages/s, Memory Alloc: %d bytes, CPU Usage: %.2f goroutines\n",
			result.Strategy, result.Latency, result.Throughput, result.MemoryAlloc, result.CPUUsage)
	}

	// Identify potential bottlenecks
	fmt.Printf("\nPotential Bottlenecks:\n")
	for i := 1; i < len(results); i++ {
		// Check for significant throughput drop (10% or more)
		if results[i].Throughput < 0.9*results[i-1].Throughput {
			fmt.Printf("%s has significantly lower throughput.\n", results[i].Strategy)
		}

		// Check for significant latency increase (10% or more)
		if results[i].Latency > 1 *results[i-1].Latency {
			fmt.Printf("%s has significantly higher latency.\n", results[i].Strategy)
		}

		// Check for resource usage anomalies
		if results[i].MemoryAlloc > results[i-1].MemoryAlloc*2 {
			fmt.Printf("%s is using significantly more memory than the previous strategy.\n", results[i].Strategy)
		}
		if results[i].CPUUsage > results[i-1].CPUUsage*2 {
			fmt.Printf("%s is using significantly more CPU than the previous strategy.\n", results[i].Strategy)
		}
	}
}

func main() {
	// Run the benchmark
	result := testing.Benchmark(BenchmarkLogging)
	fmt.Println(result)

	// Analyze and print the results
	analyzeResults()
}
