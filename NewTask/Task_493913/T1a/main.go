package main

import (
	"fmt"
)

// Feedback struct represents a generic feedback structure.
type Feedback struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// UserReview struct represents a user review.
type UserReview struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

// BugReport struct represents a bug report.
type BugReport struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// processFeedback processes feedback of any type.
func processFeedback(feedback Feedback) {
	feedbackType := feedback.Type
	content := feedback.Content

	switch feedbackType {
	case "user_review":
		// Use type assertion to access the actual struct.
		if review, ok := content.(UserReview); ok {
			processUserReview(review)
		} else {
			fmt.Println("Invalid content for user_review")
		}
	case "bug_report":
		// Use type assertion to access the actual struct.
		if report, ok := content.(BugReport); ok {
			processBugReport(report)
		} else {
			fmt.Println("Invalid content for bug_report")
		}
	default:
		fmt.Println("Unsupported feedback type:", feedbackType)
	}
}

func processUserReview(review UserReview) {
	fmt.Printf("User Review:\nRating: %d, Comment: %s\n", review.Rating, review.Comment)
}

func processBugReport(report BugReport) {
	fmt.Printf("Bug Report:\nTitle: %s, Description: %s\n", report.Title, report.Description)
}

func main() {
	feedbackMap := map[string]Feedback{
		"user_review": {
			Type: "user_review",
			Content: UserReview{
				Rating:  5,
				Comment: "Great service!",
			},
		},
		"bug_report": {
			Type: "bug_report",
			Content: BugReport{
				Title:       "Login issue",
				Description: "Unable to log in with some accounts.",
			},
		},
	}

	for _, feedback := range feedbackMap {
		processFeedback(feedback)
	}
}
