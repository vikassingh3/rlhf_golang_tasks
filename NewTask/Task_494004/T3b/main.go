package main

import (
	"fmt"
	"sort"
	"time"
)

// BenchmarkResult represents the results of a single benchmark run.
type BenchmarkResult struct {
	Strategy     string
	Latency      time.Duration
	Throughput   float64
	MemoryAlloc  uint64
	CPUUsage     float64
}

// Dummy function to simulate benchmark data collection (this should be replaced with actual benchmark results)
func collectBenchmarkResults() []BenchmarkResult {
	return []BenchmarkResult{
		{"StrategyA", 10 * time.Millisecond, 1000.0, 50000, 0.5},
		{"StrategyB", 15 * time.Millisecond, 900.0, 60000, 0.6},
		{"StrategyC", 12 * time.Millisecond, 1100.0, 55000, 0.55},
	}
}

func analyzeBenchmarkResults(results []BenchmarkResult) {
	// Sort results by latency in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Latency > results[j].Latency
	})

	fmt.Println("**Performance Bottleneck Analysis:**")
	fmt.Println("**High Latency Strategies:**")
	for _, result := range results[:3] { // Show the top 3 highest latency strategies
		fmt.Printf("- %s: Latency: %s\n", result.Strategy, result.Latency)
	}

	fmt.Println("\n**Low Throughput Strategies:**")
	sort.Slice(results, func(i, j int) bool {
		return results[i].Throughput < results[j].Throughput
	})
	for _, result := range results[:3] { // Show the top 3 lowest throughput strategies
		fmt.Printf("- %s: Throughput: %.2f messages/s\n", result.Strategy, result.Throughput)
	}

	fmt.Println("\n**High Resource Usage Strategies:**")
	sort.Slice(results, func(i, j int) bool {
		return results[i].MemoryAlloc > results[j].MemoryAlloc
	})
	for _, result := range results[:3] { // Show the top 3 strategies with high memory allocation
		fmt.Printf("- %s: Memory Alloc: %d bytes\n", result.Strategy, result.MemoryAlloc)
	}
}

func main() {
	results := collectBenchmarkResults()
	analyzeBenchmarkResults(results)
}
