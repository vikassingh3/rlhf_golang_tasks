package main

import (
	"errors"
	"fmt"
	"testing"
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

func TestListWithErrors(t *testing.T) {
	// Test case 1: Positive input
	result, err := listWithErrors(3)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err.Error())
	}
	expectedResult := []int{0, 1, 2}
	if !eq(result, expectedResult) {
		t.Errorf("Expected %#v, got %#v", expectedResult, result)
	}

	// Test case 2: Negative input
	result, err = listWithErrors(-2)
	if err == nil {
		t.Fatalf("Expected error, got: %#v", result)
	}
	expectedError := errors.New("negative input")
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected %q, got %q", expectedError.Error(), err.Error())
	}
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMain(m *testing.M) {
	if err := m.Run(); err != 0 {
		fmt.Println(err)
	}
}