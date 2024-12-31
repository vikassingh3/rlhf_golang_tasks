package main  
import (  
    "fmt"
    "sync"
    "time"
)

func processData(wg *sync.WaitGroup, task string) {
    defer wg.Done()
    fmt.Printf("Started processing %s\n", task)
    time.Sleep(time.Second * 2) // Simulate processing time
    fmt.Printf("Completed processing %s\n", task)
}

func main() {
    var wg sync.WaitGroup
    tasks := []string{"task1", "task2", "task3"}

    wg.Add(len(tasks))
    for _, task := range tasks {
        go processData(&wg, task)
    }

    fmt.Println("Waiting for all tasks to complete...")
    wg.Wait()
    fmt.Println("All tasks completed successfully!")
}