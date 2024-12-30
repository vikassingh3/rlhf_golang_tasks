package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Open a file for reading
	file, err := os.Open("./example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()  // Close the file handle automatically

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	// Make a network request
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()  // Close the response body automatically

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}