package main  
import (  
    "fmt"
)  
type Logger struct {  
    logCallback func(string)  
}  
// SetLogCallback sets the callback function for logging.  
func (l *Logger) SetLogCallback(callback func(string)) {  
    l.logCallback = callback  
}  
// Log logs a message using the set callback.  
func (l *Logger) Log(message string) {  
    if l.logCallback != nil {  
        l.logCallback(message)  
    }  
}  
func consoleLogger(message string) {  
    fmt.Println("Console Log:", message)  
}  
func fileLogger(message string) {  
    fmt.Println("File Log:", message)  
}  
func main() {  
    // Create a Logger instance.  
    logger := &Logger{}  
    // Set the console logger as the callback.  
    logger.SetLogCallback(consoleLogger)  
    logger.Log("Hello from console logger!")  
    // Set the file logger as the callback.  
    logger.SetLogCallback(fileLogger)  
    logger.Log("Hello from file logger!")  
    // You can even set different callbacks for different types of logs.  
    logger.SetLogCallback(consoleLogger)  
    logger.Log("Important info goes to console.")  
    logger.SetLogCallback(fileLogger)  
    logger.Log("Debug details go to file.")  
}  