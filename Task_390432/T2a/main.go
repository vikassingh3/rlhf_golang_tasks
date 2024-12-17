package main

import (
	"fmt"
	"testing"
)

// Custom error type for processing errors
type ProcessingError struct {
	Message string
}

func (e *ProcessingError) Error() string {
	return e.Message
}

// DataProcessor processes data
type DataProcessor struct{}

func (dp *DataProcessor) Process(data string) (string, error) {
	if data == "" {
		return "", &ProcessingError{Message: "data cannot be empty"} // return custom error
	}

	// Simulate processing logic
	processedData := fmt.Sprintf("Processed: %s", data)
	return processedData, nil
}

func TestProcessDataEmptyError(t *testing.T) {
	dp := &DataProcessor{}
	result, err := dp.Process("") // passing empty data to simulate error

	if err == nil {
		t.Fatal("Expected an error but got nil")
	}

	processingErr, ok := err.(*ProcessingError) // type assertion to check custom error
	if !ok {
		t.Fatalf("Expected error type *ProcessingError, got: %T", err)
	}

	expectedErrorMessage := "data cannot be empty"
	if processingErr.Message != expectedErrorMessage {
		t.Errorf("Expected error message '%s', got: '%s'", expectedErrorMessage, processingErr.Message)
	}

	if result != "" {
		t.Errorf("Expected result to be empty, got: '%s'", result)
	}
}

func TestProcessDataSuccess(t *testing.T) {
	dp := &DataProcessor{}
	result, err := dp.Process("sample data") // valid data processing

	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}

	expected := "Processed: sample data"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// main function to demonstrate the functionality
func main() {
	dp := &DataProcessor{}

	// Test with valid data
	result, err := dp.Process("hello world")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success:", result)
	}

	// Test with empty data (this will trigger an error)
	result, err = dp.Process("")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success:", result)
	}
}
