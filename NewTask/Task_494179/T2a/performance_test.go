package main

import (
	"testing"
)

func multipleReturns(x, y float64) (quotient, remainder float64) {
	quotient = x / y
	remainder = x - quotient*y
	return
}

type Result struct {
	Quotient, Remainder float64
}

func singleReturn(x, y float64) Result {
	quotient := x / y
	remainder := x - quotient*y
	return Result{quotient, remainder}
}

// Test functions should not use 'b *testing.B'. Use 't *testing.T' and run the function a fixed number of times.
func TestMultipleReturns(t *testing.T) {
	x := 10.0
	y := 3.0

	// Run the function a fixed number of times, for example, 1000 iterations.
	for i := 0; i < 1000; i++ {
		multipleReturns(x, y)
	}
}

func TestSingleReturn(t *testing.T) {
	x := 10.0
	y := 3.0

	// Run the function a fixed number of times, for example, 1000 iterations.
	for i := 0; i < 1000; i++ {
		singleReturn(x, y)
	}
}

// Benchmark functions use 'b *testing.B' and use 'b.N' for iteration.
func BenchmarkMultipleReturns(b *testing.B) {
	x := 10.0
	y := 3.0

	// Benchmark the function with b.N iterations.
	for i := 0; i < b.N; i++ {
		multipleReturns(x, y)
	}
}

func BenchmarkSingleReturn(b *testing.B) {
	x := 10.0
	y := 3.0

	// Benchmark the function with b.N iterations.
	for i := 0; i < b.N; i++ {
		singleReturn(x, y)
	}
}
