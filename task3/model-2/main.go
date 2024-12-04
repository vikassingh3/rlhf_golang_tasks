package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User is the database model
type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment"`
	Email    string `gorm:"unique"`
	Password string
}

// db will hold the pointer to our database connection
var db *sql.DB

func init() {
	// Database connection (replace "your_database_username": "your_database_password"@tcp(127.0.0.1:3306)/auth_db?charset=utf8&parseTime=True&loc=Local" with your MySQL connection string)
	var err error
	db, err = sql.Open("mysql", "your_database_username", "your_database_password@tcp(127.0.0.1:3306)/auth_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("Couldn't connect to db: %v", err)
	}
}

func main() {
	// Automigrate the User model
	fmt.Println("Creating tables if they don't exist...")
	// gorm.AutoMigrate creates tables if they don't exist
	if err := gorm.New(db).AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
	}

	// RUN THE SERVER
	runServer()
}

// User struct
type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment"`
	Email    string `gorm:"unique"`
	Password string
}

// Login struct represents the login request body
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostLogin - Authenticate user via email/password
func PostLogin(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var login Login
	if err := dec.Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate login credentials
	user := &User{}
	db.Where("email = ?", login.Email).First(user)
	if user.Email == "" {
		http.Error(w, "Invalid Email/Password", http.StatusUnauthorized)
		return
	}

	if user.Password != login.Password {
		http.Error(w, "Invalid Email/Password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := generateToken(user.ID)
	if token == "" {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth-token",
		Value:    token,
		HttpOnly: true, // The cookie can't be accessed by JavaScript
		Path:     "/",
		Expires:  nextMonth(), // Ensure token expires in a reasonable time frame (e.g., 1 month in this case)
		Secure:   true,        // Requires HTTPS for sending and accessing the cookie
	})

	response := map[string]string{
		"message": "Successfully logged in!",
		"token":   token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func nextMonth() time.Time {
	now := time.Now()
	return now.AddDate(0, 1, 0)
}

const secretKey = "thisIsAnVerySecretKey" // In production, use a stronger and more securely stored secret key

// generateToken creates a new jwt token
func generateToken(userID uint64) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token will expire in 24 hours (customize as needed)

	tkn, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ""
	}
	return tkn
}
