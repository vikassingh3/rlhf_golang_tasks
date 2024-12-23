package main

import (
	"errors"
	"fmt"
	"runtime"
)

// Define a new struct to hold error details
type ErrorContext struct {
	Err   error
	Stack []string
}

func (ec *ErrorContext) Error() string {
	return fmt.Sprintf("Error: %v, Stack: %v", ec.Err, ec.Stack)
}

// Function to generate a new ErrorContext with the stack trace
func newErrorContext(err error) *ErrorContext {
	var stack []string
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		f := runtime.FuncForPC(pc)
		if f == nil {
			continue
		}
		fn := f.Name()
		stack = append(stack, fmt.Sprintf("%s:%d: %s", file, line, fn))
	}
	return &ErrorContext{Err: err, Stack: stack}
}

func main() {
	err := errors.New("an example error")
	errorContext := newErrorContext(err)
	fmt.Println(errorContext)
}
