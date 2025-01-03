package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func CreateAndReturnSlice() []int {
	slice := make([]int, 10000) // Increased the slice size for illustrative purposes
	for i := range slice {
		slice[i] = i * i
	}
	return slice
}

func createAndReturnSimplifiedSlice() []int {
	var slice []int
	for i := 0; i < 10000; i++ {
		slice = append(slice, i*i)
	}
	return slice
}


func main() {
	// Create a file to save the profile data
	f, err := os.Create("profile.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start the CPU profile
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// Generate some substantial workload
	fmt.Println("Running workload...")
	time.Sleep(1 * time.Second) // Allow time to create substantial profile
	for i := 0; i < 100000; i++ {
		CreateAndReturnSlice()
	}
	fmt.Println("Workload completed.")
}