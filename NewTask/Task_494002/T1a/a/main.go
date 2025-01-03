package main

import "fmt"

func main() {
    var numbers = [3]int{1, 2, 3}
    
    // Iterate with both index and value
    for i, num := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", i, num)
    }
    
    // Iterate with only value
    for _, num := range numbers {
        fmt.Printf("Value: %d\n", num)
    }
}