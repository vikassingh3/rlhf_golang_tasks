package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	bucketSize  = 10  // Max tokens in the bucket
	fillRate    = 5   // Tokens added per second
	maxRequests = 5   // Maximum concurrent requests allowed
)

type TokenBucket struct {
	bucketSize int
	fillRate   int
	tokens     int
	lastFill   time.Time
	mu         sync.Mutex
}

func NewTokenBucket(size, rate int) *TokenBucket {
	return &TokenBucket{
		bucketSize: size,
		fillRate:   rate,
		tokens:     size,
		lastFill:   time.Now(),
	}
}

func (tb *TokenBucket) ConsumeToken() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()

	// Add new tokens to the bucket
	tokensToAdd := int((now.UnixNano() - tb.lastFill.UnixNano()) / (int64(time.Second) / int64(tb.fillRate)))
	if tokensToAdd > 0 {
		tb.tokens = min(tb.tokens+tokensToAdd, tb.bucketSize)
		tb.lastFill = now
	}

	// Check if there is a token available
	if tb.tokens == 0 {
		return false
	}

	tb.tokens--
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Endpoint struct {
	tokenBucket *TokenBucket
}

func NewEndpoint(size, rate int) *Endpoint {
	return &Endpoint{
		tokenBucket: NewTokenBucket(size, rate),
	}
}

func (e *Endpoint) handleRequest(w http.ResponseWriter, r *http.Request) {
	if !e.tokenBucket.ConsumeToken() {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	response := map[string]string{"message": "Hello, World!"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	endpoint := NewEndpoint(bucketSize, fillRate)
	http.HandleFunc("/api/example", endpoint.handleRequest)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}