package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter interface defines the behavior of a rate limiter.
type RateLimiter interface {
	Allow() bool
}

// FixedWindowRateLimiter is an implementation of the RateLimiter interface that limits
// requests to a fixed number within a fixed time window.
type FixedWindowRateLimiter struct {
	mu           sync.Mutex
	window       time.Duration
	maxRequests  int
	requestCount int
	lastRequest  time.Time
}

// NewFixedWindowRateLimiter creates a new FixedWindowRateLimiter with the specified
// window duration and maximum number of requests.
func NewFixedWindowRateLimiter(window time.Duration, maxRequests int) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		window:      window,
		maxRequests: maxRequests,
	}
}

// Allow checks if the rate limit allows a request.
func (l *FixedWindowRateLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Check if the window has elapsed since the last request.
	if time.Since(l.lastRequest) >= l.window {
		l.requestCount = 0
		l.lastRequest = time.Now()
	}

	// Check if the limit has been reached.
	if l.requestCount >= l.maxRequests {
		return false
	}

	// Allow the request and increment the request count.
	l.requestCount++
	return true
}

func main() {
	limiter := NewFixedWindowRateLimiter(time.Second, 5) // Limit to 5 requests per second.
	for i := 0; i < 20; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed", i)
		} else {
			fmt.Println("Request denied", i)
		}
		// Introduce a small delay between requests to simulate real-world usage.
		time.Sleep(200 * time.Millisecond)
	}
}
