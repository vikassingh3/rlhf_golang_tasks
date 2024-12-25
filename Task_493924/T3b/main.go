package main  
import (  
    "fmt"  
    "errors"  
)  

// Function that returns a struct and an error
func divide(dividend float64, divisor float64) (float64, error) {  
    if divisor == 0 {  
        return 0, errors.New("divisor cannot be zero")  
    }  
    result := dividend / divisor  
    return result, nil  
}  

func main() {  
    result, err := divide(10.0, 0.0)  
    if err != nil {  
        fmt.Println("Error:", err)  
    } else {  
        fmt.Println("Result:", result)  
    }  
}  