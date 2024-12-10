package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MyAPI represents the client for the MyAPI service.
type MyAPI struct {
	baseURL    string
	httpClient *http.Client
}

// NewMyAPI returns a new MyAPI client with the specified base URL.
func NewMyAPI(baseURL string) *MyAPI {
	return &MyAPI{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// GetData fetches data from the API.
func (c *MyAPI) GetData() (*DataResponse, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/data", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var dataResponse DataResponse
	err = json.NewDecoder(resp.Body).Decode(&dataResponse)
	if err != nil {
		return nil, err
	}

	return &dataResponse, nil
}

// DataResponse represents the response from the API.
type DataResponse struct {
	Message string `json:"message"`
	Data    []int  `json:"data"`
}
