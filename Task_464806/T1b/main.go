package main

import (
    "fmt"
)

func main() {
    mySlice := []int{1, 2, 3, 4, 5}
    
    // Incorrect use of range loop over slice with direct len() function call
    for i, v := range mySlice {
        if i == len(mySlice) - 1 {
            fmt.Println("Last element:", v)
        }
    }

    // Correct use of range loop over slice with a variable for slice length
    sliceLength := len(mySlice)
    for i, v := range mySlice {
        if i == sliceLength - 1 {
            fmt.Println("Last element:", v)
        }
    }
    
    myMap := map[string]int{"a": 1, "b": 2, "c": 3}
    
    // Incorrect use of range loop over map to iterate over values
    for _, value := range myMap {
        delete(myMap, "b")
        fmt.Println("Value:", value)
    }

    // Correct use of range loop over map to iterate over keys
    for key := range myMap {
        fmt.Println("Key:", key)
    }
}  