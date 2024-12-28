package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func readJSON(filePath string) (*Person, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var person Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func main() {
	filePath := "./person.json"
	person, err := readJSON(filePath)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
