package main

import (
	"fmt"
	"log"
	"sync"
)

func processSliceValue(value int) {
	fmt.Println("Processing value:", value)
}

func main() {
	// Fixing map iteration issues
	fmt.Println("\nIterating over a map:")
	myMap := map[string]int{"apple": 2, "banana": 5}
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Concurrent slice modification issue fixed
	fmt.Println("\nAppending to a slice with proper synchronization:")
	var slice []int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			slice = append(slice, val)
		}(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final slice after concurrent appending:", slice)

	// Nested map iteration
	fmt.Println("\nIterating over a nested map:")
	data := map[string][]int{"a": {1, 2}, "b": {3, 4}}
	for key, values := range data {
		for _, value := range values {
			fmt.Printf("Key: %s, Value: %d\n", key, value)
		}
	}

	// Processing slice and doubling values
	fmt.Println("\nDoubling slice values:")
	var doubledSlice []int
	for _, value := range slice {
		doubledSlice = append(doubledSlice, value*2)
	}
	fmt.Println("Doubled slice:", doubledSlice)

	// Safe concurrent slice processing
	fmt.Println("\nConcurrently processing slice values:")
	for _, value := range slice {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			processSliceValue(v)
		}(value)
	}
	wg.Wait()

	// Logging slice processing
	fmt.Println("\nLogging slice processing:")
	for _, value := range slice {
		log.Println("Processing value:", value)
	}

	// Iterating over a copy of the slice
	fmt.Println("\nProcessing a copy of the slice:")
	var copiedSlice []int
	for _, value := range slice {
		processSliceValue(value)
		copiedSlice = append(copiedSlice, value)
	}
	fmt.Println("Copied slice:", copiedSlice)

	// String iteration
	fmt.Println("\nIterating over a string with runes:")
	str := "Hello, 世界"
	for index, runeValue := range str {
		fmt.Printf("Index: %d, Rune: %c, UTF-8 bytes: % x\n", index, runeValue, []byte(string(runeValue)))
	}
}
