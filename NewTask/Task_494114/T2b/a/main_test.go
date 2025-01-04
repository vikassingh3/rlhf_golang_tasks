package main

import (
	"testing"
)

func TestStage1_ErrorHandling(t *testing.T) {
    testCases := []struct {
        name   string
        input  string
        wantErr bool
        wantErrMsg string
    }{
        {"Empty input", "", true, "Data is empty."},  // The specific error message can be changed to align with the later implementation
        {"Input with '!'", "hello!", true, "Unexpected '!' in data."},
        {"Trimmed input with '!'", " hello! ", true, "Unexpected '!' in data."},
        {"Valid input", "hello world", false, ""},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            processed, err := Stage1(tc.input)

            if tc.wantErr {
                if err == nil {
                    t.Errorf("Stage1() expected error, got nil")
                } else if err.Error() != tc.wantErrMsg {
                    t.Errorf("Stage1() expected error message '%s', got '%s'", tc.wantErrMsg, err.Error())
                }
            } else {
                if err != nil {
                    t.Errorf("Stage1() unexpected error: %v", err)
                }
                if processed != tc.input {
                    t.Errorf("Stage1() returned incorrect data. Expected: '%s', Got: '%s'", tc.input, processed)
                }
            }
        })
    }
}  