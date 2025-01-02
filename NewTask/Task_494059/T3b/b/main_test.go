package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"github.com/gorilla/mux"
	"net/http/httptest"
)

// listFilesHandler simulates returning files for a user
func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate returning file info for the user
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Send a successful response with the user ID
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Files for user %s", userID)))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

// TestConcurrentAccess tests concurrent access to files for multiple users
func TestConcurrentAccess(t *testing.T) {
	client := &http.Client{}
	const numUsers = 10
	var wg sync.WaitGroup

	// Create a new router and add the listFilesHandler
	router := mux.NewRouter()
	router.HandleFunc("/files", listFilesHandler).Methods("GET")

	// Create an in-memory test server
	server := httptest.NewServer(router)
	defer server.Close()

	wg.Add(numUsers)

	// Launch concurrent requests for different users
	for i := 0; i < numUsers; i++ {
		go func(userID int) {
			defer wg.Done()
			resp, err := client.Get(fmt.Sprintf("%s/files?user_id=user%d", server.URL, userID))
			if err != nil {
				t.Fatalf("Failed to access files for user %d: %v", userID, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Unexpected status code for user %d: %d", userID, resp.StatusCode)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
