package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

const maxWorkers = 4

func worker(wg *sync.WaitGroup, tasks chan int, startTime time.Time) {
	defer wg.Done()
	for task := range tasks {
		// Simulate CPU-bound work
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Println("Worker processed task:", task)
	}
	// Measure and log task completion time
	endTime := time.Now()
	totalTime := endTime.Sub(startTime).Seconds()
	fmt.Println("Total execution time:", totalTime, "seconds")
}

func getCPUUtilization() (float64, error) {
	cmd := exec.Command("top", "-b", "-n", "1")
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// This is a simplified way to extract CPU utilization.
	// In a real application, you'd want to parse the output correctly.
	lines := string(out)
	idleIdx := strings.Index(lines, "idle")
	if idleIdx != -1 {
		// Locate the colon following the idle keyword
		endIdleIdx := strings.Index(lines[idleIdx:], ":")
		if endIdleIdx != -1 {
			// Extract idle percentage as a string and parse it
			idleStr := strings.TrimSpace(lines[idleIdx+len("idle") : idleIdx+endIdleIdx])
			cpuUtil, err := strconv.ParseFloat(idleStr, 64)
			if err != nil {
				return 0, err
			}
			return 100 - cpuUtil, nil
		}
	}

	return 0, fmt.Errorf("could not find idle time in top output")
}

func main() {
	tasks := make(chan int, 10)
	var wg sync.WaitGroup
	startTime := time.Now()

	// Start workers
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(&wg, tasks, startTime)
	}

	// Generate tasks and send them to the channel
	for i := 0; i < 10; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()

	// Measure CPU utilization
	cpuUtilization, err := getCPUUtilization()
	if err != nil {
		fmt.Println("Error getting CPU utilization:", err)
	} else {
		fmt.Println("Average CPU utilization:", cpuUtilization, "%")
	}

	fmt.Println("All tasks completed.")
}
