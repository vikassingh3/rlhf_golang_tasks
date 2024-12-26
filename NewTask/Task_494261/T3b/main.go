package main

import (
	"fmt"
)

func main() {
	// Define two matrices
	matrix1 := []float32{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	matrix2 := []float32{
		10, 11, 12,
		13, 14, 15,
		16, 17, 18,
	}

	// Perform scalar multiplication
	result := make([]float32, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i*3+j] += matrix1[i*3+k] * matrix2[k*3+j]
			}
		}
	}

	// Print the result
	fmt.Println("Result:")
	fmt.Println(result)
}
