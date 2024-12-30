package main

import (
	"fmt"
	"reflect"
)

// Define a global map to associate function names with actual functions
var functions = map[string]interface{}{
	"add":      add,
	"multiply": multiply,
	"greet":    greet,
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func callFunction(funcName string, args ...interface{}) (interface{}, error) {
	// Retrieve the function by name
	funcValue, exists := functions[funcName]
	if !exists {
		return nil, fmt.Errorf("function %q not found", funcName)
	}

	// Use reflection to prepare the arguments
	funcVal := reflect.ValueOf(funcValue)
	if len(args) != funcVal.Type().NumIn() {
		return nil, fmt.Errorf("expected %d arguments, got %d", funcVal.Type().NumIn(), len(args))
	}

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		// Check argument types
		if reflect.TypeOf(arg) != funcVal.Type().In(i) {
			return nil, fmt.Errorf("argument %d: expected type %s, got %s", i, funcVal.Type().In(i), reflect.TypeOf(arg))
		}
		in[i] = reflect.ValueOf(arg)
	}

	// Call the function
	results := funcVal.Call(in)

	// Return the first result
	return results[0].Interface(), nil
}

func main() {
	// Call functions dynamically
	result1, err := callFunction("add", 2, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of add(2, 3):", result1)
	}

	result2, err := callFunction("multiply", 2, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of multiply(2, 3):", result2)
	}

	result3, err := callFunction("greet", "Alice")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of greet(\"Alice\"):", result3)
	}
}
