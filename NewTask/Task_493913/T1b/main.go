package main  
import (  
    "fmt"
    "reflect"
)

// Defining some predefined Feedback Types
const (
    FeatureRequest FeedbackType = "feature_request"
    BugReport      FeedbackType = "bug_report"
    Comment        FeedbackType = "comment"
)

// FeedbackType represents the type of feedback
type FeedbackType string

// Feedback is an interface that encapsulates different feedback types
type Feedback interface {
    GetFeedbackType() FeedbackType
}

// A generic struct to handle any feedback
type feedback struct {  
    Type    FeedbackType `json:"type"`
    Details map[string]interface{} `json:"details"`
}

// Implement the Feedback interface
func (f *feedback) GetFeedbackType() FeedbackType {
    return f.Type
}

// implementation for different feedback types
type featureRequest struct {
    feedback  // Embedded feedback
    Feature   string `json:"feature"`
    Benefit   string `json:"benefit"`
}

type bugReport struct {  
    feedback // Embedded feedback
    Description string `json:"description"`
    StepsToReproduce string `json:"steps_to_reproduce"`
}

type comment struct {
    feedback  // Embedded feedback
    Text     string `json:"text"`
}

// Add more feedback types as needed...
// We can define new structs for each type of feedback without modifying the above code.

// displayFeedback function uses reflection to display any feedback
func displayFeedback(feedback Feedback) {  
    val := reflect.ValueOf(feedback)
    fmt.Printf("Feedback Type: %v\n", feedback.GetFeedbackType())

    // Use reflection to iterate over the feedback details
    for i := 0; i < val.NumField(); i++ {  
        field := val.Type().Field(i)
        value := val.Field(i)
        // Skip the embedded feedback field
        if field.Name != "feedback" {
            fmt.Printf("  %s: %v\n", field.Name, value)
        }
    }
    fmt.Println("--------")
}

func main() {  
  // Example feedback data in a map struct
  feedbacks := map[string]interface{}{
      "feedback_1": featureRequest{
          feedback: feedback{
              Type:    FeatureRequest,
              Details: map[string]interface{}{"user_id": 1},
          },
          Feature:   "New Search Feature",
          Benefit:  "Faster and more efficient search",
      },
      "feedback_2": bugReport{
          feedback: feedback{
              Type:    BugReport,
              Details: map[string]interface{}{"user_id": 2, "app_version": "1.0"},
          },
          Description:      "Application crashes when clicking button X",
          StepsToReproduce: "Open the app, click button X",
      },
  }

  // process and display each feedback dynamically using reflection
  for id, feedback := range feedbacks {
        switch v := feedback.(type) {
        case Feedback:
            displayFeedback(v)
        default:
            fmt.Printf("Invalid feedback type: %v\n", id)
        }
    }
}  