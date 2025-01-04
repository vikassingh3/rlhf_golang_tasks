package main

import (
	"io"
	"log"
	"os"
	"testing"
	"time"
)

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
			totalDuration := time.Duration(0)

			for i := 0; i < b.N; i++ {
				start := time.Now()

				strategy.logger.Println("This is an informational message")
				strategy.logger.Printf("Error: %v\n", "An error occurred!")
				strategy.logger.Println("Warning: Something went wrong")
				strategy.logger.Println("Debug: Debug information goes here")

				end := time.Now()
				totalDuration += end.Sub(start)
			}

			// Calculate latency (average time per operation)
			latency := totalDuration / time.Duration(b.N)
			b.Logf("Latency: %s", latency)

			// Calculate throughput (operations per second)
			throughput := float64(b.N) / float64(totalDuration.Seconds())
			b.Logf("Throughput: %.2f ops/s", throughput)
		})
	}
}
