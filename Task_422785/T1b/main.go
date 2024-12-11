package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

// User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// GenerateRandomBytes function
func generateRandomBytes(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// CreateJWT function
func createJWT(user *User) (string, error) {
	secretKey := generateRandomBytes(32) // Use a more secure method to load secret key in production

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyJWT function
func verifyJWT(tokenString string) (*User, error) {
	secretKey := generateRandomBytes(32) // Use a more secure method to load secret key in production

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &User{
			ID:       claims["id"].(string),
			Username: claims["username"].(string),
		}
		return user, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// EncryptData function
func encryptData(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte("your-aes-key-here")) // Use a strong and secure key
	if err != nil {
		return nil, err
	}

	if len(data)%block.BlockSize() != 0 {
		data = padData(data, block.BlockSize())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	return append(nonce, ciphertext...), nil
}

// DecryptData function
func decryptData(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize+aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	block, err := aes.NewCipher([]byte("your-aes-key-here")) // Use the same key as encryption
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:aesgcm.NonceSize()], ciphertext[aesgcm.NonceSize():]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return unpadData(plaintext, block.BlockSize()), nil
}

func padData(data []byte, size int) []byte {
	padding := size - len(data)%size
	pad := make([]byte, padding)
	for i := 0; i < padding; i++ {
		pad[i] = byte(padding)
	}
	return append(data, pad...)
}

func unpadData(data []byte, size int) []byte {
	padding := int(data[len(data)-1])
	if padding < 1 || padding > size {
		panic("incorrect padding")
	}
	return data[:len(data)-padding]
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestUser User
	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// In a real application, you'd authenticate the user here
	// For simplicity, let's assume the user exists
	storedUser := User{
		ID:       "user123",
		Username: "exampleuser",
	}

	if storedUser.Username == requestUser.Username {
		token, err := createJWT(&storedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "{\"token\": \"%s\"}", token)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Token not provided", http.StatusUnauthorized)
		return
	}

	tokenString = tokenString[len("Bearer "):] // Remove "Bearer " from the token
	user, err := verifyJWT(tokenString)
	fmt.Print(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Now you can handle the protected request
	encryptedData := []byte("Hello, encrypted message!")
	ciphertext, err := encryptData(encryptedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Ciphertext string `json:"ciphertext"`
	}{
		base64.URLEncoding.EncodeToString(ciphertext),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", protectedHandler)
	fmt.Println("Started..")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
