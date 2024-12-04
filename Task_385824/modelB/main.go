package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work to be performed
type Task func()

// TaskQueue manages the task queue and worker threads
type TaskQueue struct {
	tasks   chan Task
	workers int
	wg      sync.WaitGroup
}

// NewTaskQueue creates a new TaskQueue with the specified number of workers
func NewTaskQueue(workers int) *TaskQueue {
	return &TaskQueue{
		tasks:   make(chan Task),
		workers: workers,
	}
}

// Start starts the worker threads
func (q *TaskQueue) Start() {
	for i := 0; i < q.workers; i++ {
		q.wg.Add(1)
		go q.worker()
	}
}

// worker is a goroutine that processes tasks from the queue
func (q *TaskQueue) worker() {
	defer q.wg.Done() // Ensure WaitGroup is decremented when worker exits
	for task := range q.tasks {
		task()
	}
}

// Add adds a new task to the queue
func (q *TaskQueue) Add(task Task) {
	q.tasks <- task
}

// Close signals that no more tasks will be added
func (q *TaskQueue) Close() {
	close(q.tasks)
}

// Wait waits for all workers to finish processing tasks
func (q *TaskQueue) Wait() {
	q.wg.Wait()
}

func main() {
	// Create a task queue with 3 workers
	q := NewTaskQueue(3)
	q.Start()

	// Define some sample tasks
	tasks := []Task{
		func() {
			fmt.Println("Processing task 1...")
			time.Sleep(1 * time.Second)
		},
		func() {
			fmt.Println("Processing task 2...")
			time.Sleep(1500 * time.Millisecond)
		},
		func() {
			fmt.Println("Processing task 3...")
			time.Sleep(500 * time.Millisecond)
		},
		func() {
			fmt.Println("Processing task 4...")
			time.Sleep(1 * time.Second)
		},
		func() {
			fmt.Println("Processing task 5...")
			time.Sleep(750 * time.Millisecond)
		},
	}

	// Add tasks to the queue
	for _, task := range tasks {
		q.Add(task)
	}

	// Signal that no more tasks will be added
	q.Close()

	// Wait for all tasks to be completed
	q.Wait()

	fmt.Println("All tasks completed.")
}
