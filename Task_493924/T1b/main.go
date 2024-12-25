package main  
import (  
    "fmt"
    "errors"
)

// Define a struct to hold the result
type Result struct {  
    Name string  
    Age  int  
}  

// Function that returns a struct and an error
func GetUserData(userID int) (*Result, error) {  
    // Simulate a database lookup
    if userID == 1 {  
        return &Result{Name: "Alice", Age: 25}, nil  
    } else if userID == 2 {  
        return nil, errors.New("user not found")  
    } else {  
        return nil, errors.New("invalid user ID")  
    }  
}  

func main() {  
    // Calling the function
    result, err := GetUserData(1)  
    if err != nil {  
        // Handle the error
        fmt.Println("Error:", err)  
    } else {  
        // Access the result struct
        fmt.Println("User Name:", result.Name)  
        fmt.Println("User Age:", result.Age)  
    }  
}