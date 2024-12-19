package main  
import ("fmt")  
func main() {  
    mySlice := []int{1, 2, 3, 4, 5}  
    for i, _ := range mySlice {  
        if mySlice[i] > 3 {  
            mySlice = append(mySlice[:i], mySlice[i:]...)  
        }  
    }  
    fmt.Println(mySlice)  
}  