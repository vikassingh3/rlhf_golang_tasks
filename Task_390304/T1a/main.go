// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Define the struct for incident logging
type Incident struct {
	ID         string    `json:"id"`
	Message    string    `json:"message"`
	OccurredAt time.Time `json:"occurred_at"`
}

// Example storage mechanism with a retention policy
var incidentStore []Incident

func main() {
	logIncident("Unauthorized access attempt.", "12345")
	logIncident("User data breach detected.", "12346")

	// Simulate a scheduled cleanup every hour
	go startCleanupJob()

	// Run application...
	select {} // Keep app running
}

// logIncident logs an incident while adhering to retention policies
func logIncident(message, id string) {
	incident := Incident{
		ID:         id,
		Message:    message,
		OccurredAt: time.Now(),
	}
	incidentStore = append(incidentStore, incident)
	fmt.Println("Logged Incident:", incident)
}

// startCleanupJob mimics a scheduled cleanup task
func startCleanupJob() {
	ticker := time.NewTicker(1 * time.Hour) // Cleanup interval
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cleanupIncidents()
		}
	}
}

// cleanupIncidents removes old incidents based on defined retention policies
func cleanupIncidents() {
	maxRetentionPeriod := 7 * 24 * time.Hour // Retain incidents for 7 days
	now := time.Now()

	var retainedIncidents []Incident
	for _, incident := range incidentStore {
		if now.Sub(incident.OccurredAt) <= maxRetentionPeriod {
			retainedIncidents = append(retainedIncidents, incident)
		}
	}

	incidentStore = retainedIncidents
	log.Println("Cleaned up incidents. Retained:", len(incidentStore))
}

// Utility function to log incidents (could be used for JSON logging)
func logAsJSON(incident Incident) {
	data, _ := json.Marshal(incident)
	log.Println(string(data))
}
