package plugins

import (
	"fmt"
	"pugin/manager"
)

func main() {
	pm := manager.NewPluginManager()
	pluginPath := "./plugin/uppercase/uppercase.so" // Replace with your platform-specific plugin extension
	pm.Load(pluginPath)

	data := "hello, plugin system!"
	result := pm.Process(data)

	fmt.Println("Original data:", data)
	fmt.Println("Processed data:", result)
}
