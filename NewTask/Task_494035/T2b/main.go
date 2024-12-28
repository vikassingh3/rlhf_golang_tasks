package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
)

var (
	fieldCache sync.Map // Cache for field lookups by struct type
)

// User struct for demonstration
type User struct {
	ID   int
	Name string
}

// GetFieldByName retrieves the field by name, caching it for future access.
func getFieldByName(t reflect.Type, name string) *reflect.StructField {
	// Check cache for field lookup
	v, ok := fieldCache.Load(t)
	if !ok {
		// If not in cache, perform reflection to get the field by name
		f, ok := t.FieldByName(name)
		if !ok {
			panic(fmt.Sprintf("field %q not found in %v", name, t))
		}
		// Store the pointer to reflect.StructField in the cache for future use
		fieldCache.Store(t, &f)
		v = &f
	}

	// Return the cached field
	return v.(*reflect.StructField)
}

func formatStructOptimized(structValue interface{}) string {
	var result bytes.Buffer

	// Use reflection to access the struct fields dynamically
	rv := reflect.ValueOf(structValue)
	rt := rv.Type()

	// Example using field structs to reduce repeated reflection calls
	fieldID := getFieldByName(rt, "ID")
	fieldName := getFieldByName(rt, "Name")

	// Efficiently access and format fields using the cached struct fields
	result.WriteString(fmt.Sprintf("ID: %d\n", rv.FieldByIndex(fieldID.Index).Interface()))
	result.WriteString(fmt.Sprintf("Name: %s\n", rv.FieldByIndex(fieldName.Index).Interface()))

	return result.String()
}

func main() {
	// Create a User instance
	user := User{ID: 1, Name: "Alice"}

	// Use the optimized function to format the struct
	result := formatStructOptimized(user)

	// Output the result
	fmt.Println(result)
}
