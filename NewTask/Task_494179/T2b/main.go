package main

import (
	"math/rand"
	"testing"
)

func generateSlice(size int) []int {  
    slice := make([]int, size)  
    for i := 0; i < size; i++ {  
        slice[i] = rand.Intn(10000)  
    }  
    return slice  
}

func singleValueReturn(size int) []int {  
    return generateSlice(size)  
}

func multipleValueReturn(size int) ([]int, []int, []int) {  
    return generateSlice(size), generateSlice(size), generateSlice(size)  
}

func BenchmarkSingleValueReturn(b *testing.B) {  
    for n := 0; n < b.N; n++ {  
        singleValueReturn(100000)  
    }  
}

func BenchmarkMultipleValueReturn(b *testing.B) {  
    for n := 0; n < b.N; n++ {  
        _, _, _ = multipleValueReturn(100000)  
    }  
}

func main() {  
    // Run the benchmarks  
    testing.Benchmark(BenchmarkSingleValueReturn)  
    testing.Benchmark(BenchmarkMultipleValueReturn)  
}