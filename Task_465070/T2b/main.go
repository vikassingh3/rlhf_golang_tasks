package main  
import (  
    "fmt"
    "strconv"
)  
func convertToInt(input string) (int, error) {  
    result, err := strconv.Atoi(input)  
    return result, err  
}  
func main() {  
    input := "42"
    number, err := convertToInt(input)  
    if err != nil {  
        fmt.Println("Error converting:", err)  
    } else {  
        fmt.Println("Converted number:", number)  
    }  
}