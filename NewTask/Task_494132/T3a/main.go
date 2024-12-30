package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Function that returns different types, including an error
func DynamicOperation(i int) interface{} {
	switch i {
	case 0:
		return "string"
	case 1:
		return 42
	case 2:
		return 3.14
	case 3:
		return errors.New("operation failed")
	default:
		return nil
	}
}

// Function to process a dynamically returned value using reflection
func ProcessDynamicValue(result interface{}) error {
	v := reflect.ValueOf(result)

	if !v.IsValid() || v.IsZero() {
		fmt.Println("Nil returned")
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		fmt.Println("String returned:", v.String())
	case reflect.Int:
		fmt.Println("Integer returned:", v.Int())
	case reflect.Float64:
		fmt.Println("Float64 returned:", v.Float())
	default:
		// Handle the error case separately
		if err, ok := result.(error); ok {
			fmt.Println("Error returned:", err.Error())
			return err // Propagate the error
		}
		return fmt.Errorf("unsupported return type: %v", v.Kind())
	}
	return nil
}

func main() {
	for i := 0; i <= 4; i++ {
		result := DynamicOperation(i)
		if err := ProcessDynamicValue(result); err != nil {
			fmt.Println("Error handling failed:", err.Error())
		}
	}
}
