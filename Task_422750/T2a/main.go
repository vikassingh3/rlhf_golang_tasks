package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error error `json:"error"`
}

type authError interface {
	error
	HTTPStatusCode() int
}

type invalidCredentialsError struct {
	message string
}

func (e invalidCredentialsError) Error() string {
	return e.message
}

func (e invalidCredentialsError) HTTPStatusCode() int {
	return http.StatusUnauthorized
}

// Middleware for handling errors
func errorHandler(c *gin.Context) {
	if err := recover(); err != nil {
		e, ok := err.(authError)
		if ok {
			c.JSON(e.HTTPStatusCode(), errorResponse{Error: e})
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse{Error: fmt.Errorf("internal server error: %w", err)})
		}
	}
}
func main() {
	r := gin.Default()
	r.Use(errorHandler)

	// your routes
	r.GET("/protected", protectedRoute)

	r.Run(":8080")
}

func protectedRoute(c *gin.Context) {
	username := c.Request.Header.Get("Authorization")

	if username != "admin" {
		panic(invalidCredentialsError{"Invalid username"})
	}

	// Authenticated, proceed with route handler
	c.String(http.StatusOK, "Welcome!")
}
