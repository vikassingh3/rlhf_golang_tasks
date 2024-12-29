package main

import (
	"fmt"
	"time"
)

const (
	bucketSize = 5  // Max tokens in the bucket
	fillRate    = 2  // Tokens added per second
)

type TokenBucket struct {
	bucketSize int
	fillRate   int
	tokens     int
	lastFill   time.Time
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

func rateLimitedFuncTokenBucket(bucket *TokenBucket, args ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, arg := range args {
			if bucket.ConsumeToken() {
				ch <- arg * arg
			}
		}
	}()

	return ch
}

func main() {
	bucket := NewTokenBucket(bucketSize, fillRate)

	// Generate results with a rate limit of 2 outputs per second using token bucket
	results := rateLimitedFuncTokenBucket(bucket, 1, 2, 3, 4, 5)
	for result := range results {
		fmt.Println(result)
	}
}



// func rateLimitedFuncFixedInterval(interval time.Duration, args ...int) <-chan int {
// 	ch := make(chan int)

// 	go func() {
// 		defer close(ch)
// 		ticker := time.NewTicker(interval)
// 		defer ticker.Stop()

// 		for _, arg := range args {
// 			<-ticker.C
// 			ch <- arg * arg
// 		}
// 	}()

// 	return ch
// }

// func main() {
// 	// Generate results with a rate limit of 2 outputs per second using fixed interval
// 	results := rateLimitedFuncFixedInterval(time.Second/2, 1, 2, 3, 4, 5)
// 	for result := range results {
// 		fmt.Println(result)
// 	}
// }