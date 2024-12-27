package main

import (
	"fmt"
	"reflect"
)

// Define a sample struct
type User struct {
	ID      int
	Name    string
	Age     int
	Email   string
	Active  bool
}

// Function to dynamically format and display a struct using reflection
func formatStruct(structValue interface{}) string {
	// Create a reflect.Value from the struct interface
	rv := reflect.ValueOf(structValue)

	// Check if it is a struct
	if rv.Kind() != reflect.Struct {
		return "Invalid input: not a struct"
	}

	// Prepare a string builder
	var result string

	// Iterate over each field in the struct
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)

		// Get the field name
		fieldName := rv.Type().Field(i).Name

		// Get the field value
		fieldValue := field.String()

		// Format each field name and value pair
		result += fmt.Sprintf("%s: %s\n", fieldName, fieldValue)
	}

	return result
}

func main() {
	// Create a User struct instance
	user := User{
		ID:      1,
		Name:    "John Doe",
		Age:     30,
		Email:   "johndoe@example.com",
		Active:  true,
	}

	// Format and display the User struct
	formattedOutput := formatStruct(user)
	fmt.Println(formattedOutput)
}