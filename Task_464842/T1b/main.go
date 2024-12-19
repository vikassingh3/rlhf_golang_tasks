package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define the structure of the data you expect from the API
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Replace this with the actual API URL that returns an array of users
	apiUrl := "https://api.example.com/users"

	res, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error while working with", err.Error())
		return
	}
	defer res.Body.Close()

	// Read the response body and convert it to a byte slice
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while working with", err.Error())
		return
	}

	// Create a slice to store the user data (empty slice)
	var users []User

	// Unmarshal the JSON data from the byte slice into the users slice
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error while working with", err.Error())
		return
	}

	// Print the data stored in the users slice
	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
