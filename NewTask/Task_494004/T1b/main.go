package main

import (
	"io"
	"log"
	"os"
	"testing"
)

// These are the different file-based logging strategies we'll benchmark.
var loggingStrategies = []struct {
	name string
	logger *log.Logger
}{
	{
		name: "DefaultLogger",
		logger: log.New(os.Stdout, "", 0),
	},
	{
		name: "LshortfileLogger",
		logger: log.New(os.Stdout, "", log.Lshortfile),
	},
	{
		name: "LlongfileLogger",
		logger: log.New(os.Stdout, "", log.Llongfile),
	},
	{
		name: "LdateLogger",
		logger: log.New(os.Stdout, "", log.Ldate),
	},
	{
		name: "LtimeLogger",
		logger: log.New(os.Stdout, "", log.Ltime),
	},
	{
		name: "LmicrosecondsLogger",
		logger: log.New(os.Stdout, "", log.Lmicroseconds),
	},
}

func BenchmarkLogging(b *testing.B) {
	// Initialize an output file
	f, err := os.Create("logfile.log")
	if err != nil {
		b.Fatalf("Error creating log file: %v", err)
	}
	defer f.Close()

	// Create a multi-writer to write logs to both stdout and file
	multiWriter := io.MultiWriter(os.Stdout, f)

	for _, strategy := range loggingStrategies {
		// Redirect the logger to use the multi-writer
		strategy.logger.SetOutput(multiWriter)
		b.Run(strategy.name, func(b *testing.B) {
			// Log various levels of messages to simulate different logging scenarios
			for i := 0; i < b.N; i++ {
				strategy.logger.Println("This is an informational message")
				strategy.logger.Printf("Error: %v\n", "An error occurred!")
				strategy.logger.Println("Warning: Something went wrong")
				strategy.logger.Println("Debug: Debug information goes here")
			}
		})
	}
}
 