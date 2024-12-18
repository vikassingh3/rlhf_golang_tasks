package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"math/rand"
)

// ExternalService represents an external service we want to fetch data from
type ExternalService struct {
	URL  string
	Name string
	Data interface{} // Example data structure
}

// SimpleResult represents the result of a service call
type SimpleResult struct {
	Name  string
	Data  interface{}
	Error error
}

// FetchData fetches data from an external service URL
func FetchData(ctx context.Context, svc ExternalService) (*SimpleResult, error) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate latency

	req, err := http.NewRequestWithContext(ctx, "GET", svc.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform request: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %v", err)
	}

	var result SimpleResult
	err = json.Unmarshal(body, &svc.Data)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	result.Name = svc.Name
	result.Data = svc.Data
	return &result, nil
}

func main() {
	services := []ExternalService{
		{
			URL:  "https://jsonplaceholder.typicode.com/posts/1",
			Name: "Service A",
			Data: struct {
				UserID int    `json:"userId"`
				ID     int    `json:"id"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			}{},
		},
		{
			URL:  "https://jsonplaceholder.typicode.com/comments/1",
			Name: "Service B",
			Data: struct {
				PostID int    `json:"postId"`
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Email  string `json:"email"`
			}{},
		},
	}

	var results []SimpleResult
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, svc := range services {
		wg.Add(1)
		go func(svc ExternalService) {
			defer wg.Done()
			result, err := FetchData(ctx, svc)
			results = append(results, *result)
			if err != nil {
				log.Printf("Error fetching data from %s: %v", svc.Name, err)
			}
		}(svc)
	}

	wg.Wait()

	for _, result := range results {
		fmt.Printf("Fetched from %s: %+v\n", result.Name, result.Data)
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		}
	}
}
