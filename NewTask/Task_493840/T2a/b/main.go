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

	// Safe way to execute commands
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}