package main  
import (  
    "fmt"
    "sync"
    "time"
)

// DataMiningTask represents a single data mining task
type DataMiningTask struct {
    Name string
}

// Run performs the data mining task
func (task DataMiningTask) Run(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Starting task: %s\n", task.Name)
    // Simulate task execution time
    time.Sleep(2 * time.Second)
    fmt.Printf("Task %s completed!\n", task.Name)
}

func main() {
    var wg sync.WaitGroup
    tasks := []DataMiningTask{
        {"Data Fetching"},
        {"Data Preprocessing"},
        {"Feature Extraction"},
        {"Model Training"},
        {"Result Analysis"},
    }

    for _, task := range tasks {
        wg.Add(1)
        go task.Run(&wg)
    }

    fmt.Println("Waiting for all tasks to finish...")
    wg.Wait()
    fmt.Println("All tasks completed. Research data mining process is done!")
}