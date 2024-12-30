package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesEdgeCases(t *testing.T) {
	// Bounds checking
	testSlice := []int{1, 2, 3}
	assert.Panics(t, func() { _ = testSlice[3] }, "Expected panic for index out of bounds")

	// Slice reallocation
	testSlice2 := make([]int, 0, 2)
	testSlice2 = append(testSlice2, 1, 2, 3) // Reallocation occurs here
	assert.Equal(t, []int{1, 2, 3}, testSlice2, "Reallocation failed")

	// Handling nil slices
	var nilSlice []int
	assert.Nil(t, nilSlice, "Slice should be nil")
	nilSlice = append(nilSlice, 1) // Appending to nil slices is valid
	assert.Equal(t, []int{1}, nilSlice, "Append to nil slice failed")
}
