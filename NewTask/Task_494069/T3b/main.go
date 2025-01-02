package main

import (
	"fmt"
)

func main() {
	// Initialize a slice of fruits
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Initial fruits slice:", fruits) // Output: [apple banana cherry]

	// Assign fruits to moreFruits (shallow copy, shared underlying array)
	moreFruits := fruits
	fmt.Printf("fruits: %p, %v\n", &fruits, fruits)       // Memory address of fruits
	fmt.Printf("moreFruits: %p, %v\n", &moreFruits, moreFruits) // Same memory address as fruits

	// Append an element to moreFruits
	// Since moreFruits' capacity matches fruits, the append will create a new array
	moreFruits = append(moreFruits, "date")
	fmt.Println("moreFruits after append:", moreFruits) // [apple banana cherry date]
	fmt.Println("fruits remains unchanged:", fruits)    // [apple banana cherry]

	// Make a deep copy to avoid shared memory
	moreFruits2 := make([]string, len(fruits))
	copy(moreFruits2, fruits)

	// Append to the deep copy
	moreFruits2 = append(moreFruits2, "elderberry")
	fmt.Println("moreFruits2 after append:", moreFruits2) // [apple banana cherry elderberry]
	fmt.Println("fruits remains unchanged:", fruits)      // [apple banana cherry]
}
