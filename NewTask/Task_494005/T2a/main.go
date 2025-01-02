package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Case 1: Sequential access
	numbersSequential := make([]int, 1000000)
	for i := range numbersSequential {
		numbersSequential[i] = rand.Int()
	}

	startSequential := time.Now()
	var sumSequential int
	for _, value := range numbersSequential {
		sumSequential += value
	}
	elapsedSequential := time.Since(startSequential)
	fmt.Println("Sequential access: Sum:", sumSequential, "Time taken:", elapsedSequential)

	// Case 2: Random access
	numbersRandom := make([]int, 1000000)
	for i := range numbersRandom {
		numbersRandom[i] = rand.Int()
	}

	startRandom := time.Now()
	sumRandom := 0
	randomIndices := make([]int, 1000000)
	for i := range randomIndices {
		randomIndices[i] = rand.Intn(len(numbersRandom))
	}
	for _, index := range randomIndices {
		sumRandom += numbersRandom[index]
	}
	elapsedRandom := time.Since(startRandom)
	fmt.Println("Random access: Sum:", sumRandom, "Time taken:", elapsedRandom)
}