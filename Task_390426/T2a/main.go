package main

import (
	"example.com/project/packageA"
	"example.com/project/packageB"
)

func main() {
	a := packageA.NewA()  // Create an instance of A
	b := packageB.NewB(a) // Pass the instance of A to B
	b.A.DoSomething()     // Call a method on A through the interface
}
