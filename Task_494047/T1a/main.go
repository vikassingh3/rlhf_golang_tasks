package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Event represents a scheduled task.
type Event struct {
	ID          string
	Description string
	NextRun     time.Time
	Run         func()
	index       int // Required for heap.Interface
}

// PriorityQueue implements heap.Interface for a priority queue.
type PriorityQueue []*Event

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].NextRun.Before(pq[j].NextRun)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	event := x.(*Event)
	event.index = len(*pq)
	*pq = append(*pq, event)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	event := old[n-1]
	event.index = -1 // For safety
	*pq = old[0 : n-1]
	return event
}

var (
	pq       PriorityQueue
	eventMap = make(map[string]*Event)
	wg       sync.WaitGroup
)

func scheduleEvent(id string, desc string, nextRun time.Time, runFunc func()) {
	event := &Event{
		ID:          id,
		Description: desc,
		NextRun:     nextRun,
		Run:         runFunc,
	}
	eventMap[id] = event
	heap.Push(&pq, event)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if pq.Len() == 0 {
				return
			}
			nextEvent := heap.Pop(&pq).(*Event)
			time.Sleep(time.Until(nextEvent.NextRun))
			nextEvent.Run()
			nextEvent.NextRun = nextEvent.NextRun.Add(1 * time.Minute) // Reschedule
			heap.Push(&pq, nextEvent)
		}
	}()
}

func main() {
	heap.Init(&pq)

	fmt.Println("Starting the event scheduler.")

	sendReminder := func() {
		fmt.Println("Sending reminder for event.")
	}

	// Schedule events
	scheduleEvent("exampleEvent", "Reminder for important meeting", time.Now().Add(10*time.Second), sendReminder)
	scheduleEvent("exampleEvent2", "Daily check-in", time.Now().Add(15*time.Second), sendReminder)

	// Wait for the goroutines
	wg.Wait()
}
