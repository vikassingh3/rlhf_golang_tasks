package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

// Assuming File struct
type File struct {
	ID     int
	Name   string
	UserID string
}

// Assuming we have a function to upload files (mock implementation)
func uploadFile(t *testing.T, client *http.Client, userID, filename string, data []byte) (*http.Response, error) {
	// Simulating file upload logic
	fmt.Println("Uploading file:", filename)
	// You can simulate a real HTTP request here to your upload API.
	// For now, we'll return a mock response
	return &http.Response{StatusCode: http.StatusOK}, nil
}

// Mock function to simulate retrieving files for a user
func getFilesForUser(t *testing.T, client *http.Client, userID string) ([]File, error) {
	// Simulate file retrieval
	files := []File{
		{ID: 1, Name: "file1.txt", UserID: userID},
		{ID: 2, Name: "large_file.txt", UserID: userID}, // This is the file we uploaded
	}
	return files, nil
}

// Mock function to simulate retrieving a specific file
func getFile(t *testing.T, client *http.Client, fileID int, userID string) (*File, error) {
	// Simulate getting a file by ID
	return &File{ID: fileID, Name: "large_file.txt", UserID: userID}, nil
}

func TestLargeFileUpload(t *testing.T) {
	// Generate a large file data (e.g., using bytes.Repeat)
	largeFileData := make([]byte, 10*1024*1024) // 10 MB file
	
	client := &http.Client{}
	userID := "largeFileUser"
	filename := "large_file.txt"
	
	_, err := uploadFile(t, client, userID, filename, largeFileData)
	if err != nil {
		t.Fatalf("Failed to upload large file: %v", err)
	}
	
	// Ensure the file is in the list of files for the user
	files, err := getFilesForUser(t, client, userID)
	if err != nil {
		t.Fatalf("Failed to retrieve files for user after large file upload: %v", err)
	}
	
	fileFound := false
	for _, file := range files {
		if file.Name == filename {
			fileFound = true
			break
		}
	}
	if !fileFound {
		t.Fatalf("Uploaded large file not found in the user's file list.")
	}
}

func TestMultipleUserAccess(t *testing.T) {
	const numUsers = 10
	const numAccesses = 100
	client := &http.Client{}
	userID := "largeFileUser"
	filename := "large_file.txt"
	// Mock the file upload
	_, err := uploadFile(t, client, userID, filename, []byte("dummy data"))
	if err != nil {
		t.Fatalf("Failed to upload file: %v", err)
	}

	// Simulating file access
	fileID := 1 // Example file ID (You can adjust as per actual file ID logic)

	var wg sync.WaitGroup
	wg.Add(numUsers)

	errorsOccurred := false

	for i := 0; i < numUsers; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numAccesses; j++ {
				// Simulate accessing the file (e.g., reading or writing)
				_, err := getFile(t, client, fileID, userID)
				if err != nil {
					errorsOccurred = true
					t.Logf("Error accessing file: %v", err)
				}
			}
		}(i)
	}

	wg.Wait()

	// Check if any errors occurred during the accesses
	if errorsOccurred {
		t.Fatalf("Multiple user accesses to the same file caused errors.")
	}
}
