package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID      int
	Name    string
	Address *Address
	Friends []string
}

type Address struct {
	Street string
	City   string
}

// Recursive function to format nested structs and slices
func formatStructRecursive(structValue interface{}, indent string) string {
	rv := reflect.ValueOf(structValue)
	rt := rv.Type()

	var result strings.Builder

	// If it's not a struct or a pointer to a struct, return immediately
	if rv.Kind() != reflect.Struct && rv.Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", structValue) // Basic types
	}

	// Handle struct fields recursively
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return "nil"
	}

	// If it's a pointer, dereference it
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rv.Type()
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldName := rt.Field(i).Name
		var fieldValue string

		// Handle pointers and interfaces
		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface {
			if field.IsNil() {
				fieldValue = "nil"
			} else {
				fieldValue = formatStructRecursive(field.Elem().Interface(), indent+"  ")
			}
		} else if field.Kind() == reflect.Slice {
			// Handle slices
			if field.Len() == 0 {
				fieldValue = "[]"
			} else {
				sliceElements := make([]string, field.Len())
				for j := 0; j < field.Len(); j++ {
					sliceElements[j] = formatStructRecursive(field.Index(j).Interface(), indent+"  ")
				}
				fieldValue = fmt.Sprintf("[%s]", strings.Join(sliceElements, ", "))
			}
		} else if field.Kind() == reflect.String {
			// Handle string type
			fieldValue = field.String()
		} else if field.Kind() == reflect.Int {
			// Handle int type
			fieldValue = fmt.Sprintf("%d", field.Int())
		} else {
			// Handle other types (e.g., float, bool, etc.)
			fieldValue = fmt.Sprintf("%v", field.Interface())
		}

		result.WriteString(fmt.Sprintf("%s%s: %s\n", indent, fieldName, fieldValue))
	}

	return result.String()
}

func main() {
	user := User{
		ID:   1,
		Name: "Alice",
		Address: &Address{
			Street: "Main Street",
			City:   "New York",
		},
		Friends: []string{"Bob", "Charlie"},
	}

	formattedOutput := formatStructRecursive(user, "")
	fmt.Println(formattedOutput)
}
