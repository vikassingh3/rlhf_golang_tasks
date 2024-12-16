package localize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"golang.org/x/text/message/toml"
)

var (
	bundles = make(map[string]*message.Bundle)
)

// LoadBundle loads translation files for a given language
func LoadBundle(lang string) {
	// If the bundle is already loaded, return early
	if _, ok := bundles[lang]; ok {
		return
	}

	// Create a new bundle for the specified language
	bundle := message.NewBundle(lang)
	bundles[lang] = bundle

	// Get the absolute path for the translations directory
	dir, err := filepath.Abs("translations")
	if err != nil {
		fmt.Println("Error getting translations directory:", err)
		return
	}

	// Read the files in the translations directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading translations directory:", err)
		return
	}

	// Iterate through the files and load translation files with the `.toml` extension
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".toml") {
			continue
		}

		filePath := filepath.Join(dir, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening translation file:", err)
			continue
		}
		defer f.Close()

		// Decode the TOML file and load it into the bundle's catalog
		var cat catalog.Catalog
		if err := toml.NewDecoder(f).Decode(&cat); err != nil {
			fmt.Println("Error decoding translation file:", err)
			continue
		}

		// Add the loaded catalog to the bundle
		bundle.AddMessages(lang, &cat)
	}
}

// GetText retrieves a translated message for a given key
func GetText(lang string, key string) string {
	LoadBundle(lang)
	if bundle, ok := bundles[lang]; ok {
		return bundle.String(key)
	}
	return fmt.Sprintf("Missing translation for key: %s", key)
}
