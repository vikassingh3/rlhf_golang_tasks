package plugins

type Plugin interface {
	Name() string
	Process(data string) string
}
