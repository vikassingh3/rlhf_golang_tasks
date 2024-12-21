package main  
import (  
    "fmt"
    "time"
)  
type Logger struct {  
    logCallback func(string, time.Duration)  
    startTime   time.Time  
}  

// SetLogCallback sets the callback function for logging.  
func (l *Logger) SetLogCallback(callback func(string, time.Duration)) {  
    l.logCallback = callback  
}  
// Log logs a message using the set callback.  
func (l *Logger) Log(message string) {  
    if l.logCallback != nil {  
        l.logCallback(message, 0)  
    }  
}  

// StartProfiling starts profiling the logged action.
func (l *Logger) StartProfiling() {
    l.startTime = time.Now()
}

// StopProfiling stops profiling the logged action and logs the execution time.
func (l *Logger) StopProfiling(message string) {
    elapsed := time.Since(l.startTime)
    if l.logCallback != nil {
        l.logCallback(message, elapsed)
    }
}

func consoleLogger(message string, duration time.Duration) {
    if duration > 0 {
        fmt.Println("Console Log:", message, "Execution Time:", duration)
    } else {
        fmt.Println("Console Log:", message)
    }
}

func fileLogger(message string, duration time.Duration) {
    if duration > 0 {
        fmt.Println("File Log:", message, "Execution Time:", duration)
    } else {
        fmt.Println("File Log:", message)
    }
}
func main() {  
    logger := &Logger{}  
    logger.SetLogCallback(consoleLogger)  
    
    logger.Log("Doing some work...")  
    
    // Start profiling the loop
    logger.StartProfiling()
    for i := 0; i < 1000000; i++ {
    }
    // Stop profiling and log the execution time
    logger.StopProfiling("Finished counting to 1 million")
    
    logger.SetLogCallback(fileLogger)  
    logger.Log("Debug details...")  
}  