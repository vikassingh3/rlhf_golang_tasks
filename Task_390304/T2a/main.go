package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Incident struct {
	ID         string    `json:"id"`
	Message    string    `json:"message"`
	OccurredAt time.Time `json:"occurred_at"`
}

// Global variable to store incidents
var incidentStore []Incident

// Define the retention period (e.g., 30 days)
const logRetentionPeriod = 30 * 24 * time.Hour

func main() {
	// Logging incidents
	logIncident("User login failed: invalid credentials", "001")
	logIncident("Data accessed without permission", "002")
	logIncident("System error encountered", "003")

	// Start cleanup process in a separate routine
	go startCleanupJob()

	// Keep the application running to simulate a long-running process
	select {}
}

// Function to log incidents
func logIncident(message, id string) {
	incident := Incident{
		ID:         id,
		Message:    message,
		OccurredAt: time.Now(),
	}
	incidentStore = append(incidentStore, incident)

	// Optionally, log to a file or monitoring system
	logToFile(incident)
	fmt.Println("Logged incident:", incident)
}

// Save incident log to a file
func logToFile(incident Incident) {
	file, err := os.OpenFile("incidents.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("could not open log file: %v", err)
	}
	defer file.Close()

	data, err := json.Marshal(incident)
	if err != nil {
		log.Println("Error marshaling incident for logging:", err)
		return
	}

	if _, err := file.Write(data); err != nil {
		log.Println("Error writing to log file:", err)
	}
	file.WriteString("\n") // newline for clarity
}

// CleanupJob method to periodically run retention logic
func startCleanupJob() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		<-ticker.C
		cleanupTimedOutIncidents()
	}
}

// Deletes incidents older than logRetentionPeriod
func cleanupTimedOutIncidents() {
	currentTime := time.Now()
	for i := 0; i < len(incidentStore); {
		if currentTime.Sub(incidentStore[i].OccurredAt) > logRetentionPeriod {
			incidentStore = append(incidentStore[:i], incidentStore[i+1:]...) // Remove outdated incident
		} else {
			i++ // Only increment if no deletion occurred
		}
	}
	// Output number of retained incidents
	fmt.Printf("Remaining incidents after cleanup: %d\n", len(incidentStore))
}
