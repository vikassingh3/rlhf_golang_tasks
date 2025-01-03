package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func CreateAndReturnSlice() []int {
	slice := []int{1, 2, 3, 4, 5}
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

	// Run the function for profiling
	numbers := CreateAndReturnSlice()
	fmt.Println("The slice returned is:", numbers)
} 