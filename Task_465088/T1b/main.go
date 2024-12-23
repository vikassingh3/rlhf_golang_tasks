package main  
import ("fmt")

func main() {  
    // Declare a map
    var userMap map[string]int  

    // Initialize the map
    userMap = make(map[string]int)  
    
    // Store key-value pairs
    userMap["Alice"] = 25
    userMap["Bob"] = 30
    userMap["Charlie"] = 22

    // Retrieve values using keys
    fmt.Println("Alice's age:", userMap["Alice"])  
    fmt.Println("Bob's age:", userMap["Bob"])  

    // Check if a key exists
    if _, ok := userMap["David"]; ok {
        fmt.Println("David's age:", userMap["David"])  
    } else {
        fmt.Println("David is not found.")
    }
    
    // Delete a key-value pair
    delete(userMap, "Charlie")
}