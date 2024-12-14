package main

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

// Define a data structure that matches both JSON and XML formats
type Sample struct {
	Name  string `json:"name" xml:"name"`
	Age   int    `json:"age" xml:"age"`
	Email string `json:"email" xml:"email"`
}

// Sample data for JSON and XML
var jsonData = `{"name": "John Doe", "age": 30, "email": "john.doe@example.com"}`
var xmlData = `<Sample><name>John Doe</name><age>30</age><email>john.doe@example.com</email></Sample>`

// Benchmark function for JSON parsing
func BenchmarkJSONParse(b *testing.B) {
	var sample Sample
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal([]byte(jsonData), &sample); err != nil {
			b.Fatalf("JSON Unmarshal failed: %v", err)
		}
	}
}

// Benchmark function for XML parsing
func BenchmarkXMLParse(b *testing.B) {
	var sample Sample
	for i := 0; i < b.N; i++ {
		if err := xml.Unmarshal([]byte(xmlData), &sample); err != nil {
			b.Fatalf("XML Unmarshal failed: %v", err)
		}
	}
}
