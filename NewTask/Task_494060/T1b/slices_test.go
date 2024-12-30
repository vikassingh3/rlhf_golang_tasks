package sliceutils

import (
	"testing"
)

func TestSliceLength(t *testing.T) {
	cases := []struct {
		name    string
		slice   []int
		wantLen int
	}{
		{
			name:    "Empty slice",
			slice:   []int{},
			wantLen: 0,
		},
		{
			name:    "Single element slice",
			slice:   []int{42},
			wantLen: 1,
		},
		{
			name:    "Multiple element slice",
			slice:   []int{1, 2, 3, 4, 5},
			wantLen: 5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotLen := len(tc.slice)
			if gotLen != tc.wantLen {
				t.Errorf("len(%v) = %d, want %d", tc.slice, gotLen, tc.wantLen)
			}
		})
	}
} 
