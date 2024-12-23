package main

import (
	"fmt"
)

func main() {
	// Create an empty map
	users := make(map[string]int)

	// Store data in the map
	users["Alice"] = 30
	users["Bob"] = 25
	users["Charlie"] = 35

	// Retrieve data from the map
	aliceAge := users["Alice"]
	bobAge := users["Bob"]

	// Output the data
	fmt.Println("Alice's age:", aliceAge)
	fmt.Println("Bob's age:", bobAge)

	// Check if a key exists
	if _, exists := users["David"]; !exists {
		fmt.Println("David is not in the map.")
	}

	// Delete a key-value pair
	delete(users, "Bob")

	// Output the map after deletion
	fmt.Println("Users after deleting Bob:", users)
}