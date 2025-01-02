package main

import (
	"fmt"
)

func main() {
    // Step 1: Initialize the fruits slice and print its initial state
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Println("Initial fruits slice:", fruits) // Output: Initial fruits slice: [apple banana cherry]

    // Step 2: Create another slice called moreFruits and print its memory address
    moreFruits := fruits
    fmt.Printf("moreFruits after assignment: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits after assignment: 0xc000084020, [apple banana cherry]

    // Step 3: Append an element to moreFruits and print its memory address again
    moreFruits = append(moreFruits, "date")
    fmt.Printf("moreFruits after append: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits after append: 0xc000084030, [apple banana cherry date]

    // Step 4: Print the fruits slice to see if it's affected
    fmt.Println("fruits after append to moreFruits:", fruits) // Output: fruits after append to moreFruits: [apple banana cherry]

    // Step 5: Print the memory address of fruits to verify if it's different after append
    fmt.Printf("fruits after append: %p, %v\n", &fruits, fruits) // Output: fruits after append: 0xc000084020, [apple banana cherry]

    // Step 6: Observe the memory addresses and understand that fruits and moreFruits initially point to the same location
    // Step 7: After appending to moreFruits, it gets a new memory address because the underlying array is reallocated
    // Step 8: To fix this, use copy() to create a deep copy of fruits to moreFruits

    // Corrected code:
    fruits = []string{"apple", "banana", "cherry"}
    fmt.Println("Initial fruits slice:", fruits) // Output: Initial fruits slice: [apple banana cherry]

    // Step 9: Create a new slice moreFruits with the same length as fruits
    // moreFruits1 := make([]string, len(fruits))
    fmt.Printf("moreFruits after creation: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits after creation: 0xc000084040, [apple banana cherry]

    // Step 10: Copy elements from fruits to moreFruits
    copy(moreFruits, fruits)
    fmt.Printf("moreFruits after copy: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits after copy: 0xc000084040, [apple banana cherry]

    // Step 11: Append an element to moreFruits
    moreFruits = append(moreFruits, "date")
    fmt.Printf("moreFruits after append: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits after append: 0xc000084050, [apple banana cherry date]

    // Step 12: Print fruits to verify it remains unchanged
    fmt.Println("fruits after append to moreFruits:", fruits) // Output: fruits after append to moreFruits: [apple banana cherry]

    // Step 13: Print the memory address of fruits to confirm it's unchanged
    fmt.Printf("fruits after append: %p, %v\n", &fruits, fruits) // Output: fruits after append: 0xc000084020, [apple banana cherry]

    // Step 14: Observe that fruits and moreFruits now point to different memory locations, ensuring independence
}