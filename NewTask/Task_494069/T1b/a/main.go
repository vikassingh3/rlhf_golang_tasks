package main
import ("fmt")

func main() {
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Println("Initial fruits slice:", fruits) // Output: Initial fruits slice: [apple banana cherry]

    moreFruits := fruits
    moreFruits = append(moreFruits, "date")
    fmt.Println("moreFruits after append:", moreFruits) // Output: moreFruits after append: [apple banana cherry date]

    fmt.Println("fruits after append to moreFruits:", fruits) // Output: fruits after append to moreFruits: [apple banana cherry]
}