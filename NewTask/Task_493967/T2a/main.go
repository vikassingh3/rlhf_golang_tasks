package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Example: Simple Worker Actor
type WorkerActor struct {
	id   int
	ch   chan int // Channel to receive messages
	done chan bool
	wg   *sync.WaitGroup
}

func NewWorkerActor(id int, wg *sync.WaitGroup) *WorkerActor {
	return &WorkerActor{
		id:   id,
		ch:   make(chan int),
		done: make(chan bool),
		wg:   wg,
	}
}

func (a *WorkerActor) Start() {
	defer a.wg.Done() // Notify WaitGroup when this worker exits

	// This is the main actor loop that processes messages
	for {
		select {
		case msg := <-a.ch:
			// Simulate some work
			time.Sleep(time.Duration(msg) * time.Millisecond)
			fmt.Printf("Worker %d: Completed task with delay %d ms\n", a.id, msg)
		case <-a.done:
			return // Shut down actor when done is sent
		}
	}
}

// Send a message to the actor
func (a *WorkerActor) Send(msg int) {
	a.ch <- msg
}

// Shut down the actor
func (a *WorkerActor) Shutdown() {
	close(a.done)
}

func main() {
	// Create a WaitGroup to keep track of all workers
	var wg sync.WaitGroup

	// Create multiple worker actors
	workers := make([]*WorkerActor, 3)
	for i := 0; i < len(workers); i++ {
		workers[i] = NewWorkerActor(i, &wg)
		wg.Add(1) // Increment the WaitGroup counter for each worker
		go workers[i].Start()
	}

	// Send tasks to each worker
	tasks := []int{100, 200, 300, 400, 500}
	for _, task := range tasks {
		randomWorker := int(rand.Int31n(int32(len(workers))))
		workers[randomWorker].Send(task)
	}

	// Wait for all workers to finish their tasks
	for _, worker := range workers {
		worker.Shutdown()
	}

	wg.Wait() // Ensures all workers are done
	fmt.Println("All workers have completed their tasks")
}
