package main

import (
    "fmt"
    "sync"
)

func main() {
    // Fix 1: Appending to a slice while iterating
    slice := []int{1, 2, 3, 4, 5}
    newValues := []int{}

    for _, val := range slice {
        if val > 3 {
            newValues = append(newValues, val*2)
        }
        fmt.Println(val)
    }

    slice = append(slice, newValues...) // Append new values after iteration
    fmt.Println("Final slice:", slice)

    // Fix 2: Concurrent map writes
    var m sync.Map
    for i := 0; i < 10; i++ {
        go func(i int) {
            m.Store(fmt.Sprintf("key%d", i), i) // Safely store values
        }(i)
    }

    m.Range(func(k, v interface{}) bool {
        fmt.Println(k, v)
        return true
    })

    // Fix 3: Map key and value iteration
    myMap := map[string]int{"apple": 2, "banana": 5}
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }

    // Fix 4: Processing data
    data := []int{10, 20, 30}
    process := func(item int) {
        fmt.Println("Processing item:", item)
    }

    for _, item := range data {
        process(item)
    }
}
