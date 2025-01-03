package main

import (
	"fmt"
)

func getEvenNumbers(n int) []int {
  //Initialize a slice
  evens := []int{}

  //Add elements to the slice
  for i:=2; i<=n; i+=2 {
     evens = append(evens, i)
  }

  //Return the slice
  return evens
}

func main() {
    n := 10
    evenSlice := getEvenNumbers(n)

    //Print the returned slice
    fmt.Println("Even Numbers from 1 to", n, ":", evenSlice)

    //Access elements of the slice
    fmt.Println("First even number:", evenSlice[0])
}