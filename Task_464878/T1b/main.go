package main

import (
	"fmt"
	"pugin/pluginmanager"
)

func main() {
	pm := pluginmanager.NewPluginManager()
	pluginPath := "./plugins/uppercase/uppercase.so" // Path to your compiled plugin
	pm.Load(pluginPath)

	data := "hello, plugin system!"
	result := pm.Process(data)

	fmt.Println("Original data:", data)
	fmt.Println("Processed data:", result)
}
