package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// Custom error type for pipeline errors
type PipelineError struct {
	Stage string
	Err   error
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("Pipeline error at stage '%s': %v", p.Stage, p.Err)
}

// Error handlers
type ErrorHandler func(error) error

// NewErrorHandler returns an error handler that wraps the original error with the provided message.
func NewErrorHandler(message string) ErrorHandler {
	return func(err error) error {
		return errors.New(message + ": " + err.Error())
	}
}

// Stage1: Initial data cleaning.
func Stage1(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("data is empty after trimming")
	}
	return data, nil
}

// Stage2: Data normalization and format checking.
func Stage2(data string) (string, error) {
	if data == "" {
		return data, errors.New("data is empty after normalization")
	}
	// Simulate different types of errors based on data.
	switch {
	case strings.HasPrefix(data, "invalid"):
		return data, errors.New("invalid data format")
	case strings.HasSuffix(data, "*"):
		return data, errors.New("data is incomplete")
	}
	return data, nil
}

// Stage3: Final data validation.
func Stage3(data string) (string, error) {
	if strings.Contains(data, "@") {
		return data, errors.New("unexpected '@' in data")
	}

	if len(data) >= 100 {
		return data, errors.New("data length exceeds 100 characters")
	}

	if strings.ContainsAny(data, "0123456789") {
		return data, errors.New("unexpected numbers in data")
	}

	return data, nil
}

// ProcessData processes data through a pipeline with error handling.
func ProcessData(data string, stageHandlers map[string]ErrorHandler) (string, error) {
	var err error

	// Stage1
	data, err = Stage1(data)
	if err != nil {
		if handler, ok := stageHandlers["Stage1"]; ok {
			err = handler(err)
		}
		return "", &PipelineError{"Stage1", err}
	}

	// Stage2
	data, err = Stage2(data)
	if err != nil {
		if handler, ok := stageHandlers["Stage2"]; ok {
			err = handler(err)
		}
		return "", &PipelineError{"Stage2", err}
	}

	// Stage3
	data, err = Stage3(data)
	if err != nil {
		if handler, ok := stageHandlers["Stage3"]; ok {
			err = handler(err)
		}
		return "", &PipelineError{"Stage3", err}
	}

	return data, nil
}

// Example error handlers for each stage.
var stageHandlers = map[string]ErrorHandler{
	"Stage1": NewErrorHandler("Stage1 Error"),
	"Stage2": NewErrorHandler("Stage2 Error"),
	"Stage3": NewErrorHandler("Stage3 Error"),
}

// Main function
func main() {
	// Example data
	data := "   hello! world @ "
	processed, err := ProcessData(data, stageHandlers)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Processed data: %s\n", processed)
	}

	// Example of data that triggers errors in Stage2
	data2 := "invalid data"
	processed, err = ProcessData(data2, stageHandlers)
	if err != nil {
		fmt.Printf("Error (Stage2): %v\n", err)
	}
}

// Test Cases
func TestProcessData(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantErr bool
		stage   string
	}{
		{"Valid Data", "hello world", false, ""},
		{"Empty After Trim", "   ", true, "Stage1"},
		{"Invalid Data Format", "invalid data", true, "Stage2"},
		{"Data Incomplete", "hello*", true, "Stage2"},
		{"Unexpected '@'", "hello @world", true, "Stage3"},
		{"Data Too Long", "a long string exceeding 100 characters" + string(make([]byte, 80)), true, "Stage3"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ProcessData(tc.input, stageHandlers)
			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else {
					pipelineErr, ok := err.(*PipelineError)
					if !ok {
						t.Errorf("Error is not a PipelineError")
					} else if pipelineErr.Stage != tc.stage {
						t.Errorf("Expected error at stage '%s', but got '%s'", tc.stage, pipelineErr.Stage)
					}
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}
