package main  
import ("fmt")

func main() {
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Println("Initial fruits slice:", fruits) // Output: Initial fruits slice: [apple banana cherry]
    // Step 1: Observe the slice before making any changes.

    moreFruits := fruits
    fmt.Printf("fruits: %p, %v\n", &fruits, fruits) // Output: fruits: 0xc0000101a0, [apple banana cherry]
    fmt.Printf("moreFruits: %p, %v\n", &moreFruits, moreFruits) // Output: moreFruits: 0xc0000101a0, [apple banana cherry]
    // Step 2: Check the memory addresses of fruits and moreFruits. Since they are the same, it suggests they are shallow copies.

    moreFruits = append(moreFruits, "date")
    fmt.Println("moreFruits after append:", moreFruits) // Output: moreFruits after append: [apple banana cherry date]
    // Step 3: Append an element to moreFruits and observe the changes.

    fmt.Println("fruits after append to moreFruits:", fruits) // Output: fruits after append to moreFruits: [apple banana cherry]
    // Step 4: **Unexpected behavior!** fruits is unchanged after appending to moreFruits. This is the issue.

    // Step 5: Now, let's try to make a deep copy using copy() function.
    moreFruits2 := make([]string, len(fruits))
    copy(moreFruits2, fruits)

    moreFruits2 = append(moreFruits2, "elderberry")
    fmt.Println("moreFruits2 after append:", moreFruits2) // Output: moreFruits2 after append: [apple banana cherry elderberry]
    // Step 6: Observe the changes in moreFruits2 after making a deep copy.

    fmt.Println("fruits after append to moreFruits2:", fruits) // Output: fruits after append to moreFruits2: [apple banana cherry]
    // Step 7: fruits is still unchanged, which confirms that we fixed the issue by making a deep copy.
}