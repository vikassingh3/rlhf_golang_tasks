package main

import (
	"fmt"
)

// Define your custom error types
type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

// Define another custom error type
type AnotherError struct {
	msg string
}

func (e *AnotherError) Error() string {
	return e.msg
}

// A state machine function that returns multiple values and can return errors
func myStateMachine(input int) (int, error) {
	switch input {
	case 0:
		return 1, nil // No error
	case 1:
		return 0, &MyError{"Invalid input: 1"} // Custom error MyError
	case 2:
		return 0, &AnotherError{"Internal server error"} // Custom error AnotherError
	default:
		return 0, fmt.Errorf("Unknown input: %d", input) // Generic error
	}
}

func main() {
	for i := 0; i < 5; i++ {
		result, err := myStateMachine(i)
		if err != nil {
			switch err.(type) {
			case *MyError:
				fmt.Println("Handling MyError:", err.(*MyError).msg)
			case *AnotherError:
				fmt.Println("Handling AnotherError:", err.(*AnotherError).msg)
			default:
				fmt.Println("Handling generic error:", err)
			}
		} else {
			fmt.Println("Input:", i, ", Result:", result)
		}
	}
}
