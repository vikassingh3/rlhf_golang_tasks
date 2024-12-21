package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"time"
)

type Data struct {
	Title string
	Name  string
}

// Custom error types
type TemplateParsingError struct {
	err error
}

type TemplateExecutionError struct {
	err error
}

type IOError struct {
	err error
}

func (tpe *TemplateParsingError) Error() string {
	return fmt.Sprintf("template parsing error: %v", tpe.err)
}

func (tee *TemplateExecutionError) Error() string {
	return fmt.Sprintf("template execution error: %v", tee.err)
}

func (ioe *IOError) Error() string {
	return fmt.Sprintf("IO error: %v", ioe.err)
}

// RenderTemplate renders the template with the given data and writes the output to the specified writer.
// It returns the rendered content and any error that occurred during rendering.
func RenderTemplate(templatePath string, data interface{}, w io.Writer) ([]byte, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, &TemplateParsingError{err: err}
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, &TemplateExecutionError{err: err}
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return nil, &IOError{err: err}
	}

	return buf.Bytes(), nil
}

func main() {
	templatePath := "index.html"
	data := Data{
		Title: "Welcome Page",
		Name:  "Alice",
	}

	var renderedOutput []byte
	var err error

	// Retry rendering the template for 3 times if a TemplateParsingError occurs
	for i := 0; i < 3; i++ {
		renderedOutput, err = RenderTemplate(templatePath, data, os.Stdout)
		if err == nil {
			break
		}

		if _, ok := err.(*TemplateParsingError); ok {
			log.Printf("Template parsing error occurred, retrying in 5 seconds... (%v/3)", i+1)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		// Handle other error types (TemplateExecutionError, IOError)
		switch err.(type) {
		case *TemplateExecutionError:
			log.Printf("Template execution error: %v", err)
			os.Exit(1)
		case *IOError:
			log.Printf("IO error: %v", err)
			os.Exit(1)
		default:
			log.Printf("Unhandled error: %v", err)
			os.Exit(1)
		}
	}

	// Do something with the rendered output if needed
	fmt.Println("\nRendered Output:")
	fmt.Println(string(renderedOutput))
}
