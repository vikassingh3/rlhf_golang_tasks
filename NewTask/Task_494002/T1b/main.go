package main

import (
	"fmt"
)
func main() {
   //example 1: Iterating over a slice
   slice1 := []int{1, 2, 3, 4, 5}
   for index, value := range slice1 {
      fmt.Printf("Index: %d, Value: %d\n", index, value)
   }
   //output : Index: 0, Value: 1 Index: 1, Value: 2 Index: 2, Value: 3 Index: 3, Value: 4 Index: 4, Value: 5

   //example 2: Iterating over an array
   array1 := [5]int{6, 7, 8, 9, 10}
   for index, value := range array1 {
      fmt.Printf("Index: %d, Value: %d\n", index, value)
   }
   //output : Index: 0, Value: 6 Index: 1, Value: 7 Index: 2, Value: 8 Index: 3, Value: 9 Index: 4, Value: 10 
}
