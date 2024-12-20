package main

import (
	"fmt"
	"testing"
)

// BenchmarkRangeWithIndex benchmarks the range loop with index
func BenchmarkRangeWithIndex(b *testing.B) {
	slice := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		for range slice {
		}
	}
}

func BenchmarkRangeWithoutIndex(b *testing.B) {
	slice := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		for range slice {
		}
	}
}

func BenchmarkFor(b *testing.B) {
	slice := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(slice); i++ {
		}
	}
}

func BenchmarkRangeMap(b *testing.B) {
	mapData := make(map[string]int, 1000000)
	for i := 0; i < b.N; i++ {
		for range mapData {
		}
	}
}

func BenchmarkForMap(b *testing.B) {
	mapData := make(map[string]int, 1000000)
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(mapData); i++ {
		}
	}
}

func main() {
	fmt.Println("Run the tests with 'go test -bench .':")
}
