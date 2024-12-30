package main

import (
	"fmt"
	"sync"
)

func main() {
	// Use a struct to represent a student's scores
	type StudentScores struct {
		Math    int
		Science int
	}

	// Use a map of strings to StudentScores to store data
	var studentScores map[string]StudentScores
	var wg sync.WaitGroup

	// Initialize the map
	studentScores = make(map[string]StudentScores)

	// Function to add student scores
	addStudentScores := func(name string, math, science int) {
		defer wg.Done()
		studentScores[name] = StudentScores{Math: math, Science: science}
	}

	// Add multiple students in parallel
	wg.Add(2)
	go addStudentScores("Alice", 95, 88)
	go addStudentScores("Bob", 85, 92)

	wg.Wait()

	// Access and print scores
	fmt.Println("Alice's Math score:", studentScores["Alice"].Math)
	fmt.Println("Bob's Science score:", studentScores["Bob"].Science)

	// Remove a student
	delete(studentScores, "Alice")
	fmt.Println("Alice removed. StudentScores:", studentScores)

	// Iterate over the remaining students
	for student, scores := range studentScores {
		fmt.Printf("Student: %s\n", student)
		fmt.Printf("  %s: %d\n  %s: %d\n", "Math", scores.Math, "Science", scores.Science)
	}
}