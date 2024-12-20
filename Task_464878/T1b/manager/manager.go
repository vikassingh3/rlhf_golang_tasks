package manager

import (
	"log"
	"plugin"

	"pugin/plugin"
)

type PluginManager struct {
	plugins []plugins.Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make([]plugins.Plugin, 0),
	}
}

func (pm *PluginManager) Load(pluginPath string) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		log.Printf("Error loading plugin %s: %v", pluginPath, err)
		return
	}

	symbol, err := p.Lookup("Plugin")
	if err != nil {
		log.Printf("Error looking up Plugin symbol in %s: %v", pluginPath, err)
		return
	}

	pluginInstance, ok := symbol.(plugins.Plugin)
	if !ok {
		log.Printf("Plugin symbol in %s does not implement Plugin interface", pluginPath)
		return
	}

	pm.plugins = append(pm.plugins, pluginInstance)
}

func (pm *PluginManager) Process(data string) string {
	result := data
	for _, plugin := range pm.plugins {
		result = plugin.Process(result)
	}
	return result
}
