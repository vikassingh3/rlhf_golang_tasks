package plugin_feedback

import (
	"task/main" // Path to the main package
)

type PluginFeedback struct {
	main.feedback  // Embed the feedback struct for type and common details
	Suggestion     string `json:"suggestion"`
}

// Implement the Feedback interface
func (pf *PluginFeedback) GetFeedbackType() main.FeedbackType {
	return main.FeedbackType(pf.Type)
}

func (pf *PluginFeedback) Display() string {
	return fmt.Sprintf("Suggestion: %s", pf.Suggestion)
}

// This function must be exported by the plugin package
func FeedbackCreator() main.Feedback {
	return &PluginFeedback{
		Type:   "Suggestion",
		Details: map[string]interface{}{"user_id": 3, "suggestion": "Improve UI design"},
	}
}