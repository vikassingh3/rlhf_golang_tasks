package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	_ "net/http/pprof" // for pprof
)

// Function to simulate event processing (example)
func processEvent(data []int) {
	// Simulating some processing
	time.Sleep(10 * time.Millisecond)
}

// Original loop with repeated slicing
func processSubslices(data []int) {
	for i := 0; i < len(data); i += 10 {
		processEvent(data[i:i+10]) // This creates a new slice each time
	}
}

// Optimized loop with pre-allocated slices
func processSubslicesOptimized(data []int, sliceSize int) {
	subSlices := make([][]int, (len(data)-1)/sliceSize+1)
	for i := range subSlices {
		start := i * sliceSize
		end := min(start+sliceSize, len(data))
		subSlices[i] = data[start:end]
	}

	// Now process all the subSlices without further allocations
	for _, subSlice := range subSlices {
		processEvent(subSlice)
	}
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Start HTTP server for pprof
func startPprofServer() {
	go func() {
		log.Println("Starting pprof server on :6060")
		log.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()
}

// Test function to run the profiling and optimization
func testProfileAndOptimize() {
	// Example data to be processed
	data := make([]int, 100000)

	// Profiling memory before starting the processing
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Before processing - Alloc: %v, TotalAlloc: %v, HeapAlloc: %v\n", m.Alloc, m.TotalAlloc, m.HeapAlloc)

	// Run the original function (to profile)
	processSubslices(data)

	// Profiling memory after processing
	runtime.ReadMemStats(&m)
	fmt.Printf("After processing original - Alloc: %v, TotalAlloc: %v, HeapAlloc: %v\n", m.Alloc, m.TotalAlloc, m.HeapAlloc)

	// Run the optimized function (to profile)
	processSubslicesOptimized(data, 10)

	// Profiling memory after optimized processing
	runtime.ReadMemStats(&m)
	fmt.Printf("After processing optimized - Alloc: %v, TotalAlloc: %v, HeapAlloc: %v\n", m.Alloc, m.TotalAlloc, m.HeapAlloc)
}

func main() {
	// Start pprof server for profiling
	startPprofServer()

	// Simulate running the test case and profiling
	testProfileAndOptimize()

	// Keep the program running to allow for pprof access
	select {}
}
