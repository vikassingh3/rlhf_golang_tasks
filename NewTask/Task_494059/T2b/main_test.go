package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// File struct to represent a file
type File struct {
	ID     int
	Name   string
	UserID int
	Data   []byte
	Shared bool
}

// Declare a global variable to hold the files during tests
var files []File

// Setup and teardown functions for integration tests
func setup() {
	files = nil
}

func teardown() {
	files = nil
}

// TestFileSharingIntegration tests file sharing functionality
func TestFileSharingIntegration(t *testing.T) {
	// Set up the test environment
	setup()
	defer teardown()

	// User 1 uploads a file
	user1ID := 1
	fileName := "test_file.txt"
	uploadFile(fileName, user1ID)

	// User 2 lists the files - Should see only User 1's file (empty list expected for User 2)
	user2Files := listFiles(2)
	assert.Len(t, user2Files, 0)

	// User 1 shares the file
	shareFile(fileName)

	// User 2 lists the files - Should see the shared file
	user2Files = listFiles(2)
	assert.Len(t, user2Files, 1)
	assert.Equal(t, user2Files[0].Name, fileName)

	// User 1 unshares the file
	unshareFile(fileName)

	// User 2 lists the files - Should not see the shared file anymore
	user2Files = listFiles(2)
	assert.Len(t, user2Files, 0)
}

// Helper functions for uploading, listing, and sharing files (similar to your actual handlers)
func uploadFile(fileName string, userID int) {
	// Implement your file upload logic using the 'files' slice
	newFile := File{
		ID:     len(files) + 1,
		Name:   fileName,
		UserID: userID,
		Data:   []byte("Test file data"), // Placeholder data
		Shared: false,
	}
	files = append(files, newFile)
}

func listFiles(userID int) []File {
	// Implement your file listing logic using the 'files' slice
	var userFiles []File
	for _, file := range files {
		if file.UserID == userID || file.Shared { // Allow user to see their own files or shared files
			userFiles = append(userFiles, file)
		}
	}
	return userFiles
}

func shareFile(fileName string) {
	// Implement your file sharing logic using the 'files' slice
	for i, file := range files {
		if file.Name == fileName {
			files[i].Shared = true
			break
		}
	}
}

func unshareFile(fileName string) {
	// Implement your file unsharing logic using the 'files' slice
	for i, file := range files {
		if file.Name == fileName {
			files[i].Shared = false
			break
		}
	}
}
