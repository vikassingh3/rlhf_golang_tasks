package main  

import (  
    "fmt"
)

func getData() (string, int) {  
    return "Hello", 42  
}  

func main() {  
    // Call the function and store the return value in a tuple
    result, _ := getData()  
    
    // Convert result into an interface{}
    var data interface{} = result  

    // Type assertion to extract values from the tuple
    str, ok := data.([2]interface{})  // Assert it to a tuple of 2 elements
    if !ok {
        fmt.Println("Type assertion failed")
        return
    }
    
    // Access values using type assertion
    strValue, ok := str[0].(string)
    numValue, ok := str[1].(int)

    if !ok {
        fmt.Println("Type assertion failed")
        return
    }

    // Print the values
    fmt.Println("String:", strValue)  
    fmt.Println("Number:", numValue)
}
