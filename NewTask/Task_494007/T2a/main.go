package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// Start memory profiling
	f, err = os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal(err)
	}

	// Your code to be profiled goes here
	time.Sleep(10 * time.Second) // Simulate work
}