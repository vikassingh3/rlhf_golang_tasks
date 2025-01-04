package main

import (
	"fmt"
)

func divide(x, y float64) (quotient, remainder float64) {
	quotient = x / y
	remainder = x - quotient*y
	return
}

func main() {
	x := 10.0
	y := 3.0
	
	quotient, remainder := divide(x, y)
	
	fmt.Printf("Quotient: %.2f, Remainder: %.2f\n", quotient, remainder)
}