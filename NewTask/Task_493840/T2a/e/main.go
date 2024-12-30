package main

import (
	"fmt"
	"html/template"
	"strings"
)

func main() {
	// User input
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	// Sanitize and print
	fmt.Println(template.HTMLEscapeString(name))
}