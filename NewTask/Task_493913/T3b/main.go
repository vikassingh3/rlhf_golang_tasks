package main

import (
	"fmt"
	"reflect"
)

// Feedback interface defines the contract for all feedback types
type Feedback interface {
	GetFeedbackType() string
	Display() string
}

// Registry to dynamically register feedback types
var feedbackRegistry = make(map[string]reflect.Type)

// RegisterFeedbackType registers a new feedback type dynamically
func RegisterFeedbackType(typeName string, fbType reflect.Type) {
	feedbackRegistry[typeName] = fbType
}

// CreateFeedback dynamically creates a feedback instance
func CreateFeedback(typeName string, params map[string]interface{}) (Feedback, error) {
	fbType, exists := feedbackRegistry[typeName]
	if !exists {
		return nil, fmt.Errorf("feedback type '%s' not registered", typeName)
	}
	instance := reflect.New(fbType).Interface().(Feedback)
	val := reflect.ValueOf(instance).Elem()
	for key, value := range params {
		field := val.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}
	return instance, nil
}

// Example Feedback Types
type BugReport struct {
	Title       string
	Description string
}

func (b BugReport) GetFeedbackType() string {
	return "BugReport"
}

func (b BugReport) Display() string {
	return fmt.Sprintf("Bug Report: %s - %s", b.Title, b.Description)
}

type FeatureRequest struct {
	Title       string
	Description string
}

func (f FeatureRequest) GetFeedbackType() string {
	return "FeatureRequest"
}

func (f FeatureRequest) Display() string {
	return fmt.Sprintf("Feature Request: %s - %s", f.Title, f.Description)
}

func main() {
	// Register feedback types
	RegisterFeedbackType("BugReport", reflect.TypeOf(BugReport{}))
	RegisterFeedbackType("FeatureRequest", reflect.TypeOf(FeatureRequest{}))

	// Dynamically create feedback instances
	params := map[string]interface{}{
		"Title":       "Search Feature",
		"Description": "Add advanced search capabilities.",
	}
	feedback, err := CreateFeedback("FeatureRequest", params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(feedback.Display())

	params = map[string]interface{}{
		"Title":       "App Crash",
		"Description": "Crash on clicking settings.",
	}
	feedback, err = CreateFeedback("BugReport", params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(feedback.Display())
}
