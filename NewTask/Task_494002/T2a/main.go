package main

import (
	"fmt"
	"time"
)

func arrayRangeLoop() {
	array := make([]int, 100000)
	for i := range array {
		array[i] = i
	}
}

func sliceRangeLoop() {
	slice := make([]int, 100000)
	for i := range slice {
		slice[i] = i
	}
}

func main() {
	arrayStart := time.Now()
	arrayRangeLoop()
	arrayTime := time.Since(arrayStart)

	sliceStart := time.Now()
	sliceRangeLoop()
	sliceTime := time.Since(sliceStart)

	fmt.Printf("Array range loop time: %v\n", arrayTime)
	fmt.Printf("Slice range loop time: %v\n", sliceTime)
}