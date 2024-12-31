package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type Document struct {
	Content string `json:"content"`
	Version int    `json:"version"`
}

var fileLock sync.Mutex

func main() {
	filePath := "example.txt"

	// Load the initial document
	doc, err := loadDocument(filePath)
	if err != nil {
		fmt.Printf("Error loading document: %v\n", err)
		return
	}

	fmt.Printf("Current content: %s\nVersion: %d\n", doc.Content, doc.Version)

	// Simulate a user edit
	newContent := "Updated content"
	newVersion := doc.Version + 1

	// Save the edited document with optimistic locking
	err = saveDocument(filePath, newContent, newVersion, doc.Version)
	if err != nil {
		fmt.Printf("Error saving document: %v\n", err)
	} else {
		fmt.Println("Document saved successfully.")
	}
}

func loadDocument(filePath string) (*Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Document{Content: "", Version: 0}, nil
		}
		return nil, err
	}
	defer file.Close()

	doc := &Document{}
	if err := json.NewDecoder(file).Decode(doc); err != nil && err != io.EOF {
		return nil, err
	}

	return doc, nil
}

func saveDocument(filePath, newContent string, newVersion, currentVersion int) error {
	fileLock.Lock()
	defer fileLock.Unlock()

	doc, err := loadDocument(filePath)
	if err != nil {
		return err
	}

	// Check for version conflict
	if doc.Version != currentVersion {
		return fmt.Errorf("version conflict detected: expected %d, got %d", currentVersion, doc.Version)
	}

	// Update and save the document
	doc.Content = newContent
	doc.Version = newVersion

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(doc)
}
