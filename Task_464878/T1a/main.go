package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
)

func loadPlugins(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Execute the plugin script
		cmd := exec.Command(runtime.GOOS+"-"+runtime.GOARCH, filepath.Join(dir, file.Name()))
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		// Read output from the plugin
		output := make([]byte, 1024)
		n, _ := stdout.Read(output)
		fmt.Println(string(output[:n]))

		// Wait for the plugin to complete
		err = cmd.Wait()
		if err != nil {
			log.Println("Error waiting for plugin to complete:", err)
		}
	}
}

func main() {
	loadPlugins("./plugin") // Replace "plugins" with the path to your plugin directory
}
