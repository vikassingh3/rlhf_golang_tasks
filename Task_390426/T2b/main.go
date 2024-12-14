// main.go
package main

import (
	"my_proj/server"
	"my_proj/util"
)

func main() {
	server.StartServer()       // Starts the server
	util.SomeUtilityFunction() // Calls a utility function
}
