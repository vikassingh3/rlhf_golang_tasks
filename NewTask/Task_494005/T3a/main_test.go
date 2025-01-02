package main

import (
	"math/rand"
	"testing"
)

// PreallocateSlice function pre-allocates a slice for the specified size and populates it.
func PreallocateSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Int()
	}
	return slice
}

// AppendInLoop function populates a slice by appending in a loop.
func AppendInLoop(size int) []int {
	slice := []int{}
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Int())
	}
	return slice
}

func BenchmarkPreallocateSlice(b *testing.B) {
	size := 1000000
	for i := 0; i < b.N; i++ {
		_ = PreallocateSlice(size)
	}
}

func BenchmarkAppendInLoop(b *testing.B) {
	size := 1000000
	for i := 0; i < b.N; i++ {
		_ = AppendInLoop(size)
	}
}

// The `testing.M` function is necessary to handle test setup and execution.
func main() {
	// Running tests with testing.Benchmark requires calling the test suite through `testing.M`
	testing.Main(matchString, nil, nil, nil)
}

// Helper function to match test patterns, as required by testing.Main.
func matchString(pat, str string) (bool, error) {
	return true, nil
}
