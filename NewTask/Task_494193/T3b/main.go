package main  
import (  
    "fmt"
    "sync"
    "time"
)

const (
    bucketSize    = 100  // Max tokens in the bucket
    fillRate      = 50   // Tokens added per second
    apiCallDelay = 200 * time.Millisecond  // Simulate API call duration
)

type TokenBucket struct {
    bucketSize int
    fillRate   int
    tokens     int
    lastFill   time.Time
    mux        sync.Mutex
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
    tb.mux.Lock()
    defer tb.mux.Unlock()

    now := time.Now()

    // Add new tokens to the bucket
    tokensToAdd := int((now.UnixNano() - tb.lastFill.UnixNano()) / (int64(time.Second) / int64(tb.fillRate)))
    tb.tokens = min(tb.tokens+tokensToAdd, tb.bucketSize)
    tb.lastFill = now

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

func simulateAPICall(id int) error {
    // Simulate API call latency
    time.Sleep(apiCallDelay)
    fmt.Printf("API Call %d completed.\n", id)
    return nil
}

func rateLimitedAPICalls(bucket *TokenBucket, ids ...int) {
    for _, id := range ids {
        if bucket.ConsumeToken() {
            // Make the API call
            err := simulateAPICall(id)
            if err != nil {
                fmt.Printf("API Call %d failed: %v\n", id, err)
            }
        } else {
            // Rate limit exceeded, handle the error or log it
            fmt.Printf("Rate limit exceeded for API Call %d.\n", id)
        }
    }
}

func main() {
    bucket := NewTokenBucket(bucketSize, fillRate)

    // Simulate making multiple API calls
    start := time.Now()
    rateLimitedAPICalls(bucket, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
    elapsed := time.Since(start)

    fmt.Printf("Total time taken: %s\n", elapsed)
}  