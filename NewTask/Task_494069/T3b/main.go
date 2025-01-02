package main

import (
    "fmt"
)

func main() {
    // Initializing a slice of strings called fruits
    fruits := []string{"apple", "banana", "cherry"}

    // Printing the initial fruits slice
    fmt.Println("Initial fruits slice:", fruits) // Output: Initial fruits slice: [apple banana cherry]

    // Assigning the fruits slice to a new variable called moreFruits. This creates a shallow copy,
    // meaning both variables point to the same underlying array.
    moreFruits := fruits

    // Printing the memory addresses of fruits and moreFruits to understand the shallow copy
    fmt.Printf("fruits: %p, %v\n", &fruits, fruits)    // Output: fruits: 0xc00008a0a0, [apple banana cherry]
    fmt.Printf("moreFruits: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits: 0xc00008a0a0, [apple banana cherry]

    // Now, we append the element "date" to the moreFruits slice. This works as expected, but...
    moreFruits = append(moreFruits, "date")
    fmt.Println("moreFruits after append:", moreFruits) // Output: moreFruits after append: [apple banana cherry date]

    // ...the change is unexpectedly also reflected in the fruits slice. This is because they share the same memory!
    fmt.Println("fruits after append to moreFruits:", fruits) // Output: fruits after append to moreFruits: [apple banana cherry date]

    // **The unexpected behavior arises because we intended to modify moreFruits separately from fruits.**
    // To fix this, we need to create a **deep copy** of the slice before modifying it.

    // We can use the 'make' function to create a new slice of the same length as fruits,
    // and then use the 'copy' function to copy the elements from fruits to this new slice.
    moreFruits2 := make([]string, len(fruits))
    copy(moreFruits2, fruits)

    // Printing the memory addresses of fruits, moreFruits, and moreFruits2 to illustrate the change
    fmt.Printf("fruits: %p, %v\n", &fruits, fruits)    // Output: fruits: 0xc00008a0a0, [apple banana cherry date]
    fmt.Printf("moreFruits: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits: 0xc00008a0a0, [apple banana cherry date]
    fmt.Printf("moreFruits2: %p, %v\n", &moreFruits2, moreFruits2) // Output: moreFruits2: 0xc00008c000, [apple banana cherry]

    // Now, we can safely append elements to moreFruits2 without affecting fruits
    moreFruits2 = append(moreFruits2, "elderberry")
    fmt.Println("moreFruits2 after append:", moreFruits2) // Output: moreFruits2 after append: [apple banana cherry elderberry]

    // Fruits remains unchanged, proving that we fixed the issue
    fmt.Println("fruits after append to moreFruits2:", fruits) // Output: fruits after append to moreFruits2: [apple banana cherry date]
}
 