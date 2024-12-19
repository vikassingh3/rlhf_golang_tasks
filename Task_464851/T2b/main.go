package main

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path"
)

// MAX_UPLOAD_SIZE defines the maximum allowed file size for uploads
const MAX_UPLOAD_SIZE = 1024 * 1024 * 10 // 10MB

// AllowTypes defines a list of allowed file types
var AllowTypes = []string{"image/png", "image/jpeg", "application/pdf"}

type FileHandler struct {
	uploadsDir   string
	downloadsDir string
}

func NewFileHandler(uploadsDir, downloadsDir string) *FileHandler {
	return &FileHandler{uploadsDir, downloadsDir}
}

func (fh *FileHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the uploads directory exists
	os.MkdirAll(fh.uploadsDir, 0755)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, err.Error(), http.StatusRequestEntityTooLarge)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Validate file type
	fileContentType := mime.TypeByExtension(path.Ext(handler.Filename))
	if !contains(AllowTypes, fileContentType) {
		http.Error(w, "Invalid file type", http.StatusUnsupportedMediaType)
		return
	}

	// Validate file size
	fileSize, err := file.Seek(0, os.SEEK_END)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if fileSize > MAX_UPLOAD_SIZE {
		http.Error(w, "File size too large", http.StatusRequestEntityTooLarge)
		return
	}

	fileName := path.Clean(handler.Filename)
	filePath := path.Join(fh.uploadsDir, fileName)

	dest, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	_, err = io.Copy(dest, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "File uploaded successfully!\n")
}

// contains checks if a string slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	// Rest of the code remains the same
}
