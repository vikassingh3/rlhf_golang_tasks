package main

import (
	"fmt"
	"sync"
	"time"
)

func rateLimitedFunc(args ...int) <-chan int {
	ch := make(chan int, 10) // Buffer size of 10
	wg := sync.WaitGroup{}
	wg.Add(len(args))

	go func() {
		defer close(ch)
		for _, arg := range args {
			result := arg * arg
			wg.Done()
			time.Sleep(100 * time.Millisecond) // Introduce delay for rate limiting
			ch <- result
		}
	}()

	wg.Wait()
	return ch
}

func main() {
	results := rateLimitedFunc(1, 2, 3, 4, 5)
	for result := range results {
		fmt.Println(result)
	}
}


// func rateLimitedFunc(args ...int) []int {  
//     results := make([]int, 0, len(args))
//     for _, arg := range args {
//         result := arg * arg
//         results = append(results, result)
//     }
//     return results
// }
// func main() {  
//     results := rateLimitedFunc(1, 2, 3, 4, 5)
//     for _, result := range results {
//         fmt.Println(result)
//     }
// }