package plugin1

type Plugin struct{}

func (p *Plugin) Greet() string {
    return "Hello from Plugin1!"
}

func (p *Plugin) Farewell() string {
    return "Goodbye from Plugin1!"
}
