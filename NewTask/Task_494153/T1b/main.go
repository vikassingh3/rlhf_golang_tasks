package main  
import (  
    "fmt"
)  

func getData() (string, int) {  
    return "Hello", 42  
}  

func main() {  
    // Call the function and store the return value in a tuple  
    result, data := getData()  

    // Use type assertion to extract each value from the tuple and assign it to a variable  
    str, num := result, data  

    // Print the values  
    fmt.Println("String:", str)  
    fmt.Println("Number:", num)  
}  