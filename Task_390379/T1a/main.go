package main

import (
    "log"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "time"
)

// Define a JWT secret key
var jwtSecret = []byte("your-secret-key")

// Struct to hold the token claims
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// IssueToken issues a new JWT token and logs the action
func IssueToken(username string) (string, error) {
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        log.Printf("ERROR: Issuing token for user %s: %v", username, err)
        return "", err
    }
    
    log.Printf("INFO: Token issued for user %s", username)
    return tokenString, nil
}

// ValidateToken validates the given JWT token and logs the action
func ValidateToken(tokenString string) (string, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        log.Printf("ERROR: Invalid token: %v", err)
        return "", err
    }
    
    if !token.Valid {
        log.Printf("WARN: Token is not valid")
        return "", err
    }

    log.Printf("INFO: Token validated for user %s", claims.Username)
    return claims.Username, nil
}

// HTTP handler for token issuance
func TokenHandler(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username") // Simple example, use a secure method in production
    token, err := IssueToken(username)
    if err != nil {
        http.Error(w, "Could not issue token", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(token))
}

// HTTP handler for token validation
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := r.URL.Query().Get("token")
    username, err := ValidateToken(tokenString)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    w.Write([]byte("Token valid for user: " + username))
}

func main() {
    http.HandleFunc("/issue", TokenHandler)
    http.HandleFunc("/validate", ValidateHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}