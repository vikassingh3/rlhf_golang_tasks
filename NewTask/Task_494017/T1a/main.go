package main

import (
	"fmt"
	"sync"
)

// func main() {
//     matrix := [][]int{
//         {1, 2, 3},
//         {4, 5, 6},
//         {7, 8, 9},
//     }

//     for i, row := range matrix {
//         for j, value := range row {
//             fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
//         }
//     }
// }

// package main

// import (
//     "fmt"
//     "sync"
// )

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    var wg sync.WaitGroup
    wg.Add(len(matrix))

    for i, row := range matrix {
        go func(i int, row []int) {
            for j, value := range row {
                // Use value, e.g., compute something or modify it
                // Access shared data may require synchronization
				fmt.Println(j, value)
            }
            wg.Done()
        }(i, row)
    }

    wg.Wait()
    fmt.Println("All goroutines done.")
}


// var sharedData int
// var mutex = &sync.Mutex{}

// func updateSharedData(value int) {
//     mutex.Lock()
//     defer mutex.Unlock()
//     sharedData += value
// }

// // Inside the goroutine
// for j, value := range row {
//     // ....
//     updateSharedData(value)
// }