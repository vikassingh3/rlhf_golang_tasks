package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	filePath := "person.json"
	person, err := readJSON(filePath)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}


package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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
	filePath := "person.xml"
	person, err := readXML(filePath)
	if err != nil {
		fmt.Printf("Error reading XML file: %v\n", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}


package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
)

func readCSV(filePath string) ([][]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func main() {
	filePath := "people.csv"
	records, err := readCSV(filePath)
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return
	}

	for _, record := range records {
		fmt.Printf("Name: %s, Age: %s\n", record[0], record[1])
	}
}