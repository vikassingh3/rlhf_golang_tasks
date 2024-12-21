package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

type Data struct {
	Title string
	Name  string
}

// RenderTemplate renders the specified template with the given data.
// It returns the rendered output as a string and any error that occurred.
func RenderTemplate(templatePath string, data interface{}) (string, error) {
	// Create a new template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %v", err)
	}

	// Execute the template with the provided data
	var output bytes.Buffer // Use bytes.Buffer as it implements io.Writer
	err = tmpl.Execute(&output, data)
	if err != nil {
		return "", fmt.Errorf("error executing template: %v", err)
	}

	return output.String(), nil // Convert buffer to string and return
}

func main() {
	// Define the template file path
	templatePath := "index.html"

	// Create a Data struct with the data to be passed to the template
	data := Data{
		Title: "Welcome Page",
		Name:  "John Doe",
	}

	// Render the template
	renderedOutput, err := RenderTemplate(templatePath, data)
	if err != nil {
		fmt.Println("Failed to render template:", err)
		os.Exit(1)
	}

	// Output the rendered HTML
	fmt.Println(renderedOutput)
}
