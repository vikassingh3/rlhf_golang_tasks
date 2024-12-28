package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// File struct represents a file entry with metadata
type File struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	UserID  int    `json:"user_id"`
	Shared  bool   `json:"shared"`
}

// TestFileSharing tests the file upload and sharing functionality
func TestFileSharing(t *testing.T) {
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

// uploadFile uploads a file for a user and returns the uploaded file
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

// getFile retrieves a file for a user by fileID and userID
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

// shareFile shares a file by its fileID
func shareFile(t *testing.T, client *http.Client, fileID int) {
	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/share?file_id=%d", fileID))
	if err != nil {
		t.Fatalf("Failed to share file: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status: %d", resp.StatusCode)
	}
}
