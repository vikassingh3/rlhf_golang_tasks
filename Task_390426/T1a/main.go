package main

import (
	"fmt"

	"example.com/packageA"
	"example.com/packageB"
)

func main() {
	a := packageA.NewA()
	b := packageB.NewB(a)

	// Use b as needed
	fmt.Println(b)
}
