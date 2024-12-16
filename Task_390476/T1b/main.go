package main

import (
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	// Creating a matcher
	matcher := language.NewMatcher([]language.Tag{
		language.English,
		language.French,
		language.Spanish,
	})

	// Get user's preferred language
	userLang := os.Getenv("LANG")
	tag, _ := language.MatchStrings(matcher, userLang)

	// Create a Printer for the matched language
	p := message.NewPrinter(tag)

	// Register translations
	message.SetString(language.English, "Hello, %s!", "Hello, %s!")
	message.SetString(language.French, "Hello, %s!", "Bonjour, %s!")
	message.SetString(language.Spanish, "Hello, %s!", "¡Hola, %s!")

	// Print localized messages
	name := "John"
	p.Printf("Hello, %s!\n", name)
}
