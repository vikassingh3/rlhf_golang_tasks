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

package main

import (
	"errors"
	"strings"
)

// Stage1 is the first step in the pipeline.
func Stage1(data string) (string, error) {
	if strings.Contains(data, "!") {
		return data, errors.New("Unexpected '!' in data.")
	}
	return strings.TrimSpace(data), nil
}