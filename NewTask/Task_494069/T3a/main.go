package main

import (
	"fmt"
)

// This function demonstrates the unexpected behavior of slices
// when they are shallow copied and then modified.
func main() {
	// Step 1: Initialize a slice of fruits.
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Initial fruits slice:", fruits) // Output: [apple banana cherry]

	// Step 2: Create a shallow copy of the fruits slice.
	// When a slice is copied in Go, the copy refers to the same underlying array
	// as the original slice. This means changes to either slice will affect the other.
	moreFruits := fruits

	// Step 3: Display the memory addresses of both slices to confirm they are the same.
	// This demonstrates that they are shallow copies.
	fmt.Printf("fruits: %p, %v\n", &fruits, fruits) // Output: fruits: <address>, [apple banana cherry]
	fmt.Printf("moreFruits: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits: <same address>, [apple banana cherry]

	// Step 4: Modify the moreFruits slice by appending a new element.
	moreFruits = append(moreFruits, "date")
	fmt.Println("moreFruits after append:", moreFruits) // Output: [apple banana cherry date]

	// Step 5: Check the fruits slice after modifying moreFruits.
	// Here, we observe the unexpected behavior: fruits is also modified.
	fmt.Println("fruits after append to moreFruits:", fruits) // Output: [apple banana cherry date]

	// Step 6: To resolve this, we need to make a deep copy of the slice.
	// Use make() to create a new slice with the same capacity, and then use copy()
	// to fill the new slice with the elements from the original slice.
	moreFruits2 := make([]string, len(fruits))
	copy(moreFruits2, fruits)

	// Step 7: Modify the moreFruits2 slice by appending a new element.
	// Since this is a deep copy, changes to moreFruits2 will not affect fruits.
	moreFruits2 = append(moreFruits2, "elderberry")
	fmt.Println("moreFruits2 after append:", moreFruits2) // Output: [apple banana cherry elderberry]

	// Step 8: Verify that fruits remains unchanged.
	fmt.Println("fruits after append to moreFruits2:", fruits) // Output: [apple banana cherry]

	// Conclusion: By using a deep copy, we ensure that modifications to one slice
	// do not affect the other, resolving the unexpected behavior.
}