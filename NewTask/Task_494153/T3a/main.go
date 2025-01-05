package main

import (
	"errors"
	"testing"
)

// Function that can return multiple values, including an error
func getData() (string, int, error) {
	if randomError() {
		return "", 0, errors.New("random error occurred")
	}
	return "Hello", 42, nil
}

// Simulate a random error
func randomError() bool {
	// Replace with actual random logic
	return true // For demonstration purposes, let's always return an error
}

func TestGetData(t *testing.T) {
	// Test case 1: Expect an error
	_, _, err := getData()
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != "random error occurred" {
		t.Errorf("Expected error message 'random error occurred', but got '%v'", err.Error())
	}

	// Test case 2: Expect no error and specific data
	// Note: This test case will fail because randomError() always returns true in the implementation
	str, num, err := getData()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	} else if str != "Hello" || num != 42 {
		t.Errorf("Expected string 'Hello' and number 42, but got string '%v' and number %d", str, num)
	}
}

