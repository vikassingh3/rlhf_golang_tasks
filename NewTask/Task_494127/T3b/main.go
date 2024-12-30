package main

import (
	"fmt"
	"sync"
)

type StudentScores struct {
	Math   int
	Science int
}

type SchoolScores struct {
	sync.Mutex
	Students map[string]StudentScores
}

func main() {
	schoolScores := SchoolScores{Students: make(map[string]StudentScores)}

	// Initialize and update scores
	schoolScores.Lock()
	schoolScores.Students["Alice"] = StudentScores{Math: 95, Science: 88}
	schoolScores.Students["Bob"] = StudentScores{Math: 85, Science: 92}
	schoolScores.Unlock()

	// Access and update scores concurrently
	go func() {
		schoolScores.Lock()
		aliceScores := schoolScores.Students["Alice"]
		aliceScores.Math += 10
		schoolScores.Students["Alice"] = aliceScores
		schoolScores.Unlock()
	}()

	go func() {
		schoolScores.Lock()
		bobScores := schoolScores.Students["Bob"]
		bobScores.Science += 5
		schoolScores.Students["Bob"] = bobScores
		schoolScores.Unlock()
	}()

	// Wait for concurrent updates to finish
	// (In a real application, use appropriate synchronization mechanisms)
	<-make(chan struct{})
	<-make(chan struct{})

	// Access scores
	schoolScores.Lock()
	fmt.Println("Alice's Math score:", schoolScores.Students["Alice"].Math)
	fmt.Println("Bob's Science score:", schoolScores.Students["Bob"].Science)
	schoolScores.Unlock()
}