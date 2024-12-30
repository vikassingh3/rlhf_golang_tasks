package sliceutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceAppendAndReallocation(t *testing.T) {
	t.Run("Append to nil slice", func(t *testing.T) {
		var slice []int
		require.Nil(t, slice)

		slice = append(slice, 1)
		assert.NotNil(t, slice)
		assert.Len(t, slice, 1)
		assert.Equal(t, slice[0], 1)
	})

	t.Run("Append to empty slice", func(t *testing.T) {
		slice := make([]int, 0, 0)
		assert.NotNil(t, slice)

		slice = append(slice, 2)
		assert.Len(t, slice, 1)
		assert.Equal(t, slice[0], 2)
	})

	t.Run("Append multiple elements", func(t *testing.T) {
		slice := make([]int, 0, 1)
		slice = append(slice, 3, 4, 5)
		assert.Len(t, slice, 3)
		assert.Equal(t, slice, []int{3, 4, 5})
	})

	t.Run("Reallocation on append", func(t *testing.T) {
		slice := make([]int, 0, 2)
		slice = append(slice, 1, 2, 3, 4, 5)
		assert.Len(t, slice, 5)
		assert.Equal(t, slice, []int{1, 2, 3, 4, 5})
		// Check if the capacity has increased after reallocation
		assert.Greater(t, cap(slice), 2)
	})

	t.Run("Append to nil slice multiple times", func(t *testing.T) {
		var slice []int
		for i := 1; i <= 5; i++ {
			slice = append(slice, i)
		}
		assert.NotNil(t, slice)
		assert.Len(t, slice, 5)
		assert.Equal(t, slice, []int{1, 2, 3, 4, 5})
	})
} 