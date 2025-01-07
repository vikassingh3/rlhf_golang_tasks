package main

import (
	"fmt"
	"time"
)

// func main() {
//    a := []int{1, 2, 3, 4, 5}
//    for i, v := range a {
//       fmt.Printf("Index: %d, Value: %d\n", i, v)
//    }
// }

// package main
// import (
//     "fmt"
//     "time"
// )

const largeDatasetSize = 1000000  

func main() {  
    a := make([]int, largeDatasetSize)
    for i := 0; i < largeDatasetSize; i++ {  
        a[i] = i
    }
    
    // Using range loop
    start := time.Now()
    var sum int
    for _, v := range a {  
        sum += v
    }
    fmt.Println("Sum using range loop:", sum)
    elapsed := time.Since(start)
    fmt.Println("Time taken using range loop:", elapsed)
    
    // Using for loop
    start = time.Now()
    sum = 0
    for i := 0; i < largeDatasetSize; i++ {  
        sum += a[i]
    }
    fmt.Println("Sum using for loop:", sum)
    elapsed = time.Since(start)
    fmt.Println("Time taken using for loop:", elapsed)
}


// package main  
// import "fmt"

// func main() {
//    words := []string{"Hello", "Go", "Lang"}
//    for _, word := range words {  
//       fmt.Println("Hello", word)
//    }
// }