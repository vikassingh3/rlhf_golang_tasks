package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	size := 1000000
	data := make([]float64, size)
	for i := range data {
		data[i] = rand.Float64()
	}
	// Measure the baseline performance using range loop
	baseline := time.Now()
	var sum float64
	for _, value := range data {
		sum += value
	}
	elapsedBaseline := time.Since(baseline)
	fmt.Println("Baseline Time:", elapsedBaseline)
	fmt.Println("Baseline Sum:", sum)

	// Optimized 1: Avoid unnecessary allocation of variables within the loop.
	optimized1 := time.Now()
	var sumOptimized1 float64
	for _, value := range data {
		sumOptimized1 += math.Sin(value)
	}
	elapsedOptimized1 := time.Since(optimized1)
	fmt.Println("Optimized 1 Time:", elapsedOptimized1)
	fmt.Println("Optimized 1 Sum:", sumOptimized1)

   //  (The other optimization methods will be explained below)

	// **Optimization 3:  Use inline arithmetic and preallocate output slice**
	optimized3 := time.Now()
	output := make([]float64, size)
	for i := range data {
		v := data[i]
		output[i] = math.Sqrt(v) * math.Sqrt(1-v)
	}
	elapsedOptimized3 := time.Since(optimized3)
	fmt.Println("Optimized 3 Time:", elapsedOptimized3)
	// Note: Here we are just checking the size of the output slice, you can print the sum or use it as per your need.
	fmt.Println("Optimized 3 Output Length:", len(output))

} 