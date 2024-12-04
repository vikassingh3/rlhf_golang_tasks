package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Use sync.Map for thread-safe operations
var blockList sync.Map

// Fetch the JWT secret key from an environment variable
func getSecretKey() []byte {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if len(secretKey) == 0 {
		log.Fatal("JWT_SECRET_KEY is not set")
	}
	return []byte(secretKey)
}

func isRevoked(token string) bool {
	_, exists := blockList.Load(token)
	return exists
}

// Validate the JWT token
func validateToken(tokenString string) (*jwt.Token, error) {
	secretKey := getSecretKey()

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return secretKey, nil
	})
	return token, err
}

func main() {
	router := mux.NewRouter()

	// Replace with a strong secret key management solution
	router.HandleFunc("/get-token", getTokenHandler).Methods("GET")
	router.HandleFunc("/protected", protectedHandler).Methods("GET")
	router.HandleFunc("/block-token", blockTokenHandler).Methods("POST")

	// Run on HTTPS in production
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Sample user data
	p := &Person{ID: "1", Name: "Test User"}

	secretKey := getSecretKey()

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = p
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token
	tknString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tknString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,                    // Enable HTTPS
		SameSite: http.SameSiteStrictMode, // Add SameSite for CSRF protection
		Expires:  time.Now().Add(time.Hour * 24),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token Created"))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve token from the cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value

	// Validate the token
	token, err := validateToken(tokenString)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch ve.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			case jwt.ValidationErrorExpired:
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the token is revoked
	if isRevoked(tokenString) {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// Token is valid
	w.WriteHeader(http.StatusOK)
	fmt.Println(token)
	w.Write([]byte("This is a protected resource"))
}

func blockTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
		return
	}

	// Retrieve token from the form
	token := r.FormValue("token")
	if token != "" {
		// Block the token by adding it to the block list
		blockList.Store(token, struct{}{})
		log.Printf("Blocking token: %s", token)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token blocked successfully."))
}
