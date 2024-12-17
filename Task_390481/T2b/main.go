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
    tokens     int
    capacity    int
    refillRate   int
    lastRefill   time.Time
    mutex      sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) RateLimiter {  
    return &TokenBucket{
        capacity:    capacity,
        refillRate:   refillRate,
        lastRefill:   time.Now(),
        tokens:      capacity,
    }
}

func (rb *TokenBucket) Limit() (bool, time.Time) {  
    rb.mutex.Lock()
    defer rb.mutex.Unlock()

    elapsed := time.Since(rb.lastRefill)
    toRefill := int(elapsed.Seconds() * float64(rb.refillRate))

    if toRefill > 0 {
        rb.tokens = min(rb.tokens+toRefill, rb.capacity)
        rb.lastRefill = time.Now()
    }

    if rb.tokens == 0 {
        return false, rb.lastRefill.Add(time.Second / time.Duration(rb.refillRate))
    }

    rb.tokens--
    return true, time.Now()
}

func min(a, b int) int {  
    if a < b {
        return a
    }
    return b
}

// Sample protected resource handler
func protectedResourceHandler(limiter RateLimiter, w http.ResponseWriter, r *http.Request) {
    allowed, nextRefill := limiter.Limit()
    if !allowed {
        http.Error(w, fmt.Sprintf("Request denied; next refill at %v", nextRefill), http.StatusTooManyRequests)
        return
    }

    // Simulate some work being done by the protected resource
    time.Sleep(time.Second)

    fmt.Fprintln(w, "Welcome to the protected resource!")
}

func main() {
    // Create a rate limiter with 5 tokens per second
    limiter := NewTokenBucket(5, 1)

    mux := http.NewServeMux()
    mux.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
        protectedResourceHandler(limiter, w, r)
    })

    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    fmt.Println("Server running on port :8080")
    log.Fatal(server.ListenAndServe())
}