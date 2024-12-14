package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Incident struct to hold incident data
type Incident struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Severity    string    `json:"severity"`
	Description string    `json:"description"`
}

// LogIncident logs an incident to file
func LogIncident(incident Incident) {
	file, err := os.OpenFile("incidents.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(incident); err != nil {
		log.Fatalf("Error encoding incident: %v", err)
	}

	fmt.Printf("Incident logged: %+v\n", incident)
}

// RetainAndDeleteIncidents rotates the log file based on retention policy
func RetainAndDeleteIncidents(retentionTime time.Duration) {
	for {
		time.Sleep(24 * time.Hour) // Check every 24 hours

		files, err := os.ReadDir(".")
		if err != nil {
			log.Printf("Error reading directory: %v", err)
			continue
		}

		for _, file := range files {
			if file.Info().IsDir() {
				continue
			}

			if file.Name() != "incidents.log" {
				continue
			}

			fileAge := time.Now().Sub(file.Info().ModTime())
			if fileAge > retentionTime {
				if err := os.Remove(file.Name()); err != nil {
					log.Printf("Error deleting file %s: %v", file.Name(), err)
				} else {
					log.Printf("Deleted file %s, age: %v", file.Name(), fileAge)
				}
			}
		}
	}
}

func main() {
	// Define a retention policy (e.g., keep logs for 7 days)
	retentionPolicy := 7 * 24 * time.Hour

	// Log an incident
	incident := Incident{
		ID:          "inc-001",
		Timestamp:   time.Now(),
		Severity:    "Critical",
		Description: "Server failure detected.",
	}
	LogIncident(incident)

	// Start a goroutine to handle log retention
	go RetainAndDeleteIncidents(retentionPolicy)

	select {} // Block the main thread to keep the application running
}
