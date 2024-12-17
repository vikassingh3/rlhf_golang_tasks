package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Data structure for streaming data
type Data struct {
	Feature1 float64 `json:"feature1"`
	Feature2 float64 `json:"feature2"`
}

// Payload structure for ML model API request
type MLRequest struct {
	Data [][]float64 `json:"data"`
}

// Response structure for ML model API response
type MLResponse struct {
	Predictions []int `json:"predictions"`
}

// Simulated data generation function
func generateData() []float64 {
	if rand.Intn(100) > 95 { // 5% chance of generating anomaly
		return []float64{rand.Float64()*100 + 200, rand.Float64()*100 + 200} // Simulated anomaly
	}
	return []float64{rand.Float64() * 100, rand.Float64() * 100} // Normal data
}

// Function to call the ML model API for anomaly detection
func callModelAPI(data [][]float64) ([]int, error) {
	// Prepare the request payload
	payload := MLRequest{Data: data}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON payload: %v", err)
	}

	// Send POST request to the model API
	resp, err := http.Post("http://localhost:5000/predict", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error sending request to model API: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var mlResponse MLResponse
	if err := json.NewDecoder(resp.Body).Decode(&mlResponse); err != nil {
		return nil, fmt.Errorf("error decoding API response: %v", err)
	}

	return mlResponse.Predictions, nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	dataChannel := make(chan []float64)

	// Start a goroutine to generate streaming data
	go func() {
		for {
			data := generateData()
			dataChannel <- data
			time.Sleep(100 * time.Millisecond) // Simulate real-time streaming
		}
	}()

	// Process data from the channel and detect anomalies
	for data := range dataChannel {
		predictions, err := callModelAPI([][]float64{data})
		if err != nil {
			fmt.Printf("Error during anomaly detection: %v\n", err)
			continue
		}

		// Check prediction and log the result
		if predictions[0] == -1 {
			fmt.Printf("Anomaly detected: %v\n", data)
		} else {
			fmt.Printf("Normal data: %v\n", data)
		}
	}
}
