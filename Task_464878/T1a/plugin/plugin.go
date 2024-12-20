package main

import "fmt"

type MyPlugin struct{}

func (p *MyPlugin) Greet() string {
	return "Hello, from plugin1!"
}

func (p *MyPlugin) Farewell() string {
	return "Goodbye, from plugin1!"
}

func main() {
	p := &MyPlugin{}
	fmt.Println(p.Greet())
	fmt.Println(p.Farewell())
}
