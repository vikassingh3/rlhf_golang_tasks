package main

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "testing"
)

func TestLargeFileUpload(t *testing.T) {
    const largeFileSize = 100 * 1024 * 1024 // 100 MB

    data := make([]byte, largeFileSize)
    client := &http.Client{}

    // Upload the large file
    req, err := http.NewRequest("POST", "http://localhost:8080/upload", bytes.NewReader(data))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    resp, err := client.Do(req)
    if err != nil {
        t.Fatalf("Failed to upload file: %v", err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Invalid response status: %d", resp.StatusCode)
    }

    // Download and validate the uploaded file
    fileBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    if !bytes.Equal(fileBytes, data) {
        t.Fatalf("File data mismatch.")
    }
}