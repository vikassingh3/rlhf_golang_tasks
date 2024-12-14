package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
)

// Sample struct to hold our data
type Sample struct {
	Name  string `json:"name" xml:"name"`
	Age   int    `json:"age" xml:"age"`
	Email string `json:"email" xml:"email"`
}

// Sample JSON and XML data
var jsonData = `{"name": "John Doe", "age": 30, "email": "john.doe@example.com"}`
var xmlData = `<Sample><name>John Doe</name><age>30</age><email>john.doe@example.com</email></Sample>`

// Benchmark for JSON parsing
func BenchmarkJSONParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sample Sample
		if err := json.Unmarshal([]byte(jsonData), &sample); err != nil {
			b.Fatalf("JSON Unmarshal failed: %s", err)
		}
	}
}

// Benchmark for XML parsing
func BenchmarkXMLParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sample Sample
		if err := xml.Unmarshal([]byte(xmlData), &sample); err != nil {
			b.Fatalf("XML Unmarshal failed: %s", err)
		}
	}
}

func main() {
	// Print a message indicating how to run the benchmarks
	fmt.Println("Run the benchmarks using 'go test -bench=.' in the command line.")
}

// Note: The `testing` package requires tests to be executed with `go test`, 
// so the benchmarks cannot be run directly from the `main` function.