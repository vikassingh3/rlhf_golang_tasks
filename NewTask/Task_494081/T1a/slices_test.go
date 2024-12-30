// slices_test.go
package slices

import (
	"testing"
)

func TestAddElements(t *testing.T) {
	// Test case 1: Append to an empty slice
	var emptySlice []int
	result := AddElements(emptySlice, 1, 2, 3)
	expected := []int{1, 2, 3}
	if !sliceEqual(result, expected) {
		t.Errorf("TestAddElements: expected %v, got %v", expected, result)
	}

	// Test case 2: Append to a non-empty slice
	var nonEmptySlice =[]int{4, 5}
	result = AddElements(nonEmptySlice, 6, 7)
	expected = []int{4, 5, 6, 7}
	if !sliceEqual(result, expected) {
		t.Errorf("TestAddElements: expected %v, got %v", expected, result)
	}
}

func TestSliceEquality(t *testing.T) {
	// Test case 1: Equal slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	if !sliceEqual(slice1, slice2) {
		t.Errorf("TestSliceEquality: slices %v and %v are not equal", slice1, slice2)
	}

	// Test case 2: Unequal slices
	slice3 := []int{1, 2}
	slice4 := []int{1, 2, 3}
	if sliceEqual(slice3, slice4) {
		t.Errorf("TestSliceEquality: slices %v and %v are equal", slice3, slice4)
	}
}

func TestSliceSlicing(t *testing.T) {
	// Test case 1: Slice a slice
	originalSlice := []int{1, 2, 3, 4, 5}
	slicedSlice := originalSlice[1:4]
	expectedSlicedSlice := []int{2, 3, 4}
	if !sliceEqual(slicedSlice, expectedSlicedSlice) {
		t.Errorf("TestSliceSlicing: expected %v, got %v", expectedSlicedSlice, slicedSlice)
	}

	// Test case 2: Slice with a single element
	originalSlice = []int{6}
	slicedSlice = originalSlice[:]
	expectedSlicedSlice = []int{6}
	if !sliceEqual(slicedSlice, expectedSlicedSlice) {
		t.Errorf("TestSliceSlicing: expected %v, got %v", expectedSlicedSlice, slicedSlice)
	}
}

func sliceEqual(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}