package main  
import (  
    "fmt"
    "log"
    "os"
    "strings"
    "time"
)

type LogRecord struct {
    Timestamp   time.Time
    Message     string
    Severity    string
    Sensitive   bool
}

func (l *LogRecord) String() string {
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("%s [%s] ", l.Timestamp.Format(time.RFC3339), l.Severity))
    if l.Sensitive {
        sb.WriteString("*** Sensitive Data ***")
    } else {
        sb.WriteString(l.Message)
    }
    return sb.String()
}

func main() {
    // Create a log file for non-sensitive data
    nonSensitiveLogFile, err := os.OpenFile("logs/non_sensitive.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening non-sensitive log file: %v", err)
    }
    defer nonSensitiveLogFile.Close()

    // Create a log file for sensitive data (redacted)
    sensitiveLogFile, err := os.OpenFile("logs/sensitive.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening sensitive log file: %v", err)
    }
    defer sensitiveLogFile.Close()

    // Logging function with GDPR compliance in mind
    logMessage := func(message string, sensitive bool) {
        logRecord := LogRecord{
            Timestamp:   time.Now(),
            Message:     message,
            Severity:    "INFO",  // You can set severity based on the message
            Sensitive:   sensitive,
        }

        // Log to both files, but redact sensitive data in the non-sensitive log
        log.SetOutput(nonSensitiveLogFile)
        if sensitive {
            log.Println(logRecord.String())
        } else {
            log.Println(logRecord.Message)
        }

        log.SetOutput(sensitiveLogFile)
        log.Println(logRecord.String())
    }

    // Example log entries
    logMessage("User logged in with username: user123", false)
    logMessage("User's payment card number: 4111-1111-1111-1111", true)
    logMessage("An error occurred: invalid request data", false)
    logMessage("Sensitive API key: abcdef1234567890", true)
}  
