package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket represents a token bucket limiter.
type TokenBucket struct {
	tokens  int
	rate    int // Tokens added per second
	burst   int // Maximum tokens to hold
	m       sync.Mutex
	cond    *sync.Cond
}

// NewTokenBucket creates a new token bucket with the given rate and burst.
func NewTokenBucket(rate, burst int) *TokenBucket {
	tb := &TokenBucket{
		rate:  rate,
		burst: burst,
	}
	tb.cond = sync.NewCond(&tb.m)
	return tb
}

// Take takes a token from the bucket, blocking if necessary until a token is available.
func (tb *TokenBucket) Take() error {
	tb.m.Lock()
	defer tb.m.Unlock()

	// Wait for a token to become available.
	for tb.tokens == 0 {
		tb.cond.Wait()
	}

	tb.tokens--
	return nil
}

// refill refills the token bucket at the specified rate.
func (tb *TokenBucket) refill() {
	ticker := time.NewTicker(time.Second / time.Duration(tb.rate))
	defer ticker.Stop()

	for range ticker.C {
		tb.m.Lock()
		if tb.tokens < tb.burst {
			tb.tokens++
			tb.cond.Broadcast() // Notify waiting goroutines when tokens are added.
		}
		tb.m.Unlock()
	}
}

func main() {
	// Create a token bucket with a rate of 2 tokens per second and a burst of 5 tokens.
	tb := NewTokenBucket(2, 5)

	// Start the refill goroutine.
	go tb.refill()

	for i := 1; i <= 10; i++ {
		err := tb.Take() // Block until a token is available
		if err != nil {
			fmt.Printf("Failed to take token: %v\n", err)
			continue
		}
		fmt.Println("Sending request", i)
		time.Sleep(time.Second) // Simulate some work
	}
}


// package main
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	rate := 2 // Requests per second
// 	duration := time.Second / time.Duration(rate)

// 	for i := 1; i <= 10; i++ {
// 		fmt.Println("Sending request", i)
// 		time.Sleep(duration) // Wait for the fixed interval
// 	}
// }
