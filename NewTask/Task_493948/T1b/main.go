
package main  
import (  
    "fmt"
    "sync"
    "time"
)  
//function to aggregate data
func aggregateData(data []int, wg *sync.WaitGroup, result chan int) {  
    defer wg.Done()
    total := 0
    for _, value := range data {
        total += value
    }
    result <- total
}  
func main() {  
    //create a waitgroup
    var wg sync.WaitGroup  
    //number of tasks
    const numTasks = 3
    //initialize result channel
    result := make(chan int, numTasks)
    //start the timer
    start := time.Now()
    wg.Add(numTasks)
    //create data for aggregation
    data1 := []int{1, 2, 3, 4, 5}
    data2 := []int{6, 7, 8, 9, 10}
    data3 := []int{11, 12, 13, 14, 15}
    //run the aggregation tasks in parallel
    go aggregateData(data1, &wg, result)
    go aggregateData(data2, &wg, result)
    go aggregateData(data3, &wg, result)
    //wait for all tasks to complete
    wg.Wait()
    //close the result channel
    close(result)
    fmt.Println("Total time taken:", time.Since(start))
    total := 0
    //aggregate the results from the channel
    for r := range result {
        total += r
    }
    fmt.Println("Final Aggregated Result:", total)
}  
