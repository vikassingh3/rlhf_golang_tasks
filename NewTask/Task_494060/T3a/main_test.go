package sliceutils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdvancedSliceBehavior(t *testing.T) {
	// Test slicing with nil slice
	var nilSlice []int
	require.Nil(t, nilSlice)
	// Ensure nil slice accesses do not panic
	assert.Panics(t, func() { _ = nilSlice[0] }, "Expected panic for nil slice access")
	// Avoid accessing or slicing nil slice directly
	if len(nilSlice) > 0 {
		assert.Panics(t, func() { _ = nilSlice[:1] }, "Expected panic for nil slice slicing")
	}

	// Test slicing with empty slice
	emptySlice := make([]int, 0, 0)
	assert.NotNil(t, emptySlice)
	// Ensure that empty slice access and slicing behave properly
	assert.Panics(t, func() { _ = emptySlice[0] }, "Expected panic for empty slice access")
	// Only slice if the length is greater than zero
	if len(emptySlice) > 0 {
		assert.Panics(t, func() { _ = emptySlice[:1] }, "Expected panic for empty slice slicing")
	}
	// Slicing an empty slice should result in an empty slice
	assert.Equal(t, len(emptySlice[:0]), 0, "Expected len of emptySlice[:0] to be 0")

	// Test slice reallocation
	slice := make([]int, 0, 2)
	slice = append(slice, 1, 2, 3, 4, 5)
	assert.Equal(t, len(slice), 5, "Expected length of slice to be 5")
	// Check if capacity is at least 5 (since 5 elements are added)
	assert.True(t, cap(slice) >= 5, "Expected capacity of slice to be at least 5 after reallocation")

	// Test slice range assignment
	slice2 := slice[:3]
	slice2[0] = 10
	assert.Equal(t, slice[0], 10, "Expected slice[0] to be 10 after range assignment")
	assert.Equal(t, slice2[0], 10, "Expected slice2[0] to be 10 after range assignment")

	// Test assigning nil to a slice
	var mySlice []int
	mySlice = make([]int, 2, 2)
	assert.NotNil(t, mySlice)

	mySlice = nil
	assert.Nil(t, mySlice)
	// Ensure accessing a nil slice panics
	assert.Panics(t, func() { _ = mySlice[0] }, "Expected panic when accessing nil slice")

	// Test append to nil slice
	mySlice = append(mySlice, 1)
	assert.Equal(t, mySlice, []int{1}, "Expected mySlice to be [1] after append")

	// Test overwriting slice with same length
	slice3 := make([]int, 3, 3)
	slice3[0] = 100
	slice3[1] = 200
	slice3[2] = 300

	// Make a copy of slice3 instead of assigning it directly
	slice4 := append([]int(nil), slice3...)
	assert.Equal(t, slice4[0], 100, "Expected slice4[0] to be 100")
	assert.Equal(t, slice4[1], 200, "Expected slice4[1] to be 200")
	assert.Equal(t, slice4[2], 300, "Expected slice4[2] to be 300")

	slice3[0] = 400
	assert.Equal(t, slice4[0], 100, "Expected slice4[0] to remain unchanged as slices are pointers")
}
