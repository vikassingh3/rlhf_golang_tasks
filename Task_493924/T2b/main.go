package main

import (
    "fmt"
)

type UserData struct {
    Name string
    Age  int
}

func GetUserData(userID int) UserData {
    var userData UserData

    // Some logic to retrieve user data from database or other sources
    // For simplicity, we'll just hardcode some data
    if userID == 1 {
        userData.Name = "Alice"
        userData.Age = 25
    }
    return userData
}

func main() {
    userID := 1
    userData := GetUserData(userID)

    // Check if user exists (but you have to add an extra condition if you want to detect an error)
    if userData.Name != "" {
        fmt.Println("User Name:", userData.Name)
        fmt.Println("User Age:", userData.Age)        
    } else {
        fmt.Println("User not found.")
    }
}