package manager

// Callback interface defines the methods that plugins must implement.
type Callback interface {
    Greet() string
    Farewell() string
}

type PluginManager struct {
    plugins map[string]Callback
}

func NewPluginManager() *PluginManager {
    return &PluginManager{plugins: make(map[string]Callback)}
}

// RegisterPlugin registers a plugin with the manager.
func (pm *PluginManager) RegisterPlugin(name string, cb Callback) {
    pm.plugins[name] = cb
}

// GreetAll invokes the Greet method of all registered plugins.
func (pm *PluginManager) GreetAll() []string {
    greetings := []string{}
    for _, cb := range pm.plugins {
        greetings = append(greetings, cb.Greet())
    }
    return greetings
}

// FarewellAll invokes the Farewell method of all registered plugins.
func (pm *PluginManager) FarewellAll() []string {
    farewells := []string{}
    for _, cb := range pm.plugins {
        farewells = append(farewells, cb.Farewell())
    }
    return farewells
}
