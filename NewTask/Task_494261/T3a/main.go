package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Set the number of goroutines to use for parallel processing
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Dimensions of matrices
	m, n, p := 1000, 1000, 1000

	// Initialize matrices with random values
	matrixA := make([][]float64, m)
	matrixB := make([][]float64, n)
	matrixC := make([][]float64, m)

	// Initialize matrixA with random values
	for i := range matrixA {
		matrixA[i] = make([]float64, n)
		for j := range matrixA[i] {
			matrixA[i][j] = rand.Float64()
		}
	}

	// Initialize matrixB with random values
	for i := range matrixB {
		matrixB[i] = make([]float64, p)
		for j := range matrixB[i] {
			matrixB[i][j] = rand.Float64()
		}
	}

	// Initialize matrixC with zero values
	for i := range matrixC {
		matrixC[i] = make([]float64, p)
	}

	// Perform matrix multiplication using parallelism
	var wg sync.WaitGroup

	for i := 0; i < m; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < p; j++ {
				for k := 0; k < n; k++ {
					matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Print the result (printing only a small part of matrixC for readability)
	fmt.Printf("Matrix C [0][0]: %v\n", matrixC[0][0])
}
