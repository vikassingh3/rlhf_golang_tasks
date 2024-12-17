package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// RateLimiter interface
type RateLimiter interface {
	Limit() (bool, time.Time)
}

// TokenBucket implementation
type TokenBucket struct {
	// Current number of tokens in the bucket
	tokens int
	// Maximum number of tokens the bucket can hold
	capacity int
	// Rate at which tokens are added to the bucket
	refillRate int
	// Last time tokens were added to the bucket
	lastRefill time.Time
	// Mutex for concurrent access
	mutex sync.Mutex
}

// NewTokenBucket creates a new TokenBucket rate limiter
func NewTokenBucket(capacity int, refillRate int) RateLimiter {
	return &TokenBucket{
		capacity:     capacity,
		refillRate:   refillRate,
		lastRefill:   time.Now(),
		tokens:       capacity,
	}
}

// Limit checks if the token bucket allows a request, and returns true if allowed
func (rb *TokenBucket) Limit() (bool, time.Time) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	// Calculate the number of tokens to refill
	now := time.Now()
	elapsed := now.Sub(rb.lastRefill)
	toRefill := int(elapsed.Seconds() * float64(rb.refillRate))

	if toRefill > 0 {
		// Refill tokens but don't exceed capacity
		rb.tokens = min(rb.tokens+toRefill, rb.capacity)
		rb.lastRefill = now
	}

	// If the bucket is empty, return false and when the next refill occurs
	if rb.tokens == 0 {
		return false, rb.lastRefill.Add(time.Second / time.Duration(rb.refillRate))
	}

	// Decrease token count and allow request
	rb.tokens--
	return true, now
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func simulateResourceAccess(limiter RateLimiter) {
	allowed, _ := limiter.Limit()
	if allowed {
		// Simulate work
		time.Sleep(time.Duration(100 * time.Millisecond))
		fmt.Println("Accessing resource...")
	} else {
		fmt.Println("Rate limit exceeded. Access denied.")
	}
}

func resourceHandler(w http.ResponseWriter, r *http.Request, limiter RateLimiter) {
	// Simulate multiple goroutines accessing the resource
	const n = 5
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			simulateResourceAccess(limiter)
			wg.Done()
		}()
	}

	wg.Wait()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Requests processed successfully\n")
}

func main() {
	// Create a rate limiter with 5 tokens per second
	limiter := NewTokenBucket(5, 1)

	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		resourceHandler(w, r, limiter)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}