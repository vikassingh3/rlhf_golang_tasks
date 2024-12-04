package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// User represents a user model
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// secretKey is used to sign the JWT tokens
const secretKey = "mySecretKey" // Please replace this with a more secure key in a production environment

var users = []User{
	{
		ID:   "1",
		Name: "john",
	},
	{
		ID:   "2",
		Name: "jane",
	},
}

func main() {
	r := gin.Default()

	// Authentication route
	r.POST("/auth", authenticate)

	// Protected route that requires a valid JWT
	r.GET("/protected", protect, func(c *gin.Context) {
		userID := c.GetString("user_id")
		user := getUserById(userID)
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + user.Name + "!"})
	})

	r.Run(":8080")
}

func authenticate(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Implement proper user authentication logic here
	var authenticatedUser User
	for _, user := range users {
		if user.Name == username && user.Name == password {
			authenticatedUser = user
			break
		}
	}

	if authenticatedUser.ID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token := generateToken(authenticatedUser)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken(user User) string {
	claims := &jwt.StandardClaims{
		Issuer:    "your-app-name",
		Subject:   user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return signedToken
}

func protect(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token required"})
		c.Abort()
		return
	}

	// Verify the token
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		if verr, ok := err.(*jwt.ValidationError); ok {
			if verr.Errors&jwt.ValidationErrorMalformed != 0 {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Token is malformed"})
			} else if verr.Errors&jwt.ValidationErrorExpired != 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is expired"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is invalid"})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}
		c.Abort()
		return
	}

	c.Set("user_id", claims.Subject)
}

func getUserById(userID string) User {
	for _, user := range users {
		if user.ID == userID {
			return user
		}
	}
	return User{}
}