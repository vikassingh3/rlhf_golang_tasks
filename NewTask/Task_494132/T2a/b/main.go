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

    if t.Kind() != reflect.Struct {
        fmt.Println("Invalid input type:", t)
        return
    }

    // Type assertion for safety
    if person, ok := val.(struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }); ok {
        fmt.Println("Name:", person.Name, ", Age:", person.Age)
    } else {
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