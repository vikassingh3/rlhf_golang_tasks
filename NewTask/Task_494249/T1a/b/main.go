package main

import (
	"fmt"
	"log"
	"os"
)

type customLogger struct {
	logger *log.Logger
}

func newCustomLogger(filename string) *customLogger {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	return &customLogger{
		logger: log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (cl *customLogger) Log(message string) {
	cl.logger.Println(message)
}

func main() {
	logger := newCustomLogger("app.log")

	logger.Log("This is a custom log message.")
}