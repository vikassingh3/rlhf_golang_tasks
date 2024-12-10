package main

import (
	"fmt"
	"os"
)

func main() {
	// Assuming func1.GetData returns error

	data, err := os.OpenFile("data.txt", 2, os.ModeAppend.Perm())

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}
