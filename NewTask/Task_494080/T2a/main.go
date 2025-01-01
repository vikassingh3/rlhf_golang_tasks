package main

import (
	"fmt"
	"math/rand"
	"time"
)

// CircuitBreaker defines the circuit breaker logic
type CircuitBreaker struct {
	failCount int
	maxFail   int
	resetTime time.Time
	open      bool
}

func (cb *CircuitBreaker) Reset() {
	cb.failCount = 0
	cb.open = false
}

func (cb *CircuitBreaker) IncrementFailCount() {
	cb.failCount++
}

func (cb *CircuitBreaker) IsOpen() bool {
	return cb.open
}

func (cb *CircuitBreaker) TryCall(call func() error) error {
	if cb.IsOpen() {
		return fmt.Errorf("circuit is open")
	}

	err := call()
	if err != nil {
		cb.IncrementFailCount()
		if cb.failCount >= cb.maxFail {
			cb.open = true
			cb.resetTime = time.Now().Add(time.Second * 30)
		}
		return err
	}

	return nil
}

func (cb *CircuitBreaker) Close() {
	if cb.resetTime.IsZero() {
		return
	}
	if time.Now().After(cb.resetTime) {
		cb.Reset()
	}
}

func externalServiceCall() error {
	if rand.Intn(2) == 0 {
		return fmt.Errorf("external service failed")
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cb := &CircuitBreaker{
		maxFail: 3,
	}

	for i := 0; i < 10; i++ {
		err := cb.TryCall(externalServiceCall)
		cb.Close()
		if err != nil {
			fmt.Printf("Call %d failed: %s\n", i, err)
		} else {
			fmt.Printf("Call %d succeeded\n", i)
		}
		time.Sleep(time.Second)
	}
}