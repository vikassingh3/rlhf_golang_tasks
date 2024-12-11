package main

import (
	"fmt"
)

type InvalidInputError struct {
	msg string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.msg)
}

func validateInput(input string) error {
	input = "2222"
	if input == "" {
		return &InvalidInputError{msg: "input cannot be empty"}
	}
	return nil
}

func main() {
	input := ""
	err := validateInput(input)
	if err != nil {
		switch err := err.(type) {
		case *InvalidInputError:
			fmt.Println(err.msg)
		default:
			fmt.Println("An unexpected error occurred:", err)
		}
		return
	}
}
