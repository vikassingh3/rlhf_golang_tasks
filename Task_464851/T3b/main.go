package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	uploadDir        = "uploads"
	maxFileSize      = 1024 * 1024 * 10 // 10MB limit
	allowedFileTypes = "image/jpeg,image/png,application/pdf"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form data
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type
	fileType := handler.Header.Get("Content-Type")
	if !strings.Contains(allowedFileTypes, fileType) {
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}

	// Create the upload directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a new file path with the original file name
	filePath := filepath.Join(uploadDir, handler.Filename)

	// Create a new file on the server
	f, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy the file content to the server
	if _, err := io.Copy(f, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}
