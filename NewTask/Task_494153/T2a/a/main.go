package main

import (
	"fmt"
	"errors"
)

// Example function that returns a list of integers and an error.
func listWithErrors(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("negative input")
	}
	l := make([]int, n)
	for i := range l {
		l[i] = i
	}
	return l, nil
}

func main() {
	// Assertion for error handling
	list, err := listWithErrors(-2)
	if err == nil {
		fmt.Println("Expected error, but got:", list)
	} else {
		fmt.Println("Got expected error:", err.Error())
	}
}