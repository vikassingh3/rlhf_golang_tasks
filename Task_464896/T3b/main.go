package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Create a file for the CPU profile
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// Your application code goes here
	doSomeWork()
}

func doSomeWork() {
	// This function represents the action you want to profile
	time.Sleep(2 * time.Second) // Simulate work
	// Add your actual work here
}
