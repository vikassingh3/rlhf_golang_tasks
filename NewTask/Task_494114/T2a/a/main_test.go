package main

import (
	"testing"
)

func TestStage1WithError(t *testing.T) {
	dataWithError := "hello!"

	_, err := Stage1(dataWithError)

	if err == nil {
		t.Error("Expected an error, got nil")
	} else if err.Error() != "Unexpected '!' in data." {
		t.Errorf("Expected 'Unexpected ! in data.' error, got %v", err)
	}
}
