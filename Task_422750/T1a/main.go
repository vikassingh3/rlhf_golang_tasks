package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthResponse is the response structure for authentication errors.
type AuthResponse struct {
	Message string `json:"message"`
}

// BasicAuthMiddleware performs basic authentication.
func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.Header.Get("Authorization")
		if username == "" {
			c.JSON(http.StatusUnauthorized, AuthResponse{Message: "Username is required"})
			c.Abort()
			return
		}

		// For demonstration purposes, we'll just check if the username is "admin".
		if username != "admin" {
			c.JSON(http.StatusForbidden, AuthResponse{Message: "Invalid username"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Apply the BasicAuthMiddleware to protected routes
	r.GET("/protected", BasicAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the protected route!")
	})

	r.GET("/unprotected", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the unprotected route!")
	})

	r.Run(":8080")
}
