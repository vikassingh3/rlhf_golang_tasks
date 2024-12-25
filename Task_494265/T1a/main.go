package main

import "fmt"


func main(){
s := []int{1, 2, 3}
for i, _ := range s {
    fmt.Println(i, s[i]) // Print indices to check out-of-range access
    s = append(s, 4)     // Modifying slice while iterating can cause issues
}


// s := []int{1, 2, 3, 4, 5}
// for _, v := range s {
//     if v > 3 {
//         break
//     }
//     fmt.Println(v)
// }


// s := []int{1, 2, 3}
// for i, _ := range s {
//     if s[i] == 2 {
//         delete(s, i) // Deleting while iterating is not allowed
//     }
// }


// var count int
// for _, v1 := range s1 {
//     for _, v2 := range s2 {
//         if v1 == v2 {
//             count++
//         }
//     }
// }
// fmt.Println("Count:", count)
}