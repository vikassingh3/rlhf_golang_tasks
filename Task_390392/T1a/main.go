package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	windowSize       = 5   // Size of the moving average window
	anomalyThreshold = 2.0 // Threshold for considering a value as an anomaly
)

type AnomalyDetector struct {
	data      []float64
	movingAvg float64
	count     int
}

func (ad *AnomalyDetector) Add(value float64) {
	ad.data = append(ad.data, value)

	if len(ad.data) > windowSize {
		ad.data = ad.data[1:] // Keep only the latest windowSize values
	}

	// Calculate the moving average
	sum := 0.0
	for _, v := range ad.data {
		sum += v
	}
	ad.movingAvg = sum / float64(len(ad.data))

	// Detect anomaly
	ad.detectAnomaly(value)
}

func (ad *AnomalyDetector) detectAnomaly(value float64) {
	if math.Abs(value-ad.movingAvg) > anomalyThreshold {
		fmt.Printf("Anomaly detected! Value: %.2f, Moving Average: %.2f\n", value, ad.movingAvg)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ad := &AnomalyDetector{}

	// Simulate stream of data
	for i := 0; i < 100; i++ {
		// Simulate normal data
		value := rand.Float64() * 10
		if i == 50 { // Inject an anomaly
			value = 50
		}
		ad.Add(value)
		time.Sleep(100 * time.Millisecond) // Simulate delay in data streaming
	}
}
