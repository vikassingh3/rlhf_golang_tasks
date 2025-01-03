package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// City struct holds data about a city, including Density
type City struct {
	Name       string
	Population int
	Area       float64
	Density    int // Added the Density field
}

func main() {
	const filePath = "./cities.csv" // Assume the CSV file is named "cities.csv"

	// Calculate population density using a slice
	start := time.Now()
	citiesSlice := []City{}
	calculateDensityWithSlice(filePath, &citiesSlice)
	sliceElapsed := time.Since(start)
	fmt.Printf("Time taken for slice: %s\n", sliceElapsed)

	// Calculate population density using an array
	start = time.Now()
	const maxCities = 1000 // Example maximum number of cities
	citiesArray := [maxCities]City{}
	calculateDensityWithArray(filePath, &citiesArray)
	arrayElapsed := time.Since(start)
	fmt.Printf("Time taken for array: %s\n", arrayElapsed)
}

func calculateDensityWithSlice(filePath string, cities *[]City) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ','

	var row []string
	for {
		if row, err = reader.Read(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if len(row) != 3 {
			continue // Skip rows with incorrect number of columns
		}

		city := City{
			Name:     row[0],
			Population: int(parseInt(row[1])),
			Area:      parseFloat64(row[2]),
		}

		*cities = append(*cities, city)
	}

	// Calculate density after reading the data
	for i := range *cities {
		(*cities)[i].Density = (*cities)[i].Population / int((*cities)[i].Area) // Density = Population / Area
	}
}

func calculateDensityWithArray(filePath string, cities *[1000]City) { // maxCities defined here
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ','

	var row []string
	var cityIndex int
	for {
		if row, err = reader.Read(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if len(row) != 3 {
			continue // Skip rows with incorrect number of columns
		}

		city := City{
			Name:     row[0],
			Population: int(parseInt(row[1])),
			Area:      parseFloat64(row[2]),
		}

		cities[cityIndex] = city
		cityIndex++
	}

	// Calculate density after reading the data
	for i := 0; i < cityIndex; i++ {
		cities[i].Density = cities[i].Population / int(cities[i].Area) // Density = Population / Area
	}
}

func parseInt(s string) int {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("Error parsing int:", err)
		return 0
	}
	return int(v)
}

func parseFloat64(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("Error parsing float64:", err)
		return 0.0
	}
	return v
}
