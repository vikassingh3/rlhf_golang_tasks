package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

type File struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"filename"`
}

func TestFileSharing(t *testing.T) {
	// Create a mock HTTP server
	server := http.NewServeMux()
	server.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// Mock the upload response
		file := File{ID: 1, UserID: r.FormValue("user_id"), Name: r.FormValue("filename")}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(file)
	})

	server.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		// Mock the file retrieval response
		files := []File{
			{ID: 1, UserID: "user1", Name: "file1.txt"},
			{ID: 2, UserID: "user2", Name: "file2.txt"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	})

	server.HandleFunc("/share", func(w http.ResponseWriter, r *http.Request) {
		// Mock the share response
		w.WriteHeader(http.StatusOK)
	})

	// Start the server in a goroutine to avoid blocking the test
	go http.ListenAndServe(":8080", server)

	// Create an HTTP client
	client := &http.Client{}

	// Step 1: Upload files for two users
	user1File := uploadFile(t, client, "user1", "file1.txt")
	user2File := uploadFile(t, client, "user2", "file2.txt")

	// Step 2: Ensure files are accessible to the respective users
	if _, err := getFile(t, client, user1File.ID, "user1"); err != nil {
		t.Fatalf("Failed to retrieve user1 file: %v", err)
	}

	if _, err := getFile(t, client, user2File.ID, "user2"); err != nil {
		t.Fatalf("Failed to retrieve user2 file: %v", err)
	}

	// Step 3: Share user1's file with user2
	shareFile(t, client, user1File.ID)

	// Step 4: Ensure user2 can now access user1's shared file
	if _, err := getFile(t, client, user1File.ID, "user2"); err != nil {
		t.Fatalf("Failed to retrieve shared user1 file as user2: %v", err)
	}

	// Step 5: Ensure user1 can still access user2's non-shared file
	if _, err := getFile(t, client, user2File.ID, "user1"); err != nil {
		t.Fatalf("Failed to retrieve user2 non-shared file as user1: %v", err)
	}
}

func uploadFile(t *testing.T, client *http.Client, userID, filename string) *File {
	resp, err := client.Post("http://localhost:8080/upload", "application/x-www-form-urlencoded", strings.NewReader(fmt.Sprintf("filename=%s&user_id=%s", filename, userID)))
	if err != nil {
		t.Fatalf("Failed to upload file: %v", err)
	}

	var file File
	if err := json.NewDecoder(resp.Body).Decode(&file); err != nil {
		t.Fatalf("Failed to decode file response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status: %d", resp.StatusCode)
	}

	return &file
}

func getFile(t *testing.T, client *http.Client, fileID int, userID string) (*File, error) {
	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/files?user_id=%s", userID))
	if err != nil {
		t.Fatalf("Failed to get files for user: %v", err)
	}

	var files []File
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		t.Fatalf("Failed to decode files response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status: %d", resp.StatusCode)
	}

	for _, file := range files {
		if file.ID == fileID {
			return &file, nil
		}
	}

	return nil, fmt.Errorf("File with ID %d not found", fileID)
}

func shareFile(t *testing.T, client *http.Client, fileID int) {
	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/share?file_id=%d", fileID))
	if err != nil {
		t.Fatalf("Failed to share file: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status: %d", resp.StatusCode)
	}
}
