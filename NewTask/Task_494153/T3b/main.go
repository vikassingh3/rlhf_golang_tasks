package main

import (
	"fmt"
)

// Function returning multiple value types
func getData() (interface{}, interface{}) {
	return "Hello", 42
}

// Helper function to simulate a test runner
func runTests() {
	tests := []struct {
		name      string
		expected1 interface{}
		expected2 interface{}
	}{
		{"test success", "Hello", 42},
	}

	for _, tt := range tests {
		fmt.Printf("Running %s...\n", tt.name)
		result1, result2 := getData()
		if result1 != tt.expected1 || result2 != tt.expected2 {
			fmt.Printf("FAIL: %s - expected (%v, %v), got (%v, %v)\n",
				tt.name, tt.expected1, tt.expected2, result1, result2)
		} else {
			fmt.Printf("PASS: %s\n", tt.name)
		}
	}
}

func main() {
	fmt.Println("Running manual test runner...")
	runTests()
}
