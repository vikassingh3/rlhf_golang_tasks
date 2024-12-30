package main

import (
	"fmt"
)

func main() {
	// Define a nested map: a map of strings to maps of strings to integers
	nestedMap := make(map[string]map[string]int)

	// Populate the nested map
	nestedMap["Alice"] = make(map[string]int)
	nestedMap["Alice"]["Math"] = 95
	nestedMap["Alice"]["Science"] = 88

	nestedMap["Bob"] = make(map[string]int)
	nestedMap["Bob"]["Math"] = 85
	nestedMap["Bob"]["Science"] = 92

	// Accessing values in a nested map
	fmt.Println("Alice's Math score:", nestedMap["Alice"]["Math"])
	fmt.Println("Bob's Science score:", nestedMap["Bob"]["Science"])

	// Check if a key exists in a nested map
	_, exists := nestedMap["Charlie"]
	if !exists {
		fmt.Println("Charlie does not exist in the map.")
	}

	// Iterating over a nested map
	for student, subjects := range nestedMap {
		fmt.Printf("Student: %s\n", student)
		for subject, score := range subjects {
			fmt.Printf("  %s: %d\n", subject, score)
		}
	}
}