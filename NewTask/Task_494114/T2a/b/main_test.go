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