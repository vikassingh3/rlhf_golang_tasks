package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"
)

func CreateAndReturnSlice() []int {
	slice := []int{1, 2, 3, 4, 5}
	return slice
}

func main() {
	go func() {
		log.Println("Starting HTTP server for pprof...")
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	defer func() {
		log.Println("Stopping pprof server...")
		time.Sleep(1 * time.Second)
	}()

	numbers := CreateAndReturnSlice()
	fmt.Println("The slice returned is:", numbers)

	// Enable CPU profiling
	cpuProfFile := "cpu.prof"
	cpuProf, err := os.Create(cpuProfFile)
	if err != nil {
		log.Fatalf("Failed to open file for CPU profiling: %v", err)
	}
	defer cpuProf.Close()

	if err := pprof.StartCPUProfile(cpuProf); err != nil {
		log.Fatalf("Failed to start CPU profiling: %v", err)
	}

	time.Sleep(5 * time.Second) // Perform the operation for 5 seconds

	pprof.StopCPUProfile() // Corrected: StopCPUProfile does not return a value

	log.Printf("CPU profile saved to %s. Analyze it using 'go tool pprof'.\n", cpuProfFile)
}
