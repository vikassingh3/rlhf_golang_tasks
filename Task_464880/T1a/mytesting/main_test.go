package main

import (
	"fmt"
	"testing"
)

func BenchmarkRange(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		for index, _ := range slice {
		}
	}
}

func BenchmarkFor(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(slice); i++ {
		}
	}
}

func main() {
	fmt.Println("Benchmarking range and for loops:")
	testing.Benchmark(BenchmarkRange)
	testing.Benchmark(BenchmarkFor)
}
