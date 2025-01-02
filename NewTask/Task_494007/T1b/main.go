package main  
import (  
    "fmt"
)  
func main() {  
    slice := []int{1, 2, 3, 4, 5}

    // Using range loop to iterate over the slice
    for index, value := range slice {  
        fmt.Printf("Index: %d, Value: %d\n", index, value)  
    }  
}