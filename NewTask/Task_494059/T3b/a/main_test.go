package main

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "testing"
    "github.com/gorilla/mux"
    "net/http/httptest"
)

// uploadHandler simulates file upload handling
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Simple handler that just writes the data back with OK status
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("Upload successful"))
    if err != nil {
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
    }
}

// TestLargeFileUpload tests uploading a large file
func TestLargeFileUpload(t *testing.T) {
    const largeFileSize = 100 * 1024 * 1024 // 100 MB

    data := make([]byte, largeFileSize)

    // Create a new router and add the upload handler
    router := mux.NewRouter()
    router.HandleFunc("/upload", uploadHandler).Methods("POST")

    // Create an in-memory test server
    server := httptest.NewServer(router)
    defer server.Close()

    client := &http.Client{}

    // Upload the large file
    req, err := http.NewRequest("POST", server.URL+"/upload", bytes.NewReader(data))
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

    // Read the response body
    fileBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    // Validate the response data (it should match the expected response)
    expectedResponse := "Upload successful"
    if string(fileBytes) != expectedResponse {
        t.Fatalf("Unexpected response body: %s", string(fileBytes))
    }
}
