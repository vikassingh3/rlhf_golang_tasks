package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type PluginMessage struct {
	PluginName string
	Message    string
}

func loadPlugins(dir string) {
	plugins := make([]*exec.Cmd, 0)
	messageCh := make(chan PluginMessage)
	var wg sync.WaitGroup

	// Read the plugin directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the files and start the plugins
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		pluginPath := filepath.Join(dir, file.Name())

		// Start the plugin as a command
		cmd := exec.Command(pluginPath)
		cmd.Env = append(os.Environ(), fmt.Sprintf("MESSAGE_CHANNEL=%s", "1234"))
		// Start the plugin
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		plugins = append(plugins, cmd)

		// Increment the wait group counter
		wg.Add(1)

		// Handle the plugin's messages in a goroutine
		go func(cmd *exec.Cmd, pluginName string) {
			defer wg.Done()
			// Monitor the plugin's output
			handlePluginMessages(cmd, pluginName, messageCh)
		}(cmd, file.Name())
	}

	// Main loop to read messages from plugins
	go func() {
		wg.Wait()
		close(messageCh)
	}()

	for msg := range messageCh {
		fmt.Printf("Received from %s: %s\n", msg.PluginName, msg.Message)
	}
}

func handlePluginMessages(cmd *exec.Cmd, pluginName string, messageCh chan<- PluginMessage) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error reading from plugin %s stdout: %v", pluginName, err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Error with plugin %s: %v", pluginName, err)
	}
	messageCh <- PluginMessage{PluginName: pluginName, Message: "Plugin has finished processing."}

	// Process the plugin's output
	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Printf("Error reading output from plugin %s: %v", pluginName, err)
			break
		}
		messageCh <- PluginMessage{PluginName: pluginName, Message: string(buf[:n])}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go plugin_directory")
	}
	loadPlugins(os.Args[1])
}
