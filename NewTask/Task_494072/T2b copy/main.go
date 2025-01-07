package main

import "fmt"

func main() {
    // Sample slice
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    // Traditional for loop (iterating backward)
    fmt.Println("Traditional for loop (reverse):")
    for i := len(slice) - 1; i >= 0; i-- {
        fmt.Println(slice[i]) // Perform operations on slice[i]
    }

    // Traditional for loop with skip (steps of 2)
    fmt.Println("\nTraditional for loop with skip (steps of 2):")
    for i := 0; i < len(slice); i += 2 {
        fmt.Println(slice[i]) // Perform operations on slice[i]
    }

    // Traditional for loop with element modification (doubling each element)
    fmt.Println("\nTraditional for loop with element modification:")
    for i := range slice {
        slice[i] *= 2  // Modifying the element directly
    }
    fmt.Println("Modified slice:", slice) // Display modified slice
}
