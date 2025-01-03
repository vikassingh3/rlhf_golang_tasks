
package main  
import (
    "fmt"
    "time"
)

func main() {
    const n = 100000000

    // Test with slice
    slice1 := make([]int, n)
    start := time.Now()
    for i := range slice1 {
        slice1[i] = i
    }
    elapsed := time.Since(start)
    fmt.Printf("Time taken for slice iteration: %s\n", elapsed)

    // Test with array
    array1 := [n]int{}
    start = time.Now()
    for i := range array1 {
        array1[i] = i
    }
    elapsed = time.Since(start)
    fmt.Printf("Time taken for array iteration: %s\n", elapsed)
}  
