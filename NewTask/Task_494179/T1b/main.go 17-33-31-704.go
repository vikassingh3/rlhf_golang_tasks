package main  
import "fmt"

func main() {  
    a, b := getMultiple()  
    fmt.Println(a, b)
}  
func getMultiple() (int, int) {  
    // Calculate values for a and b
    a := 10
    b := 20
    return a, b // Return by value
}



// package main  
// import "fmt"

// func main() {  
//     result := getAsStruct()  
//     fmt.Println(result.a, result.b)
// }  
// type result struct {
//     a int
//     b int
// }
// func getAsStruct() result {  
//     // Calculate values for a and b
//     a := 10
//     b := 20  
//     return result{a,b} 
// }