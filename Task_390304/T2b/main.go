package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type Incident struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Tags      []string  `json:"tags"`
}

var incidentLogs []Incident
var logMutex sync.Mutex

const (
	retentionPeriod = 7 * 24 * time.Hour // 7 days
)

func logIncident(level string, message string, tags ...string) {
	logMutex.Lock()
	defer logMutex.Unlock()

	incident := Incident{
		ID:        generateIncidentID(),
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Tags:      tags,
	}

	incidentLogs = append(incidentLogs, incident)
	logAsJSON(incident)
}

func generateIncidentID() string {
	return fmt.Sprintf("incident-%v", time.Now().UnixNano())
}

func logAsJSON(incident Incident) {
	data, _ := json.MarshalIndent(incident, "", "  ")
	log.Println(string(data))
}

func enforceRetentionPolicy() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logMutex.Lock()
			defer logMutex.Unlock()

			now := time.Now()
			retainedLogs := []Incident{}
			for _, log := range incidentLogs {
				if now.Sub(log.Timestamp) <= retentionPeriod {
					retainedLogs = append(retainedLogs, log)
				}
			}

			incidentLogs = retainedLogs
			fmt.Println("Retention policy enforced. Retained logs:", len(incidentLogs))
		}
	}
}

func main() {
	logIncident("ERROR", "Server crash detected.", "critical", "infrastructure")
	logIncident("WARN", "Unusual login attempt.", "security", "auth")

	go enforceRetentionPolicy()

	// Simulate application running
	select {}
}
