
package main

import (  
    "fmt"
    "io"
    "os"
)

func main() {  
    //open a file for reading  
    file, err := os.Open("example.txt")  
    if err != nil {  
        fmt.Println("Error opening file: ", err)  
        return  
    }  
    //close the file when we are done  
    defer file.Close()  
  
    //create a new buffer to store data  
    buffer := make([]byte, 1024)  
    for {  
        //read from the file into the buffer  
        numBytesRead, err := file.Read(buffer)  
        if err != nil {  
            if err == io.EOF {  
                fmt.Println("Reached the end of the file")  
                break  
            }  
            fmt.Println("Error reading file: ", err)  
            return  
        }  
        //convert the bytes to string and print  
        data := string(buffer[:numBytesRead])  
        fmt.Println(data)  
    }  
}
