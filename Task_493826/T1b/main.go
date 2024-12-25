package main  
import "fmt"

func main() {
    numbers := []int{1, 2, 3}
    safeAccess(numbers) // Passing the slice to the function
}

// This function attempts to access the last element of the slice
func safeAccess(s []int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    fmt.Println("Last element:", s[len(s)]) // Out-of-bounds access will cause panic
}
