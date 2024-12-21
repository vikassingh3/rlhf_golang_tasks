package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func ParseCoordinates(input string) (float64, float64, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Invalid input format: expected 'lat,long'")
	}

	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid latitude: %w", err)
	}

	lon, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid longitude: %w", err)
	}

	return lat, lon, nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, fmt.Errorf("Error retrieving user: %w", err)
	}
	return &user, nil
}

func StringToUpperAndLower(input string) (string, string, error) {
	if len(input) > 10 {
		return "", "", errors.New("Input string is too long")
	}
	return strings.ToUpper(input), strings.ToLower(input), nil
}

func main() {
	inputCoordinates := "37.7749,-122.4194"
	lat, lon, err := ParseCoordinates(inputCoordinates)
	if err != nil {
		log.Printf("Error parsing coordinates: %v\n", err)
	} else {
		fmt.Printf("Parsed coordinates: Latitude = %.4f, Longitude = %.4f\n", lat, lon)
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	defer db.Close()

	user, err := GetUserByID(db, 1)
	if err != nil {
		log.Printf("Error retrieving user: %v\n", err)
	} else {
		fmt.Printf("Retrieved user: %+v\n", user)
	}

	inputString := "HelloGo"
	upper, lower, err := StringToUpperAndLower(inputString)
	if err != nil {
		log.Printf("Error transforming string: %v\n", err)
	} else {
		fmt.Printf("Original: %s, Upper: %s, Lower: %s\n", inputString, upper, lower)
	}
}
