package main

import (
	"fmt"
	"reflect"
)

// A generic interface for data processing
type DataProcessor interface {
	Process()
}

type StringProcessor struct {
	Data string `process:"true"`
}

func (sp *StringProcessor) Process() {
	fmt.Println("Processing string:", sp.Data)
}

type IntProcessor struct {
	Data int `process:"true"`
}

func (ip *IntProcessor) Process() {
	fmt.Println("Processing int:", ip.Data)
}

func ProcessData(data interface{}) {
	v := reflect.ValueOf(data)

	// Dereference pointer if the input is a pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	// Ensure the type is still a struct after dereferencing
	if t.Kind() != reflect.Struct {
		fmt.Println("Invalid data type. Expecting a struct implementing DataProcessor.")
		return
	}

	// Check if the value implements the DataProcessor interface
	if _, ok := data.(DataProcessor); !ok {
		fmt.Println("Invalid data type. Expecting DataProcessor.")
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("process")

		if tag == "true" {
			fieldValue := v.Field(i)
			switch field.Type.Kind() {
			case reflect.String:
				str := fieldValue.Interface().(string)
				fmt.Println("Field Name:", field.Name, ", Value:", str)
			case reflect.Int:
				intVal := fieldValue.Interface().(int)
				fmt.Println("Field Name:", field.Name, ", Value:", intVal)
			default:
				fmt.Println("Unsupported field type:", field.Type.Kind())
			}
		}
	}
}

func main() {
	sp := &StringProcessor{Data: "Hello, Go!"}
	ip := &IntProcessor{Data: 42}

	ProcessData(sp) // Output: Field Name: Data, Value: Hello, Go!
	ProcessData(ip) // Output: Field Name: Data, Value: 42

	// Type safety check will fail for invalid data
	ProcessData("Invalid Data") // Output: Invalid data type. Expecting DataProcessor.
}
