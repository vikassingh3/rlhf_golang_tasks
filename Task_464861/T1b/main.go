package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type TemplateData struct {
	Title   string
	Message string
	Now     time.Time
}

func renderTemplate(templateFile string, data interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	data := TemplateData{
		Title:   "Welcome to My Website!",
		Message: "Hello, this is a dynamic message!",
		Now:     time.Now(),
	}

	renderTemplate("./index.html", data)
}  