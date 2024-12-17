package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// DataProcessor processes data and returns an error if something goes wrong
type DataProcessor struct {
	errorSimulated bool // for simulating an error in tests
}

func (dp *DataProcessor) Process(data string) (string, error) {
	if dp.errorSimulated {
		return "", errors.New("simulated processing error")
	}
	processedData := fmt.Sprintf("Processed: %s", data)
	return processedData, nil
}

// ToggleErrorSimulation switches error simulation on or off
func (dp *DataProcessor) ToggleErrorSimulation(flag bool) {
	dp.errorSimulated = flag
}

// Mock of the DataProcessor
type MockDataProcessor struct {
	mock.Mock
}

func (m *MockDataProcessor) Process(data string) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}

func main() {
	// Example usage of DataProcessor
	dp := &DataProcessor{}
	result, err := dp.Process("Hello World")
	if err != nil {
		fmt.Println("Error processing data:", err)
	} else {
		fmt.Println(result)
	}

	// Simulate error
	dp.ToggleErrorSimulation(true)
	_, err = dp.Process("Hello World")
	if err != nil {
		fmt.Println("Error during simulation:", err)
	}

	// Example usage of MockDataProcessor (Mock for testing)
	mockProcessor := &MockDataProcessor{}
	mockProcessor.On("Process", "test data").Return("", errors.New("mock error")) // simulate error

	_, err = mockProcessor.Process("test data")
	if err != nil {
		fmt.Println("Error from mock processor:", err)
	}

	mockProcessor.AssertExpectations(nil) // Passing nil because we are not calling testing.T here.
}

// Unit Tests
func TestProcessWithoutError(t *testing.T) {
	dp := &DataProcessor{}
	result, err := dp.Process("test data")

	if err != nil {
		t.Errorf("Expected no error, got: %s", err)
	}

	expected := "Processed: test data"
	if result != expected {
		t.Errorf("Expected %s, got: %s", expected, result)
	}
}

func TestProcessWithError(t *testing.T) {
	dp := &DataProcessor{}
	dp.ToggleErrorSimulation(true) // simulate error

	_, err := dp.Process("test data")
	if err == nil || err.Error() != "simulated processing error" {
		t.Errorf("Expected simulated processing error, got: %v", err)
	}
}

func TestIntegrationWithMock(t *testing.T) {
	mockProcessor := &MockDataProcessor{}
	mockProcessor.On("Process", "test data").Return("", errors.New("mock error")) // simulate error

	_, err := mockProcessor.Process("test data")
	if err == nil || err.Error() != "mock error" {
		t.Errorf("Expected mock error, got: %v", err)
	}

	mockProcessor.AssertExpectations(t)
}
