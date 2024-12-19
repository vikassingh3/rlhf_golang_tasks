package main

import (
	"fmt"
)

func removeDuplicates(s []int) {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				// Modifying the slice during iteration
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
}

func main() {
	data := []int{1, 2, 2, 3, 4, 4, 5}

	removeDuplicates(data)

	fmt.Println(data)

	data2 := []int{1, 2, 2, 3, 4, 4, 5}

	uniqueData := removeDuplicates2(data2)

	fmt.Println(uniqueData)
}

func removeDuplicates2(s []int) []int {
	seen := make(map[int]bool)
	var unique []int
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}
