package plugin2

type Plugin struct{}

func (p *Plugin) Greet() string {
    return "Hello from Plugin2!"
}

func (p *Plugin) Farewell() string {
    return "Goodbye from Plugin2!"
}
