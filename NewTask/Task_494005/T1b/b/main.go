package main  
import (  
    "fmt"
    "math/rand"
    "time"
)  
func main() {  
    rand.Seed(time.Now().UnixNano())  
    numbers := make([]int, 1000000)  
    for i := range numbers {  
        numbers[i] = rand.Int()  
    }  
    start := time.Now()  
    var sum int  
    for _, value := range numbers {  
        sum += value  
    }  
    elapsed := time.Since(start)  
    fmt.Println("Sum:", sum)  
    fmt.Println("Time taken:", elapsed)  
}  