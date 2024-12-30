package main

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

type Sample struct {
	Name  string `json:"name" xml:"name"`
	Age   int    `json:"age" xml:"age"`
	Email string `json:"email" xml:"email"`
}

var jsonData = `{"name": "John Doe", "age": 30, "email": "john.doe@example.com"}`
var xmlData = `<Sample><name>John Doe</name><age>30</age><email>john.doe@example.com</email></Sample>`

func BenchmarkJSONParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sample Sample
		if err := json.Unmarshal([]byte(jsonData), &sample); err != nil {
			b.Fatalf("JSON Unmarshal failed: %s", err)
		}
	}
}

func BenchmarkXMLParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sample Sample
		if err := xml.Unmarshal([]byte(xmlData), &sample); err != nil {
			b.Fatalf("XML Unmarshal failed: %s", err)
		}
	}
}

func main() {
	// Running benchmarks via command line
}