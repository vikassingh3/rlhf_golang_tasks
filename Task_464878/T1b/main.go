package main

import (
	"fmt"
	"pugin/manager"
)

func main() {
	pm := manager.NewPluginManager()
	pluginPath := "./plugin/upercase/uper.so" // Path to your compiled plugin
	pm.Load(pluginPath)

	data := "hello, plugin system!"
	result := pm.Process(data)

	fmt.Println("Original data:", data)
	fmt.Println("Processed data:", result)
}
