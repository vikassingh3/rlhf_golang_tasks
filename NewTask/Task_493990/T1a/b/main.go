package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func readXML(filePath string) (*Person, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var person Person
	err = xml.Unmarshal(data, &person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func main() {
	filePath := "./person.xml"
	person, err := readXML(filePath)
	if err != nil {
		fmt.Printf("Error reading XML file: %v\n", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
