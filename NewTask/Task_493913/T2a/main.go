package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// FeedbackType represents the type of feedback
type FeedbackType string

const (
	FeatureRequest FeedbackType = "feature_request"
	BugReport      FeedbackType = "bug_report"
	Comment        FeedbackType = "comment"
)

// Feedback is an interface that encapsulates different feedback types
type Feedback interface {
	GetFeedbackType() FeedbackType
	GetDetails() map[string]interface{}
}

// A generic struct to handle any feedback
type feedback struct {
	Type    FeedbackType         `json:"type"`
	Details map[string]interface{} `json:"details"`
}

// Implement the Feedback interface
func (f *feedback) GetFeedbackType() FeedbackType {
	return f.Type
}

func (f *feedback) GetDetails() map[string]interface{} {
	return f.Details
}

// Implementation for different feedback types
type featureRequest struct {
	feedback  // Embedded feedback
	Feature   string `json:"feature"`
	Benefit   string `json:"benefit"`
}

type bugReport struct {
	feedback // Embedded feedback
	Description      string `json:"description"`
	StepsToReproduce string `json:"steps_to_reproduce"`
}

type comment struct {
	feedback  // Embedded feedback
	Text      string `json:"text"`
}

// ValidateFeedback uses reflection to validate the feedback data
func validateFeedback(feedback Feedback) error {
	// Get the feedback type
	feedbackType := feedback.GetFeedbackType()
	fmt.Printf("Validating Feedback Type: %v\n", feedbackType)

	// Map of expected fields for each feedback type
	expectedFields := map[FeedbackType]map[string]bool{
		FeatureRequest: {
			"user_id": true,
			"feature": true,
			"benefit": true,
		},
		BugReport: {
			"user_id":           true,
			"app_version":       true,
			"description":       true,
			"steps_to_reproduce": true,
		},
		Comment: {
			"user_id": true,
			"text":    true,
		},
	}

	// Define the expected field types for each feedback type
	expectedFieldTypes := map[FeedbackType]map[string]reflect.Kind{
		FeatureRequest: {
			"user_id": reflect.Int,
			"feature": reflect.String,
			"benefit": reflect.String,
		},
		BugReport: {
			"user_id":           reflect.Int,
			"app_version":       reflect.String,
			"description":       reflect.String,
			"steps_to_reproduce": reflect.String,
		},
		Comment: {
			"user_id": reflect.Int,
			"text":    reflect.String,
		},
	}

	// Iterate over the expected fields for the given feedback type
	for fieldName, required := range expectedFields[feedbackType] {
		// Check if the field exists in the details map
		if value, exists := feedback.GetDetails()[fieldName]; exists {
			// Check if the value type matches the expected type
			actualKind := reflect.TypeOf(value).Kind()
			expectedKind := expectedFieldTypes[feedbackType][fieldName]
			if actualKind != expectedKind {
				// Type mismatch
				return fmt.Errorf("Field '%s' has type %v, expected %v", fieldName, actualKind, expectedKind)
			}

			// Check if required fields are present and valid
			if required {
				switch actualKind {
				case reflect.Int:
					_, err := strconv.Atoi(fmt.Sprintf("%v", value))
					if err != nil {
						return fmt.Errorf("Field '%s' must be a valid integer", fieldName)
					}
				default:
					// Non-integer fields don't need further validation for this example
				}
			}
		} else {
			// Required field missing
			return fmt.Errorf("Missing required field '%s'", fieldName)
		}
	}

	return nil
}

func main() {
	// Example feedback data in a map struct
	feedbacks := map[string]Feedback{
		"feedback_1": &featureRequest{
			feedback: feedback{
				Type:    FeatureRequest,
				Details: map[string]interface{}{"user_id": 1, "feature": "New Search Feature", "benefit": "Faster and more efficient search"},
			},
			Feature: "New Search Feature",
			Benefit: "Faster and more efficient search",
		},
		"feedback_2": &bugReport{
			feedback: feedback{
				Type:    BugReport,
				Details: map[string]interface{}{"user_id": 2, "app_version": "1.0", "description": "Application crashes when clicking button X", "steps_to_reproduce": "Open the app, click button X"},
			},
			Description:      "Application crashes when clicking button X",
			StepsToReproduce: "Open the app, click button X",
		},
	}

	// Process and validate each feedback dynamically using reflection
	for id, feedback := range feedbacks {
		if err := validateFeedback(feedback); err != nil {
			fmt.Printf("Validation failed for %v: %v\n", id, err)
		} else {
			fmt.Printf("Feedback %v is valid\n", id)
		}
	}
}
