package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	// Open file with append and create flags
	file, err := os.OpenFile("incidents.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
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
func RetainAndDeleteIncidents(retentionTime time.Duration, stopChan <-chan struct{}) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Check log files for retention
			fileInfo, err := os.Stat("incidents.log")
			if err != nil {
				if os.IsNotExist(err) {
					continue // Log file doesn't exist, skip
				}
				log.Printf("Error stating file: %v", err)
				continue
			}

			fileAge := time.Since(fileInfo.ModTime())
			if fileAge > retentionTime {
				if err := os.Remove("incidents.log"); err != nil {
					log.Printf("Error deleting file incidents.log: %v", err)
				} else {
					log.Printf("Deleted log file incidents.log, age: %v", fileAge)
				}
			}

		case <-stopChan:
			// Graceful shutdown
			log.Println("Stopping log retention handler...")
			return
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

	// Set up channel to handle shutdown
	stopChan := make(chan struct{})
	go RetainAndDeleteIncidents(retentionPolicy, stopChan)

	// Handle OS signals for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan // Wait for termination signal
	close(stopChan)
	log.Println("Application shutting down gracefully.")
}
