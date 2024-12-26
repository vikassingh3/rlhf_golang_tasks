package main

import (
	"container/heap"
	"fmt"
	"time"
)

// Event represents a scheduled event.
type Event struct {
	Name string
	Time time.Time
}

// PriorityQueue implements heap.Interface for managing Event scheduling.
type PriorityQueue []*Event

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Time.Before(pq[j].Time) }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Event)
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority (earliest time).
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	// Initialize the priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add events to the queue
	events := []Event{
		{"Event 1", time.Now().Add(10 * time.Second)},
		{"Event 2", time.Now().Add(5 * time.Second)},
		{"Event 3", time.Now().Add(15 * time.Second)},
	}

	for _, event := range events {
		heap.Push(pq, &event)
	}

	// Process events in priority order
	for pq.Len() > 0 {
		event := heap.Pop(pq).(*Event)
		fmt.Printf("Processing %s scheduled for %v\n", event.Name, event.Time)
	}
}
