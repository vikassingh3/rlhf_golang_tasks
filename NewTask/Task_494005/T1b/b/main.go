package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Initialize a large slice
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = rand.Intn(10000)
	}

	// Start CPU profiling
	cpuProfFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("Error creating CPU profile: %v\n", err)
		return
	}
	defer cpuProfFile.Close()

	pprof.StartCPUProfile(cpuProfFile)
	defer pprof.StopCPUProfile()

	// Perform computation
	start := time.Now()
	sum := 0
	for _, value := range numbers {
		sum += value * value
	}
	fmt.Printf("Sum of squares: %d, Time taken: %v\n", sum, time.Since(start))
}
