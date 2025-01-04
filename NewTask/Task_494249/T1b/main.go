
package main

import (
    "log"
    "os"
)

func main() {
    // Create a file handle
    logFile, err := os.OpenFile("./logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }
    defer logFile.Close()

    // Set the output of the logger to the file handle
    log.SetOutput(logFile)

    // Log some messages
    log.Println("This is a log message")
    log.Printf("This is a formatted log message: %s", "value")
    log.Fatalln("This is a fatal log message")
}
