package main

import "fmt"

func main() {
    // Case 1: Appending to a slice during iteration
    slice := []int{1, 2, 3, 4, 5}
    for _, val := range slice {
        slice = append(slice, val*2) // This can lead to undefined behavior.
    }
    fmt.Println("Final slice:", slice)

    // Case 2: Deleting map entries during iteration
    m := map[string]int{"key1": 1, "key2": 2, "key3": 3}
    for value,  _ := range m {
        delete(m, value)
    }

    // Case 3: Slicing a slice during iteration
    slice = []int{1, 2, 3, 4, 5}
    for _, val := range slice {
        if val == 3 {
            slice = slice[1:] // This will skip element 4.
        }
    }
    fmt.Println("Final slice:", slice)

    // Case 4: Removing an element from a slice during iteration
    slice = []int{1, 2, 3, 4, 5}
    for index, val := range slice {
        if val == 3 {
            slice = append(slice[:index], slice[index+1:]...) // This can shift indices.
        }
    }
    fmt.Println("Final slice:", slice)
}
