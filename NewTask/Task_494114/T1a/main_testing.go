package main

import (
	"errors"
	"testing"
)

func TestProcessDataWithError(t *testing.T) {
	dataWithError := "   hello! world @ "

	_, err := ProcessData(dataWithError)

	if err == nil {
		t.Error("Expected an error, got nil")
	} else if !errors.Is(err, &PipelineError{}) {
		t.Error("Expected PipelineError, got", err)
	} else if err.(*PipelineError).Stage != "Stage1" {
		t.Error("Expected error in Stage1, got in stage", err.(*PipelineError).Stage)
	} else if err.(*PipelineError).Err.Error() != "Unexpected '!' in data." {
		t.Error("Expected 'Unexpected ! in data.' error, got", err.(*PipelineError).Err)
	}
}