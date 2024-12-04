package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gomodule/redigo/redis"
)

const (
	tokenSecret  = "super-secret-key"
	redisAddr    = "localhost:6379"
	blacklistTTL = 30 * time.Minute
)

// Blacklist manages token revocation using Redis.
type Blacklist struct {
	client redis.Conn
}

// NewBlacklist creates a new Blacklist instance.
func NewBlacklist() *Blacklist {
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	return &Blacklist{client: c}
}

// BlockToken adds a token to the blacklist with a TTL.
func (bl *Blacklist) BlockToken(tokenString string) error {
	_, err := bl.client.Do("SETEX", "blacklist:"+tokenString, int(blacklistTTL.Seconds()), "1")
	if err != nil {
		return fmt.Errorf("failed to block token: %w", err)
	}
	return nil
}

// IsBlacklisted checks if a token is in the blacklist.
func (bl *Blacklist) IsBlacklisted(tokenString string) (bool, error) {
	exists, err := redis.Bool(bl.client.Do("EXISTS", "blacklist:"+tokenString))
	if err != nil {
		return false, fmt.Errorf("failed to check blacklist: %w", err)
	}
	return exists, nil
}

func main() {
	blacklist := NewBlacklist()
	defer blacklist.client.Close()

	// Generate a new JWT token
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		log.Fatalf("Error creating token: %v", err)
	}

	fmt.Println("Generated Token:", tokenString)

	// Block the token
	err = blacklist.BlockToken(tokenString)
	if err != nil {
		log.Fatalf("Error blocking token: %v", err)
	}

	// Validate the token
	parsedToken, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		log.Printf("Error parsing token: %v", err)
	} else {
		claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
		if ok && parsedToken.Valid {
			fmt.Println("Token is valid. Claims:", claims)
		} else {
			fmt.Println("Token is not valid or expired.")
		}
	}

	// Check if the token is blacklisted
	isBlacklisted, err := blacklist.IsBlacklisted(tokenString)
	if err != nil {
		log.Fatalf("Error checking blacklist: %v", err)
	}

	if isBlacklisted {
		fmt.Println("Token is blacklisted")
	} else {
		fmt.Println("Token is not blacklisted")
	}
}
