// auth/user.go
package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID          string
	Username    string
	Password    string
	JWTToken    *jwt.Token
	TokenString string
}

func NewUser(id, username, password string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var username string
	var password string

	err := json.NewDecoder(r.Body).Decode(&struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Here, you should validate the user's credentials against a database or storage
	user := NewUser("user1", username, password) // Replace this with actual user authentication

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires after 24 hours

	tokenString, err := token.SignedString([]byte(main.jwtSecret))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.TokenString = tokenString
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
