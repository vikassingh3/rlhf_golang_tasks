package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/gorilla/mux"
)

// Dummy in-memory storage for files
var files []string

// addFile is the handler for uploading a file
func addFile(w http.ResponseWriter, r *http.Request) {
	// Simple logic for adding file
	filename := r.FormValue("filename")
	files = append(files, filename)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s uploaded", filename)
}

// listFiles is the handler for listing files
func listFiles(w http.ResponseWriter, r *http.Request) {
	// Simple logic for listing files
	userID := r.URL.Query().Get("user_id")
	// For simplicity, just return all files, we could filter based on userID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Files for user %s: %v", userID, files)
}

// TestLargeFileUpload checks if the service can handle a large file upload
func TestLargeFileUpload(t *testing.T) {
	// Generate large file data
	largeFileData := make([]byte, 10*1024*1024) // 10 MB

	// Create a new test HTTP server using Gorilla Mux Router
	router := mux.NewRouter()
	router.HandleFunc("/upload", addFile)
	server := httptest.NewServer(router)
	defer server.Close()

	// Set up HTTP client
	client := &http.Client{}

	// Send large file upload request
	resp, err := client.Post(server.URL+"/upload", "application/x-www-form-urlencoded", strings.NewReader(fmt.Sprintf("filename=large_file.txt&user_id=1&data=%s", largeFileData)))
	if err != nil {
		t.Fatalf("Failed to upload large file: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status: %d", resp.StatusCode)
	}
}

// TestConcurrentAccess checks for consistent results when multiple users access the same file
func TestConcurrentAccess(t *testing.T) {
	// Assume user1 uploaded file 'file1.txt' with ID 1
	// fileID := 1

	// Create a new test HTTP server using Gorilla Mux Router
	router := mux.NewRouter()
	router.HandleFunc("/files", listFiles)
	server := httptest.NewServer(router)
	defer server.Close()

	// Set up HTTP client
	client := &http.Client{}

	// Launch concurrent clients to access the same file
	numGoRoutines := 100 // You can adjust the number of clients
	var wg sync.WaitGroup

	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Access the same file as user1
			resp, err := client.Get(server.URL + fmt.Sprintf("/files?user_id=1"))
			if err != nil {
				t.Errorf("Concurrent access failed: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Invalid response status: %d", resp.StatusCode)
			}
			resp.Body.Close()
		}()
	}

	wg.Wait()

	// Ensure all concurrent accesses were successful
	if t.Failed() {
		t.Error("Concurrent access test failed")
	}
}
