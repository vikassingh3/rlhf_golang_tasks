package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOperations(t *testing.T) {
	testCases := []struct {
		name           string
		initialSlice   []int
		operations     func([]int) []int
		expectedSlice  []int
		expectedCap    int
		panicExpected bool
	}{
		{
			name:           "Append to empty slice",
			initialSlice:   []int{},
			operations:     func(slice []int) []int { return append(slice, 1) },
			expectedSlice:  []int{1},
			expectedCap:    1,
			panicExpected: false,
		},
		{
			name:           "Append to non-empty slice",
			initialSlice:   []int{1, 2, 3},
			operations:     func(slice []int) []int { return append(slice, 4) },
			expectedSlice:  []int{1, 2, 3, 4},
			expectedCap:    4,
			panicExpected: false,
		},
		{
			name:           "Reallocate when appending",
			initialSlice:   make([]int, 0, 2),
			operations:     func(slice []int) []int { return append(slice, 1, 2, 3) },
			expectedSlice:  []int{1, 2, 3},
			expectedCap:    4,
			panicExpected: false,
		},
		{
			name:           "Append to nil slice",
			initialSlice:   nil,
			operations:     func(slice []int) []int { return append(slice, 1) },
			expectedSlice:  []int{1},
			expectedCap:    1,
			panicExpected: false,
		},
		{
			name:           "Index out of bounds",
			initialSlice:   []int{1, 2, 3},
			operations:     func(slice []int) []int { return slice},
			expectedSlice:  []int{}, // Not used in this case, but needed for the struct
			expectedCap:    0, // Not used in this case, but needed for the struct
			panicExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicExpected {
				assert.Panics(t, func() { tc.operations(tc.initialSlice) })
			} else {
				resultSlice := tc.operations(tc.initialSlice)
				assert.Equal(t, tc.expectedSlice, resultSlice, "Slice values mismatch")
				assert.Equal(t, tc.expectedCap, cap(resultSlice), "Slice capacity mismatch")
			}
		})
	}
} 

 