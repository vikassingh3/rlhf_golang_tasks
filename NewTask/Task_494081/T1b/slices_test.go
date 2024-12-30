package main

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
    // Test 1: Append an element to a slice
    testSlice1 := []int{1, 2, 3}
    expectedSlice1 := []int{1, 2, 3, 4}
    testSlice1 = append(testSlice1, 4)
    if fmt.Sprintf("%v", testSlice1) != fmt.Sprintf("%v", expectedSlice1) {
        t.Errorf("Test failed! Expected: %v, Actual: %v", expectedSlice1, testSlice1)
    }

    // Test 2: Insert an element at a specific index
    testSlice2 := []int{1, 2, 3}
    expectedSlice2 := []int{1, 4, 2, 3}
    testSlice2 = append(testSlice2[:1], append([]int{4}, testSlice2[1:]...)...)
    if fmt.Sprintf("%v", testSlice2) != fmt.Sprintf("%v", expectedSlice2) {
        t.Errorf("Test failed! Expected: %v, Actual: %v", expectedSlice2, testSlice2)
    }
  
    // Test 3: Delete an element by index
    testSlice3 := []int{1, 2, 3, 4}
    expectedSlice3 := []int{1, 2, 4}
    testSlice3 = append(testSlice3[:2], testSlice3[3:]...)
    if fmt.Sprintf("%v", testSlice3) != fmt.Sprintf("%v", expectedSlice3) {
        t.Errorf("Test failed! Expected: %v, Actual: %v", expectedSlice3, testSlice3)
    }
    
    // Test 4: Access element by index
    testSlice4 := []int{1, 2, 3}
    expectedValue4 := 2
    actualValue4 := testSlice4[1]
    if actualValue4 != expectedValue4 {
        t.Errorf("Test failed! Expected: %v, Actual: %v", expectedValue4, actualValue4)
    }
}
