package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	// Incorrect way: Deleting elements during iteration
	for i, num := range nums {
		if num%2 == 0 {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	fmt.Println(nums) // Output: [1 3 5 7]
}

// package main

// import "fmt"

// func main() {
// 	myMap := map[string]int{"a": 1, "b": 2, "c": 3}

// 	// Incorrect way: Iterating over values using delete during iteration
// 	for _, value := range myMap {
// 		delete(myMap, "b")
// 		fmt.Println("Value:", value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	myMap := map[string]int{"a": 1, "b": 2, "c": 3}

// 	// Correct way: Iterating over keys
// 	for key := range myMap {
// 		fmt.Println("Key:", key)
// 	}
// }
