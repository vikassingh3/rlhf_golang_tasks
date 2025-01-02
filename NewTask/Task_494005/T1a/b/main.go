package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize a large slice
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = rand.Intn(10000)
	}

	// Start CPU profiling
	runtime.GC()
	go func() {
		time.Sleep(2 * time.Second)
		runtime.GC()
		cpuProfFile, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatalf("could not create CPU profile: %v", err)
		}
		pprof.StartCPUProfile(cpuProfFile)
		defer pprof.StopCPUProfile()
	}()

	// Perform a computationally expensive operation
	processSlice(numbers)
}

func processSlice(numbers []int) {
	start := time.Now()
	sum := 0
	for _, value := range numbers {
		sum += value * value
	}
	fmt.Printf("Sum of squares: %d, Time taken: %v\n", sum, time.Since(start))
}