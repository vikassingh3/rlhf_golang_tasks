package main

import (
    "fmt"
    "reflect"
)

// Function that receives an interface and handles different types with reflection
func ProcessValue(val interface{}) {
    // Get the reflected value and type
    v := reflect.ValueOf(val)
    t := v.Type()

    // Check the type
    if t.Kind() != reflect.Struct {
        fmt.Println("Invalid input type:", t)
        return
    }

    // Use a type switch to ensure type safety
    switch val := val.(type) {
    case struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }:
        // Perform operations on the known struct
        fmt.Println("Name:", val.Name, ", Age:", val.Age)
    default:
        fmt.Println("Unsupported struct type:", t)
    }
}

func main() {
    // Process a struct value
    person := struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }{"Alice", 30}
    ProcessValue(person)
}