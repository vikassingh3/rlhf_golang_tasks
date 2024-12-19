package main

import (
	"fmt"
)

func main() {
	// Example of a range loop over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Example of a range loop over a map
	people := map[string]int{"Alice": 30, "Bob": 25}
	for name, age := range people {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Example of a range loop with index and value
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Example of an infinite loop with range over a string
	// This will result in an error with `staticcheck`
	for _, char := range "infinite" {
		fmt.Println(char)
	}
}