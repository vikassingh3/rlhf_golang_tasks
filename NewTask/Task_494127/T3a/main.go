package main

import (
	"fmt"
	"sync"
)

// Flattened structure to replace nested maps
type StudentScores struct {
	Math   int `json:"math"`
	Science int `json:"science"`
}

type SchoolScores struct {
	Alice   StudentScores `json:"alice"`
	Bob     StudentScores `json:"bob"`
	Charlie StudentScores `json:"charlie"`
}

func main() {
	// Initialize the school scores with default values
	schoolScores := SchoolScores{
		Alice:   StudentScores{Math: 0, Science: 0},
		Bob:     StudentScores{Math: 0, Science: 0},
		Charlie: StudentScores{Math: 0, Science: 0},
	}

	// Update scores
	schoolScores.Alice.Math = 95
	schoolScores.Alice.Science = 88
	schoolScores.Bob.Math = 85
	schoolScores.Bob.Science = 92
	schoolScores.Charlie.Math = 78

	// Access scores
	fmt.Println("Alice's Math score:", schoolScores.Alice.Math)
	fmt.Println("Bob's Science score:", schoolScores.Bob.Science)

	// Synchronized access for concurrent updates
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		schoolScores.Alice.Math = 97
	}()

	go func() {
		defer wg.Done()
		schoolScores.Bob.Science = 94
	}()

	wg.Wait()

	fmt.Println("Updated Alice's Math score:", schoolScores.Alice.Math)
	fmt.Println("Updated Bob's Science score:", schoolScores.Bob.Science)
}