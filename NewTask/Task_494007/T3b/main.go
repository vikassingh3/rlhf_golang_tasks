package main

import (
	"fmt"
)

//go:inline
func sum(a, b int) int {
	return a + b
}

func main() {
	result := sum(10, 20)
	fmt.Println("Result:", result)
}