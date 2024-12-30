package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlicesEdgeCases(t *testing.T) {
	// Test 1: Bounds checking
	testSlice := []int{1, 2, 3}
	require.Equal(t, 1, testSlice[0], "Failed to access the first element")
	require.Equal(t, 3, testSlice[2], "Failed to access the last element")

	// Testing bounds checking by trying to access an out-of-bounds index
	defer func() {
		if r := recover(); r != nil {
			// Use fmt.Sprintf to safely handle the panic message as a string
			errorMessage := fmt.Sprintf("%v", r)
			assert.Contains(t, errorMessage, "index out of range", "Expected index out of range panic")
		}
	}()
	_ = testSlice[3] // Out-of-bounds access

	// Test 2: Slice reallocation
	testSlice2 := make([]int, 0, 2) // Initial capacity of 2
	testSlice2 = append(testSlice2, 1)
	assert.Equal(t, []int{1}, testSlice2, "Failed to append element")
	testSlice2 = append(testSlice2, 2)
	assert.Equal(t, []int{1, 2}, testSlice2, "Failed to append another element")
	testSlice2 = append(testSlice2, 3) // Should trigger reallocation
	assert.Equal(t, []int{1, 2, 3}, testSlice2, "Failed to reallocate slice")

	// Test 3: Handling nil slices
	var nilSlice []int
	assert.Nil(t, nilSlice, "Expected nil slice")

	// Appending to a nil slice
	defer func() {
		if r := recover(); r != nil {
			assert.Fail(t, "Did not expect a panic when appending to a nil slice")
		}
	}()
	nilSlice = append(nilSlice, 1)
	assert.Equal(t, []int{1}, nilSlice, "Expected successful append to nil slice")

	// Assigning a slice to nil
	var sliceToNil []int = []int{1, 2}
	sliceToNil = nil
	assert.Nil(t, sliceToNil, "Expected nil slice after assignment")
}
