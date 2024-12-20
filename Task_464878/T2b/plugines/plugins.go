package main

import (
	"fmt"
	"time"
)

func main() {
	pluginName := "plugin1"

	// Print a greeting message
	fmt.Printf("Plugin %s started.\n", pluginName)

	// Simulate some work in the plugin (e.g., processing, network call)
	time.Sleep(2 * time.Second)

	// Send a message back to the main program
	fmt.Printf("Plugin %s is finished processing.\n", pluginName)

	// Simulate sending periodic updates to the main program
	for i := 0; i < 3; i++ {
		fmt.Printf("Plugin %s: Processing step %d\n", pluginName, i+1)
		time.Sleep(1 * time.Second)
	}

	// End the plugin
	fmt.Printf("Plugin %s: Finished work and exiting.\n", pluginName)
}
