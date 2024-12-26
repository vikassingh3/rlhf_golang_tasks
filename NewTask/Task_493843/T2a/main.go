package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"time"
)

type QueryParameters struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var userIDRegex = regexp.MustCompile(`^[0-9]+$`)

type AnomalyLog struct {
	Param   string `json:"param"`
	Value   string `json:"value"`
	Count   int    `json:"count"`
	Suspicious bool   `json:"suspicious"`
	FirstSeen  string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}

var anomalyLogs []AnomalyLog
var anomaliesMu sync.Mutex

func logAnomaly(param, value string, firstSeen, lastSeen time.Time, suspicious bool) {
	anomaliesMu.Lock()
	defer anomaliesMu.Unlock()
	for i, log := range anomalyLogs {
		if log.Param == param && log.Value == value {
			anomalyLogs[i].Count++
			if suspicious {
				anomalyLogs[i].Suspicious = true
			}
			anomalyLogs[i].LastSeen = lastSeen.Format(time.RFC3339)
			return
		}
	}
	anomalyLogs = append(anomalyLogs, AnomalyLog{
		Param:      param,
		Value:      value,
		Count:      1,
		Suspicious: suspicious,
		FirstSeen:  firstSeen.Format(time.RFC3339),
		LastSeen:   lastSeen.Format(time.RFC3339),
	})
}

func filterAnomalies(anomalies *[]AnomalyLog) {
	// Example logic to filter anomalies: considera low threshold for setting suspicious
	for i := len(*anomalies) - 1; i >= 0; i-- {
		if (*anomalies)[i].Count < 3 { // Adjust threshold as needed
			*anomalies = append((*anomalies)[:i], (*anomalies)[i+1:]...)
		}
	}
}

func parameterHandler(w http.ResponseWriter, r *http.Request) {
	var anomalySuspicious = false

	queryParams := r.URL.Query()
	firstSeen := time.Now()
	lastSeen := firstSeen

	for key, value := range queryParams {
		switch key {
		case "userID":
			if !userIDRegex.MatchString(value[0]) {
				logAnomaly(key, value[0], firstSeen, lastSeen, anomalySuspicious)
			}
		case "email":
			if !emailRegex.MatchString(value[0]) {
				logAnomaly(key, value[0], firstSeen, lastSeen, anomalySuspicious)
			}
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Anomaly Handling Complete")
}

func printAnomalies() {
	anomaliesMu.Lock()
	defer anomaliesMu.Unlock()
	filterAnomalies(&anomalyLogs)
	fmt.Println("Anomalies:")
	for _, log := range anomalyLogs {
		data, err := json.MarshalIndent(log, "", "  ")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(data))
	}
}

func main() {
	http.HandleFunc("/api/parameters", parameterHandler)
	go func() {
		ticker := time.NewTicker(time.Minute)
		for range ticker.C {
			printAnomalies()
		}
	}()

	fmt.Println("Server running on http://localhost:8080/api/parameters")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}