package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomError represents a custom error structure for the API.
type CustomError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Detail string `json:"detail,omitempty"`
}

// Error returns the error message as a string.
func (e CustomError) Err() string {
	return fmt.Sprintf("status: %d, error: %s", e.Status, e.Error)
}

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Bind JSON input to credentials struct
		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, CustomError{
				Status: http.StatusBadRequest,
				Error:  "Invalid request format",
			})
			return
		}

		// Authenticate user logic
		if credentials.Username != "admin" || credentials.Password != "secret" {
			c.JSON(http.StatusUnauthorized, CustomError{
				Status: http.StatusUnauthorized,
				Error:  "Invalid credentials",
			})
			return
		}

		// Respond with success
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	})

	r.GET("/protected", func(c *gin.Context) {
		// Check if user is authenticated (e.g., using a header for simplicity)
		token := c.GetHeader("X-Auth-Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, CustomError{
				Status: http.StatusUnauthorized,
				Error:  "Authentication required",
			})
			return
		}

		c.String(http.StatusOK, "Welcome to the protected route!")
	})

	r.Run(":8080")
}
