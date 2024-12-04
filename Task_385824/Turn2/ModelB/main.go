package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type TaskQueue struct {
	tasks chan Task
	stop  chan struct{}
	wg    sync.WaitGroup
	mc    *sync.Mutex
}

func NewTaskQueue(workers int) *TaskQueue {
	q := &TaskQueue{
		tasks: make(chan Task, workers*10), // Buffered channel to reduce blocking
		stop:  make(chan struct{}),
		mc:    &sync.Mutex{},
	}

	q.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go q.worker()
	}

	return q
}

func (q *TaskQueue) worker() {
	for {
		select {
		case task := <-q.tasks:
			task()
		case <-q.stop:
			q.wg.Done()
			return
		}
	}
}

func (q *TaskQueue) Add(task Task) {
	q.tasks <- task
}

func (q *TaskQueue) Shutdown() {
	close(q.tasks)
	q.wg.Wait()
	close(q.stop)
}

func main() {
	q := NewTaskQueue(3)

	tasks := []Task{
		func() {
			fmt.Println("Processing task 1...")
			time.Sleep(time.Second)
		},
		func() {
			fmt.Println("Processing task 2...")
			time.Sleep(1500 * time.Millisecond)
		},
		// Add more tasks...
	}

	for _, task := range tasks {
		q.Add(task)
	}

	q.Shutdown()
	fmt.Println("All tasks completed.")
}
