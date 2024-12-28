package main

import (
	"fmt"
	"reflect"
)

// Define a sample struct with nested structs and slices
type User struct {
	ID      int
	Name    string
	Age     int
	Email   string
	Active  bool
	Address Address
	Photos  []Photo
}

type Address struct {
	Street    string
	City      string
	Zip       string
	Country   string
}

type Photo struct {
	URL       string
	Caption   string
}

// Function to dynamically format and display a struct using reflection
func formatStruct(structValue interface{}) string {
	// Create a reflect.Value from the struct interface
	rv := reflect.ValueOf(structValue)

	// Check if it is a struct
	if rv.Kind() != reflect.Struct {
		return "Invalid input: not a struct"
	}

	// Prepare a string builder for efficient concatenation
	var result string

	// Iterate over each field in the struct
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)

		// Get the field name
		fieldName := rv.Type().Field(i).Name

		// Recursively handle structs, slices, and other complex values
		valueString := formatField(field)

		// Format each field name and value pair
		result += fmt.Sprintf("%s: %s\n", fieldName, valueString)
	}

	return result
}

// Recursive function to handle different value kinds
func formatField(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Struct:
		return formatStruct(value.Interface())
	case reflect.Slice:
		return formatSlice(value)
	default:
		return value.String()
	}
}

// Function to format slices
func formatSlice(value reflect.Value) string {
	if value.IsNil() {
		return "nil"
	}

	// Prepare a string builder for efficient concatenation
	var result string

	// Iterate over each element in the slice
	for i := 0; i < value.Len(); i++ {
		// Recursively handle each element
		result += formatField(value.Index(i)) + "\n"
	}

	return result
}

func main() {
	// Create a User struct instance with nested structs and slices
	user := User{
		ID:      1,
		Name:    "John Doe",
		Age:     30,
		Email:   "johndoe@example.com",
		Active:  true,
		Address: Address{
			Street:    "123 Main St",
			City:      "Example City",
			Zip:       "12345",
			Country:   "Example Country",
		},
		Photos: []Photo{
			{URL: "https://example.com/photo1.jpg", Caption: "Photo 1"},
			{URL: "https://example.com/photo2.jpg", Caption: "Photo 2"},
		},
	}

	// Format and display the User struct
	formattedOutput := formatStruct(user)
	fmt.Println(formattedOutput)
}