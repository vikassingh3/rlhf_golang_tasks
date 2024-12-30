package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {  
    fmt.Print("Enter a command to run: ")
    var command string
    fmt.Scanln(&command)

    // Split the command into arguments
    args := strings.Fields(command)

    // Run the command using exec.Command
    cmd := exec.Command(args[0], args[1:]...)
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Output:", string(output))
}  
  