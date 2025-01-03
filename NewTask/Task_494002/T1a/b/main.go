package main

import "fmt"

func main() {
    var letters []string = []string{"a", "b", "c", "d"}
    
    // Iterate with both index and value
    for i, letter := range letters {
        fmt.Printf("Index: %d, Value: %d\n", i, letter)
    }
    
    // Iterate with only value
    for _, letter := range letters {
        fmt.Printf("Value: %s\n", letter)
    }
}