// slices/slices.go
package slices

// AddElements appends a list of elements to a slice.
func AddElements(slice []int, elements ...int) []int {
	for _, e := range elements {
		slice = append(slice, e)
	}
	return slice
}