package main

import (
	"fmt"
	"plugin"
	"reflect"
)

type FeedbackType string

const (
	FeatureRequest FeedbackType = "FeatureRequest"
	BugReport      FeedbackType = "BugReport"
	Comment        FeedbackType = "Comment"
)

// Feedback interface encapsulates feedback functionality
type Feedback interface {
	GetFeedbackType() FeedbackType
	Display() string
}

var plugins map[FeedbackType]Feedback = make(map[FeedbackType]Feedback)

// RegisterFeedback registers a feedback type with the system
func RegisterFeedback(fb Feedback) {
	if fb != nil {
		plugins[fb.GetFeedbackType()] = fb
	}
}

// FeedbackCreator function is expected to be exported by each plugin package
func FeedbackCreator() Feedback {
	return nil
}

// main function loads and registers plugins
func main() {
	// Load plugins from separate packages
	if err := loadPlugins(); err != nil {
		fmt.Printf("Error loading plugins: %v\n", err)
		return
	}

	// Example feedback data
	feedbacks := []Feedback{
		&FeatureRequest{
			Type:   FeatureRequest,
			Details: map[string]interface{}{"user_id": 1, "feature": "New Search Feature", "benefit": "Faster and more efficient search"},
		},
		&BugReport{
			Type:    BugReport,
			Details: map[string]interface{}{"user_id": 2, "app_version": "1.0", "description": "Crash when clicking button X", "steps_to_reproduce": "Open app, click button X"},
		},
		// Add new feedback types dynamically
		&PluginFeedback{
			Type:   "Suggestion",
			Details: map[string]interface{}{"user_id": 3, "suggestion": "Improve UI design"},
		},
	}

	// Process each feedback dynamically using reflection
	for _, feedback := range feedbackbacks {
		fmt.Printf("Feedback Type: %v\n", feedback.GetFeedbackType())
		fmt.Println(feedback.Display())
		fmt.Println("--------")
	}
}

// loadPlugins loads and registers all plugins in the GOPATH
func loadPlugins() error {
	gopath := getGOPATH()
	files, err := listFiles(gopath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file == "main.go" {
			continue // Skip the main package
		}

		// Load the plugin from the file path
		p, err := plugin.Open(file)
		if err != nil {
			fmt.Printf("Error loading plugin: %v\n", err)
			continue
		}

		// Look for the FeedbackCreator symbol in the plugin
		fc, err := p.Lookup("FeedbackCreator")
		if err != nil {
			fmt.Printf("Error looking up FeedbackCreator in plugin: %v\n", err)
			continue
		}

		// Create and register the feedback
		feedback := fc.(func() Feedback)()
		RegisterFeedback(feedback)
	}

	return nil
}

// getGOPATH retrieves the GOPATH environment variable
func getGOPATH() string {
	gopath := "." // Default to current directory if GOPATH is not set
	if gopathEnv := os.Getenv("GOPATH"); gopathEnv != "" {
		gopath = gopathEnv
	}
	return gopath
}

// listFiles lists all Go source files in the specified directory
func listFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var goFiles []string
	for _, file := range files {
		if file.Name() == "main.go" || !strings.HasSuffix(file.Name(), ".go") {
			continue // Skip main.go and non-Go files
		}
		goFiles = append(goFiles, filepath.Join(dir, file.Name()))
	}

	return goFiles, nil
}