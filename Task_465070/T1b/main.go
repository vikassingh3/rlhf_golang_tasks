package main  
import "fmt"
func sumAverage(numbers []int) (sum int, average float64) {
    total := 0
    count := len(numbers)
    for _, num := range numbers {
        total += num
    }
    sum = total
    if count > 0 {
        average = float64(total) / float64(count)
    }
    return
}
func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sum, average := sumAverage(numbers)
    fmt.Printf("Sum: %d, Average: %f\n", sum, average)  
}