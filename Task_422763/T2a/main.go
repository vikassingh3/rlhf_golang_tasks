// main.go in your application directory
package main

import (
	"fmt"

	"myapp/myutils"
)

func main() {
	input := "Hello, world!"
	reversed := myutils.Reverse(input)
	fmt.Println("Reversed string:", reversed) // !dlrow ,olleH"
}
