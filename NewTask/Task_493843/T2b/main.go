package main

import "fmt"

// func main() {
//     var slice []int // Nil slice
//     for _, val := range slice {
//         fmt.Println(val)
//     }
// }




func main() {
    slice := []int{1, 2, 3}
    ptr := &slice
    for _, val := range *ptr {
        fmt.Println(val)
    }
    
    // Modify the slice through the pointer
    *ptr = append(*ptr, 4)
    
    for _, val := range *ptr {
        fmt.Println(val)
    }
}


// func main() {
//     str := "Hello, 世界"
//     for index, rune := range str {
//         fmt.Printf("index: %d, rune: %c, UTF-8 bytes: % x\n", index, rune, []byte(string(rune)))
//     }
// }
