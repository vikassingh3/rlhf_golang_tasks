package main

import (
	"fmt"
	"os"

	"context-management/examples"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <example_name>")
		return
	}

	exampleName := os.Args[1]
	switch exampleName {
	case "pitfall1":
		examples.Pitfall1()
	case "correct_pitfall1":
		examples.CorrectPitfall1()
	case "pitfall2":
		examples.Pitfall2()
	case "correct_pitfall2":
		examples.CorrectPitfall2()
	case "pitfall3":
		examples.Pitfall3()
	case "correct_pitfall3":
		examples.CorrectPitfall3()
	default:
		fmt.Println("Unknown example. Please use one of: pitfall1, correct_pitfall1, pitfall2, correct_pitfall2, pitfall3, correct_pitfall3")
	}
}
