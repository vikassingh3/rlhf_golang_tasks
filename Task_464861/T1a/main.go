package main

import (
	"fmt"
	"html/template"
	"os"
)

type Data struct {
	Title string
	Name  string
}

func main() {
	// Define the template file path
	templatePath := "./index.html"

	fmt.Println(templatePath,"sdkjdkjaskjs")

	// Create a new template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}

	// Create a Data struct with the data to be passed to the template
	data := Data{
		Title: "Welcome Page",
		Name:  "John Doe",
	}

	// Execute the template with the provided data
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}
}
