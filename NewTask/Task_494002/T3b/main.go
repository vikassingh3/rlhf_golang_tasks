package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const filePath = "./cities.txt" // Replace this with the path to a text file containing a list of city names

func main() {
	// Process cities using slice
	start := time.Now()
	processCitiesSlice()
	elapsed := time.Since(start)
	fmt.Printf("Time taken using slice: %s\n", elapsed)

	// Process cities using array
	start = time.Now()
	processCitiesArray()
	elapsed = time.Since(start)
	fmt.Printf("Time taken using array: %s\n", elapsed)
}

func processCitiesSlice() {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	cities := strings.Split(string(data), "\n")
	for _, city := range cities {
		fmt.Println(city)
	}
}

func processCitiesArray() {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	cities := make([]string, len(lines))
	for i, line := range lines {
		cities[i] = line
	}

	for _, city := range cities {
		fmt.Println(city)
	}
}