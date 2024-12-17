package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// FetchData fetches data from an API endpoint
func FetchData(ctx context.Context, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	return ioutil.ReadAll(res.Body)
}

func TestFetchData_NetworkError(t *testing.T) {
	// Mock an error-inducing context
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the context to simulate a timeout

	_, err := FetchData(ctx, "https://example.com/data")

	require.Error(t, err)
	require.Contains(t, err.Error(), "context canceled")
}

func TestFetchData_BadURL(t *testing.T) {
	_, err := FetchData(context.Background(), "invalid-url")

	require.Error(t, err)
	require.Contains(t, err.Error(), "parse")
}

func TestFetchData_HTTPError(t *testing.T) {
	_, err := FetchData(context.Background(), "https://example.com/error")

	require.Error(t, err)
	require.Contains(t, err.Error(), "bad status: 500")
}

func main() {
	// Define the URL to fetch
	url := "https://jsonplaceholder.typicode.com/todos/1" // Example API endpoint

	// Create a context with a timeout
	ctx := context.Background()

	// Fetch data from the API
	data, err := FetchData(ctx, url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	// Print the fetched data
	fmt.Println("Fetched Data:", string(data))
}
