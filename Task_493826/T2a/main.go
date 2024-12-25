// package main

// import (
// 	"fmt"
// 	"log"
// 	"log/syslog"
// 	"os"
// 	"time"

// 	"github.com/sirupsen/logrus"
// )

// func main() {
// 	// Configure logrus
// 	logger := logrus.New()
// 	logger.SetLevel(logrus.InfoLevel)
// 	logger.SetFormatter(&logrus.JSONFormatter{})
// 	mw := logrus.NewMultiWriter(
// 		os.Stdout,
// 		syslog.New(syslog.LOG_ALERT|syslog.LOG_LOCAL0, "slice-monitor"),
// 	)
// 	logger.SetOutput(mw)

// 	// Function to handle slice errors
// 	handleSliceError := func(err error, slice interface{}, action string, index int) {
// 		if err != nil {
// 			logger.WithFields(logrus.Fields{
// 				"error":     err,
// 				"slice":     slice,
// 				"action":    action,
// 				"index":     index,
// 				"timestamp": time.Now(),
// 			}).Error("Slice error detected")
// 		}
// 	}

// 	var numbers []int = []int{1, 2, 3}

// 	// Simulate a slice out-of-bounds access
// 	handleSliceError(nil, numbers, "read", 3)

// 	// Simulate a nil slice access
// 	var nilSlice []int
// 	handleSliceError(nil, nilSlice, "read", 0)
// }

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// SlackAlert is a JSON payload for Slack alerts.
type SlackAlert struct {
	Text     string            `json:"text"`
	Username string            `json:"username"`
	IconEmoji string           `json:"icon_emoji"`
	Channel  string            `json:"channel"`
	Ts       time.Time         `json:"ts"`
	Attachments []SlackAttachment `json:"attachments"`
}

// SlackAttachment defines the structure for Slack attachments.
type SlackAttachment struct {
	Title   string      `json:"title"`
	TitleID string      `json:"title_id"`
	Fields  []SlackField `json:"fields"`
}

// SlackField defines the structure for Slack fields.
type SlackField struct {
	Title  string `json:"title"`
	Value  string `json:"value"`
	Short  bool   `json:"short"`
}

func sendSlackAlert(url string, alert *SlackAlert) error {
	data, err := json.Marshal(alert)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	// Slack Webhook URL
	slackURL := "https://hooks.slack.com/services/T12345678/B65432109/XXXXXXXXXXXXXXXXXXXXXXXX"

	// Function to handle slice errors and send alerts
	handleSliceError := func(err error, slice interface{}, action string, index int) {
		if err != nil {
			logger := logrus.New()
			logger.WithFields(logrus.Fields{
				"error":     err,
				"slice":     slice,
				"action":    action,
				"index":     index,
				"timestamp": time.Now(),
			}).Error("Slice error detected")

			// Create a Slack alert
			alert := &SlackAlert{
				Text:     "Slice error detected!",
				Username: "Slice Monitor Bot",
				IconEmoji: ":warning:",
				Channel:  "#alerts",
				Ts:       time.Now(),
				Attachments: []SlackAttachment{
					{
						Title:   "Error Details",
						TitleID: "error_details",
						Fields: []SlackField{
							{Title: "Error Message", Value: err.Error(), Short: false},
							{Title: "Slice", Value: fmt.Sprintf("%v", slice), Short: false},
							{Title: "Action", Value: action, Short: true},
							{Title: "Index", Value: fmt.Sprintf("%d", index), Short: true},
						},
					},
				},
			}

			// Send the alert to Slack
			if err := sendSlackAlert(slackURL, alert); err != nil {
				logrus.Error("Failed to send alert to Slack:", err)
			}
		}
	}

	var numbers []int = []int{1, 2, 3}

	// Simulate a slice out-of-bounds access
	handleSliceError(nil, numbers, "read", 3)

	// Simulate a nil slice access
	var nilSlice []int
	handleSliceError(nil, nilSlice, "read", 0)
}