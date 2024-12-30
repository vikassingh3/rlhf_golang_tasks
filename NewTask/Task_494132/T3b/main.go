package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Some function that returns different types, including an error
func SomeFunction(i int) interface{} {
	switch i {
	case 0:
		return "string"
	case 1:
		return 42
	case 2:
		return 3.14
	default:
		return errors.New("Invalid input")
	}
}

// ProcessDynamicValue processes a dynamically returned value using reflection
func ProcessDynamicValue(value interface{}) error {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		fmt.Println("Type: string, Value:", v.String())
	case reflect.Int:
		fmt.Println("Type: int, Value:", v.Int())
	case reflect.Float64:
		fmt.Println("Type: float64, Value:", v.Float())
	case reflect.Interface:
		// Check if the interface value is an error
		if err, ok := v.Interface().(error); ok {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println("Unexpected interface type:", v.Type())
	default:
		fmt.Println("Unsupported or nil type")
	}
	return nil
}

func main() {
	for i := 0; i < 5; i++ {
		value := SomeFunction(i)
		if err := ProcessDynamicValue(value); err != nil {
			fmt.Println("Error encountered:", err)
		}
	}
}
