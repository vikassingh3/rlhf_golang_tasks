package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Define FeedbackType as a string type
type FeedbackType string

const (
	FeatureRequest FeedbackType = "FeatureRequest"
	BugReport      FeedbackType = "BugReport"
	Comment        FeedbackType = "Comment"
)

// Define generic feedback struct
type feedback struct {
	Type    FeedbackType         `json:"type"`
	Details map[string]interface{} `json:"details"`
}

// Define specific feedback structs for each type
type featureRequestStruct struct {
	Type    FeedbackType `json:"type"`
	UserID  int          `json:"user_id"`
	Feature string       `json:"feature"`
	Benefit string       `json:"benefit"`
}

type bugReportStruct struct {
	Type             FeedbackType `json:"type"`
	UserID           int          `json:"user_id"`
	AppVersion       string       `json:"app_version"`
	Description      string       `json:"description"`
	StepsToReproduce string       `json:"steps_to_reproduce"`
}

type commentStruct struct {
	Type   FeedbackType `json:"type"`
	UserID int          `json:"user_id"`
	Text   string       `json:"text"`
}

// Function to validate the feedback data
func validateFeedbackData(feedbackData map[string]interface{}) error {
	for id, data := range feedbackData {
		// Convert the feedback data to JSON bytes
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("error marshaling feedback data for %s: %w", id, err)
		}

		// Determine the feedback type
		fbType, ok := feedbackData[id].(map[string]interface{})["type"].(string)
		if !ok {
			return fmt.Errorf("feedback data for %s is missing or has invalid type", id)
		}

		// Create a new instance of the corresponding feedback struct to unmarshal into
		var feedbackStruct interface{}
		switch FeedbackType(fbType) {
		case FeatureRequest:
			feedbackStruct = &featureRequestStruct{}
		case BugReport:
			feedbackStruct = &bugReportStruct{}
		case Comment:
			feedbackStruct = &commentStruct{}
		default:
			return fmt.Errorf("unknown feedback type: %s", fbType)
		}

		// Unmarshal the data into the struct
		err = json.Unmarshal(jsonBytes, feedbackStruct)
		if err != nil {
			return fmt.Errorf("error unmarshaling feedback data for %s: %w", id, err)
		}

		// Access the validated data using reflection
		v := reflect.ValueOf(feedbackStruct).Elem()
		fmt.Printf("Validated Feedback %s:\n", id)
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			value := v.Field(i)
			fmt.Printf("  %s: %v\n", field.Name, value.Interface())
		}
		fmt.Println("--------")
	}
	return nil
}

func main() {
	// Example feedback data
	feedbacks := map[string]interface{}{
		"feedback_1": map[string]interface{}{
			"type":    "FeatureRequest",
			"user_id": 1,
			"feature": "Dark Mode",
			"benefit": "Improved user experience",
		},
		"feedback_2": map[string]interface{}{
			"type":             "BugReport",
			"user_id":          2,
			"app_version":      "1.2.3",
			"description":      "App crashes on login",
			"steps_to_reproduce": "Open app, click login",
		},
		"feedback_3": map[string]interface{}{
			"type":    "Comment",
			"user_id": 3,
			"text":    "Great app! Keep up the good work.",
		},
	}

	// Validate the feedback data
	err := validateFeedbackData(feedbacks)
	if err != nil {
		fmt.Println("Error validating feedback:", err)
	}
}
