package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

const (
	uploadsDir       = "uploads"
	downloadsDir     = "downloads"
	maxFileSize      = 1024 * 1024 * 10                       // 10MB limit
	allowedFileTypes = "image/jpeg,image/png,application/pdf" // Allow JPEG, PNG, and PDF
)

// FileHandler is a struct to handle file-related operations
type FileHandler struct {
	uploadsDir   string
	downloadsDir string
}

// NewFileHandler creates a new FileHandler
func NewFileHandler(uploadsDir, downloadsDir string) *FileHandler {
	return &FileHandler{uploadsDir, downloadsDir}
}

// UploadHandler handles file uploads with validation
func (fh *FileHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the uploads directory exists
	os.MkdirAll(fh.uploadsDir, 0755)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(1024 * 1024 * 10) // Max memory for form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Check file size
	seeker, ok := file.(io.Seeker)
	if !ok {
		http.Error(w, "Unable to determine file size", http.StatusInternalServerError)
		return
	}

	// Seek to the end to determine file size
	fileSize, err := seeker.Seek(0, io.SeekEnd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Reset the file pointer to the beginning
	_, err = seeker.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fileSize > maxFileSize {
		http.Error(w, "File size exceeds limit", http.StatusBadRequest)
		return
	}

	// Check file type
	fileType := handler.Header["Content-Type"][0]
	if !strings.Contains(allowedFileTypes, fileType) {
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
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

// DownloadHandler handles file downloads
func (fh *FileHandler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["filename"]

	filePath := path.Join(fh.downloadsDir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Close()

	http.ServeFile(w, r, filePath)
}

func main() {
	r := mux.NewRouter()
	fh := NewFileHandler(uploadsDir, downloadsDir)

	r.HandleFunc("/upload", fh.UploadHandler).Methods(http.MethodPost)
	r.HandleFunc("/download/{filename}", fh.DownloadHandler).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
