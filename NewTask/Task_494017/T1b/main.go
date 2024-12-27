package main

import (
	"fmt"
	"sync"
)

// func main() {
//     numbers := [][]int{
//         {1, 2, 3},
//         {4, 5},
//         {6, 7, 8, 9},
//     }

//     for _, row := range numbers {
//         for _, value := range row {
//             fmt.Println(value)
//         }
//     }
// }

// func main() {
//     numbers := [][]int{
//         {1, 2, 3},
//         {4, 5},
//         {6, 7, 8, 9},
//     }
//     wg := &sync.WaitGroup{}
//     wg.Add(len(numbers))
//     for _, row := range numbers {
//         go func() {
//             defer wg.Done()
//             for _, value := range row {
//                 fmt.Println(value)
//             }
//         }()
//     }
//     wg.Wait()
// }



func main() {
    numbers := [][]int{
        {1, 2, 3},
        {4, 5},
        {6, 7, 8, 9},
    }
    wg := &sync.WaitGroup{}
    wg.Add(len(numbers))
    for _, row := range numbers {
        go func(r []int) {
            defer wg.Done()
            for _, value := range r {
                fmt.Println(value)
            }
        }(row) // Pass the copy of the row value
    }
    wg.Wait()
}