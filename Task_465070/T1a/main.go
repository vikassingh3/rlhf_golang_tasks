package main

import "fmt"

func addAndSubtract(a int, b int) (int, int) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func main() {
	resultSum, resultDifference := addAndSubtract(5, 3)
	fmt.Printf("Sum: %d, Difference: %d\n", resultSum, resultDifference)
}