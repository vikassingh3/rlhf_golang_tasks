package main  
import (  
    "fmt"
)
func removeElement(s []int, val int) []int {
    var result []int
    for _, v := range s {
        if v != val {
            result = append(result, v)
        }
    }
    return result
}
func main() {
    nums := []int{3, 2, 2, 3}
    valToRemove := 3
    result := removeElement(nums, valToRemove)
    fmt.Println("Original slice:", nums)
    fmt.Println("Slice after removing", valToRemove, ":", result)

    // Handling an empty slice
    emptySlice := []int{}
    result = removeElement(emptySlice, 10)
    fmt.Println("Result from removing from an empty slice:", result)

    // Handling a slice with only one element
    singleElementSlice := []int{42}
    result = removeElement(singleElementSlice, 42)
    fmt.Println("Result from removing from a slice with one element:", result)
}