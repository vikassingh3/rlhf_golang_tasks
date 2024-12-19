package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Error string `json:"error"`
}

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorResponse{Error: "Method not allowed"})
		return
	}

	// Max memory for form data
	err := r.ParseMultipartForm(1024 * 1024 * 10)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{Error: "Failed to parse form data"})
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}
	defer file.Close()

	// Check file size using io.Seeker to calculate the size
	seeker, ok := file.(io.Seeker)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: "Unable to determine file size"})
		return
	}

	// Seek to the end to determine file size
	fileSize, err := seeker.Seek(0, io.SeekEnd)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}

	// Reset the file pointer to the beginning
	_, err = seeker.Seek(0, io.SeekStart)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}

	// Check if file size exceeds the limit
	if fileSize > maxFileSize {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{Error: "File size exceeds limit"})
		return
	}

	// Check file type
	fileType := handler.Header["Content-Type"][0]
	if !strings.Contains(allowedFileTypes, fileType) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{Error: "Unsupported file type"})
		return
	}

	// Proceed with file saving logic (existing code)
	fileName := path.Clean(handler.Filename)
	filePath := path.Join(fh.uploadsDir, fileName)

	dest, err := os.Create(filePath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}
	defer dest.Close()

	_, err = io.Copy(dest, file)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "File uploaded successfully!\n")
}

func main() {
	r := mux.NewRouter()
	fh := NewFileHandler(uploadsDir, downloadsDir)

	r.HandleFunc("/upload", fh.UploadHandler).Methods(http.MethodPost)
	// r.HandleFunc("/download/{filename}", fh.DownloadHandler).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
