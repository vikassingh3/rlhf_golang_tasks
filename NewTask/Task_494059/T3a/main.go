package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Define the File struct
type File struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	Shared bool   `json:"shared"`
	Data   []byte `json:"data"`
}

func TestLargeFileUpload(t *testing.T) {
	client := &http.Client{}

	// Generate large file data
	largeData := make([]byte, 10*1024*1024) // 10 MB
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// Prepare form data
	formData := url.Values{}
	formData.Set("filename", "largefile.txt")

	// Upload large file
	resp, err := client.Post("http://localhost:8080/upload", "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatalf("Failed to upload large file: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status: %d", resp.StatusCode)
	}

	// Verify large file upload
	if _, err := getFile(t, client, resp.Body); err != nil {
		t.Fatalf("Failed to retrieve large file: %v", err)
	}
}

func TestMultipleConcurrentUploads(t *testing.T) {
	client := &http.Client{}

	// Upload multiple files concurrently
	numUsers := 10
	for i := 0; i < numUsers; i++ {
		go func(i int) {
			userID := strconv.Itoa(i + 1)
			filename := fmt.Sprintf("file%d.txt", i+1)

			// Prepare form data
			formData := url.Values{}
			formData.Set("filename", filename)
			formData.Set("user_id", userID)

			resp, err := client.Post("http://localhost:8080/upload", "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
			if err != nil {
				t.Fatalf("Failed to upload file %s as user %s: %v", filename, userID, err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Invalid response status: %d", resp.StatusCode)
			}
		}(i)
	}

	// Wait for all uploads to complete
	time.Sleep(5 * time.Second)
}

func getFile(t *testing.T, client *http.Client, body io.Reader) (*File, error) {
	resp, err := client.Get("http://localhost:8080/files?user_id=1")
	if err != nil {
		t.Fatalf("Failed to get files: %v", err)
	}

	defer resp.Body.Close()
	fileData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var file File
	if err := json.Unmarshal(fileData, &file); err != nil {
		t.Fatalf("Failed to unmarshal file data: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status: %d", resp.StatusCode)
	}

	return &file, nil
}
