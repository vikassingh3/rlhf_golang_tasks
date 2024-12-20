package main

import (
	"testing"
)

func BenchmarkInefficientAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var result []int
		for i := 0; i < 1000; i++ {
			result = append(result, i) // Inefficient append() to the existing slice.
		}
	}
}

func main() {
	testing.Benchmark(BenchmarkInefficientAppend)
	testing.Benchmark(BenchmarkEfficientPreallocation)

}

func BenchmarkEfficientPreallocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := make([]int, 1000) // Efficient preallocation
		for i := 0; i < 1000; i++ {
			result[i] = i
		}
	}
}
