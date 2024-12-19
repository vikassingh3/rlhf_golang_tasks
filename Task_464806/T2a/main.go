package main

import (
    "fmt"
)

func removeDuplicates(m map[int]int) {
    keys := make([]int, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }

    for _, key := range keys {
        value := m[key]
        if value > 1 {
            delete(m, key)
        }
    }
}

func main() {
    data := map[int]int{1: 1, 2: 2, 3: 3, 4: 3, 5: 4}

    removeDuplicates(data)

    fmt.Println(data)
}