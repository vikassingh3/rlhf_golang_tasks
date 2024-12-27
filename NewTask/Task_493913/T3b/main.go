package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Defining some predefined Feedback Types
const (
	FeatureRequest FeedbackType = "feature_request"
	BugReport      FeedbackType = "bug_report"
	Comment        FeedbackType = "comment"
	UserSurvey     FeedbackType = "user_survey" // New feedback type added
)

// FeedbackType represents the type of feedback
type FeedbackType string

// Feedback is an interface that encapsulates different feedback types
type Feedback interface {
	GetFeedbackType() FeedbackType
	GetDetails() interface{}
}

// A generic struct to handle any feedback
type feedback struct {
	Type    FeedbackType `json:"type"`
	Details interface{}  `json:"details"`
}

// Implement the Feedback interface
func (f *feedback) GetFeedbackType() FeedbackType {
	return f.Type
}

func (f *feedback) GetDetails() interface{} {
	return f.Details
}

// Specific feedback types
type featureRequest struct {
	Feature string `json:"feature"`
	Benefit string `json:"benefit"`
}

type bugReport struct {
	Description       string `json:"description"`
	StepsToReproduce  string `json:"steps_to_reproduce"`
}

type userSurvey struct {
	Question1 string `json:"question1"`
	Question2 string `json:"question2"`
	Question3 string `json:"question3"`
}

// RegisterFeedbackTypes function registers new feedback types dynamically
var feedbackTypeRegistry = map[FeedbackType]reflect.Type{}

func init() {
	RegisterFeedbackTypes()
}

func RegisterFeedbackTypes() {
	feedbackTypeRegistry[FeatureRequest] = reflect.TypeOf(featureRequest{})
	feedbackTypeRegistry[BugReport] = reflect.TypeOf(bugReport{})
	feedbackTypeRegistry[UserSurvey] = reflect.TypeOf(userSurvey{})
}

// CreateFeedback function creates feedback instances dynamically
func CreateFeedback(feedbackData map[string]interface{}) (Feedback, error) {
	feedbackTypeStr, ok := feedbackData["type"].(string)
	if !ok {
		return nil, errors.New("missing or invalid 'type' field")
	}

	feedbackType, exists := feedbackTypeRegistry[FeedbackType(feedbackTypeStr)]
	if !exists {
		return nil, fmt.Errorf("unknown feedback type: %s", feedbackTypeStr)
	}

	// Create a new instance of the feedback type
	feedbackInstance := reflect.New(feedbackType).Interface()

	// Decode JSON-like map into the instance
	detailsJSON, err := json.Marshal(feedbackData["details"])
	if err != nil {
		return nil, fmt.Errorf("failed to encode details: %v", err)
	}

	if err := json.Unmarshal(detailsJSON, feedbackInstance); err != nil {
		return nil, fmt.Errorf("failed to decode details: %v", err)
	}

	return &feedback{
		Type:    FeedbackType(feedbackTypeStr),
		Details: feedbackInstance,
	}, nil
}

// Display feedback details
func displayFeedback(f Feedback) {
	fmt.Printf("Feedback Type: %v\n", f.GetFeedbackType())
	fmt.Println("Details:")
	details, _ := json.MarshalIndent(f.GetDetails(), "  ", "  ")
	fmt.Println(string(details))
	fmt.Println("--------")
}

// Main function
func main() {
	// Example feedback data
	feedbacks := []map[string]interface{}{
		{
			"type": FeatureRequest,
			"details": map[string]interface{}{
				"feature": "New Search Feature",
				"benefit": "Faster and more efficient search",
			},
		},
		{
			"type": BugReport,
			"details": map[string]interface{}{
				"description":       "Application crashes when clicking button X",
				"steps_to_reproduce": "Open the app, click button X",
			},
		},
		{
			"type": UserSurvey,
			"details": map[string]interface{}{
				"question1": "Answer 1",
				"question2": "Answer 2",
				"question3": "Answer 3",
			},
		},
	}

	// Process and display each feedback
	for i, data := range feedbacks {
		fb, err := CreateFeedback(data)
		if err != nil {
			fmt.Printf("Error creating feedback for item %d: %v\n", i+1, err)
			continue
		}
		displayFeedback(fb)
	}
}
