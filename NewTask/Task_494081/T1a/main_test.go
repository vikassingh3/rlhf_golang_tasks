package sliceops

// AddToSlice appends a new element to the slice.
func AddToSlice(slice []int, newElement int) []int {
	return append(slice, newElement)
}

// RemoveFromSlice removes the first occurrence of an element from the slice.
func RemoveFromSlice(slice []int, elementToRemove int) []int {
	for i, val := range slice {
		if val == elementToRemove {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// IsElementInSlice checks if an element exists in the slice.
func IsElementInSlice(slice []int, element int) bool {
	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}

// IterateSlice applies a function to each element of the slice.
func IterateSlice(slice []int, funcToApply func(int)) {
	for _, val := range slice {
		funcToApply(val)
	}
}