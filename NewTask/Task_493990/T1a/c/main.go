package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
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
	filePath := "./people.csv"
	records, err := readCSV(filePath)
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return
	}

	for _, record := range records {
		fmt.Printf("Name: %s, Age: %s\n", record[0], record[1])
	}
}