package main

import (
	"errors"
	"strings"
)


func Stage1(data string) (string, error) {
	if data == "" {
		return "", errors.New("Data is empty.")
	}
	if strings.Contains(data, "!") {
		return "", errors.New("Unexpected '!' in data.")
	}
	return data, nil
}