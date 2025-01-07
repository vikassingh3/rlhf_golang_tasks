package sliceutils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestSliceAccess(t *testing.T) {
	// Test edge case: Accessing an element outside the slice bounds
	slice := []int{1, 2, 3}

	want := fmt.Errorf("index out of range")
	_, err := slice[10]  // Accessing index 10 out of bounds
	assert.EqualError(t, err, want.Error(), "expected index out of range error")

    // Test scenario: Slice reallocation on append
	slice = []int{}
	slice = append(slice, 1)
	wantLen := 1
	gotLen := len(slice)
	assert.Equal(t, wantLen, gotLen, "expected slice to reallocate and resize appropriately")
}

func TestSliceHandlingNil(t *testing.T) {
	// Test scenario: Handling a nil slice
	var slice []int
	assert.Equal(t, slice, []int(nil), "expected a nil slice")
	wantLen := 0
	gotLen := len(slice)
	assert.Equal(t, wantLen, gotLen, "expected len of nil slice to be 0")
	
	// Test : append to nil slice
	slice = append(slice, 1)
	want = []int{1}
	assert.Equal(t, slice, want, "expected slice to be appended when it's nil")
}


func TestSliceLength(t *testing.T) {
	cases := []struct {
		name    string
		slice   []int
		wantLen int
	}{
		// ... (Same test cases as before)
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Using testify's require.Len() assertion
			require.Len(t, tc.slice, tc.wantLen, "slice length should be as expected")
		})
	}
}
