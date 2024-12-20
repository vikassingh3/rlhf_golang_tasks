package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"project/plugins/manager"
	"project/plugins/plugin1"
	"project/plugins/plugin2"
)

func loadPlugins(dir string, pm *manager.PluginManager) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("sdcjhbsdajkhbsdh1")
	for _, file := range files {
		fmt.Println("sdcjhbsdajkhbsd2h")
		if file.IsDir() {
			continue // Skip directories
		}

		// Get the plugin's full path and name
		pluginPath := filepath.Join(dir, file.Name())  // Use file.Name() to get the string
		pluginName := file.Name()[:len(file.Name())-3] // Assuming you want to remove a ".go" suffix

		// Execute the plugin file
		cmd := exec.Command("go", "run", pluginPath)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		err = cmd.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// Read output from the plugin
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			output := scanner.Text()
			fmt.Println("Plugin output:", output)

			// Hypothetical registration based on plugin name
			switch pluginName {
			case "plugin1":
				pm.RegisterPlugin(pluginName, &plugin1.Plugin{})
			case "plugin2":
				pm.RegisterPlugin(pluginName, &plugin2.Plugin{})
			default:
				fmt.Printf("Unknown plugin: %s\n", pluginName)
			}
		}

		// Wait for the plugin to complete
		err = cmd.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for plugin to complete:", err)
		}
	}
}

func main() {
	pm := manager.NewPluginManager()
	loadPlugins("./plugins", pm)

	greetings := pm.GreetAll()
	fmt.Println("Greetings from plugins:")
	for _, greeting := range greetings {
		fmt.Println(greeting)
	}

	farewells := pm.FarewellAll()
	fmt.Println("\nFarewells from plugins:")
	for _, farewell := range farewells {
		fmt.Println(farewell)
	}
}
