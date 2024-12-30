package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Example DB connection
	fmt.Print("Enter a command: ")
	var command string
	fmt.Scanln(&command)
	command = strings.TrimSpace(command)

	// Example: safely executing a command with arguments
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}