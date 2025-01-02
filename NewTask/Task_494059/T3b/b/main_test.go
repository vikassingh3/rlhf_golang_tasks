package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestConcurrentAccess(t *testing.T) {
    client := &http.Client{}
    const numUsers = 10
    var wg sync.WaitGroup

    wg.Add(numUsers)

    for i := 0; i < numUsers; i++ {
        go func(userID int) {
            defer wg.Done()
            _, err := client.Get(fmt.Sprintf("http://localhost:8080/files?user_id=user%d", userID))
            if err != nil {
                t.Fatalf("Failed to access files for user: %v", err)
            }
        }(i)
    }

    wg.Wait()
}