package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesEdgeCases(t *testing.T) {
	// Test 1: Bounds checking (Index out of range)
	testSlice := []int{1, 2, 3}
	assert.Panics(t, func() { _ = testSlice[3] }, "Expected panic for index out of bounds")

	// Test 2: Slice reallocation
	// Initialize a slice with pre-allocated capacity
	testSlice2 := make([]int, 0, 2)

	// Append more elements than the pre-allocated capacity
	testSlice2 = append(testSlice2, 1, 2, 3, 4) // Reallocation occurs
	assert.Equal(t, []int{1, 2, 3, 4}, testSlice2, "Expected [1, 2, 3, 4], got %v", testSlice2)

	// Test 3: Appending to a nil slice
	var nilSlice []int
	assert.Nil(t, nilSlice, "Expected slice to be nil")

	// Append an element to the nil slice (this will initialize the slice)
	nilSlice = append(nilSlice, 1)
	assert.NotNil(t, nilSlice, "Expected slice to be initialized to non-nil")
	assert.Equal(t, []int{1}, nilSlice, "Expected [1], got %v", nilSlice)

	// Test 4: Appending more than 100 elements to a slice of capacity 10
	testSlice3 := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		testSlice3 = append(testSlice3, i)
	}
	expectedSlice := make([]int, 100)
	for i := 0; i < 100; i++ {
		expectedSlice[i] = i
	}
	assert.Equal(t, expectedSlice, testSlice3, "Expected %v, got %v", expectedSlice, testSlice3)
}

// Example for debugging and inspection
func Example() {
	slice := make([]int, 0, 2)
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
		// Print slice details for inspection
		println("Slice:", slice, "Capacity:", cap(slice))
	}
}

func main() {
	// Running tests using `go test` is the standard way.
	// Do not run tests manually in main.
}
