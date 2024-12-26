package main

import (
	"fmt"
	"sync"
)

// func main() {
//     slice := []int{1, 2, 3}
//     for _, val := range slice {
//         slice = append(slice, val+10) // Race condition!
//         fmt.Println(val)
//     }
// }

// func main() {
//     mapA := map[string]int{"a": 1, "b": 2}
//     for key, _ := range mapA {
//         mapA["c"] = 3 // Race condition!
//         fmt.Println(key)
//     }
// }

// func main() {
//     array := [3]int{1, 2, 3}
//     slice := array[:]

//     for index, value := range array {
//         fmt.Printf("array[%d]: %d\n", index, value)
//     }

//     fmt.Println()
//     for index, value := range slice {
//         fmt.Printf("slice[%d]: %d\n", index, value)
//     }
// }





func main() {
	var wg sync.WaitGroup
	mapA := make(map[int]int)
	var mu sync.Mutex // Mutex to prevent data races

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()          // Lock before accessing the map
			mapA[i] = i        // Safely write to the map
			mu.Unlock()        // Unlock after modifying the map
		}(i) // Pass `i` as a parameter to avoid closure issues
	}

	wg.Wait()
	mu.Lock() // Lock to safely read the map
	for key, val := range mapA {
		fmt.Printf("%d: %d\n", key, val)
	}
	mu.Unlock()
}

