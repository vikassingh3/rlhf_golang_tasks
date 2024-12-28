package main  
import (  
    "fmt"
    "sync"
    "time"
)

// Simple Counter Actor
type CounterActor struct {
    count int
    ch    chan int
    wg    *sync.WaitGroup
}

func NewCounterActor(wg *sync.WaitGroup) *CounterActor {
    return &CounterActor{
        ch: make(chan int),
        wg: wg,
    }
}

func (a *CounterActor) Start() {
    defer a.wg.Done()
    for msg := range a.ch {
        a.count += msg
        fmt.Printf("Current count: %d\n", a.count)
        // Simulate some work
        time.Sleep(time.Duration(msg) * time.Millisecond)
    }
}

func (a *CounterActor) Send(msg int) {
    a.ch <- msg
}

func main() {
    var wg sync.WaitGroup

    // Create multiple actors
    numActors := 3
    actors := make([]*CounterActor, numActors)
    for i := 0; i < numActors; i++ {
        actors[i] = NewCounterActor(&wg)
    }

    // Start all actors
    wg.Add(numActors)
    for _, actor := range actors {
        go actor.Start()
    }

    // Send messages to the actors
    messagesToSend := 10
    for _, actor := range actors {
        for i := 0; i < messagesToSend; i++ {
            actor.Send(i + 1)
        }
    }

    // Close channels to signal actors to stop
    for _, actor := range actors {
        close(actor.ch)
    }

    // Wait for all actors to finish
    wg.Wait()
    fmt.Println("All actors finished.")
}