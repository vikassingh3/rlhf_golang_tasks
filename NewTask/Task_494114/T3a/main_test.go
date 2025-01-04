package main

import (
	"errors"
	"strings"
	"testing"
)

// Stage3 processes the input data and validates multiple conditions.
func Stage3(data string) (string, error) {
	if strings.Contains(data, "@") {
		return data, errors.New("Unexpected '@' in data.")
	}

	if len(data) >= 100 {
		return data, errors.New("Data length exceeds 100 characters.")
	}

	if strings.ContainsAny(data, "0123456789") {
		return data, errors.New("Unexpected numbers in data.")
	}

	return data, nil
}

// TestStage3 tests the Stage3 function for various input scenarios.
func TestStage3(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		wantOutput string
		wantErr    bool
		wantErrMsg string
	}{
		{"Valid input", "hello world", "hello world", false, ""},
		{"Input with '@'", "hello@world", "hello@world", true, "Unexpected '@' in data."},
		{"Input too long", strings.Repeat("a", 101), strings.Repeat("a", 101), true, "Data length exceeds 100 characters."},
		{"Input with numbers", "hello 123", "hello 123", true, "Unexpected numbers in data."},
		{"Input with multiple errors", "hello 123@world", "hello 123@world", true, "Unexpected '@' in data."},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := Stage3(tc.input)
			if tc.wantErr {
				if err == nil || err.Error() != tc.wantErrMsg {
					t.Errorf("Expected error: %v, got: %v", tc.wantErrMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if output != tc.wantOutput {
					t.Errorf("Expected output: '%s', got: '%s'", tc.wantOutput, output)
				}
			}
		})
	}
}
