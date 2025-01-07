package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// Configure Logrus logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(logrus.StandardLogger().Out)

	// Log an error
	logrus.WithFields(logrus.Fields{
		"service": "my-service",
		"method":  "handleUserTransaction",
	}).Error("Error occurred during transaction")
}
