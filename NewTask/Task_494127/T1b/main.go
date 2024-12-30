package main  
import (  
    "fmt"
)  
func main() {  
    // Initialize a nested map
    nestedMap := make(map[string]map[string]int)
    
    // Populate the nested map manually
    nestedMap["country1"] = make(map[string]int)
    nestedMap["country1"]["city1"] = 100
    nestedMap["country1"]["city2"] = 200
    
    nestedMap["country2"] = make(map[string]int)
    nestedMap["country2"]["city3"] = 300
    nestedMap["country2"]["city4"] = 400
    nestedMap["country2"]["city5"] = 500

    // Access and print values from the nested map
    fmt.Println("Country1:")
    fmt.Println("City1:", nestedMap["country1"]["city1"])
    fmt.Println("City2:", nestedMap["country1"]["city2"])
    
    fmt.Println("\nCountry2:")
    fmt.Println("City3:", nestedMap["country2"]["city3"])
    fmt.Println("City4:", nestedMap["country2"]["city4"])
    fmt.Println("City5:", nestedMap["country2"]["city5"])
} 