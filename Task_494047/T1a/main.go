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
	mu       sync.Mutex
	cond     = sync.NewCond(&sync.Mutex{})
)

// Scheduler is the centralized function to process events.
func scheduler() {
	for {
		cond.L.Lock()
		for pq.Len() == 0 {
			cond.Wait()
		}

		mu.Lock()
		nextEvent := heap.Pop(&pq).(*Event)
		mu.Unlock()

		cond.L.Unlock()

		timeUntilNextRun := time.Until(nextEvent.NextRun)
		if timeUntilNextRun > 0 {
			time.Sleep(timeUntilNextRun)
		}

		// Run the event
		nextEvent.Run()

		// Reschedule the event
		nextEvent.NextRun = nextEvent.NextRun.Add(1 * time.Minute)

		mu.Lock()
		heap.Push(&pq, nextEvent)
		mu.Unlock()

		cond.Signal() // Signal the scheduler for new events
	}
}

func scheduleEvent(id string, desc string, nextRun time.Time, runFunc func()) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := eventMap[id]; exists {
		fmt.Printf("Event with ID '%s' already exists. Skipping.\n", id)
		return
	}

	event := &Event{
		ID:          id,
		Description: desc,
		NextRun:     nextRun,
		Run:         runFunc,
	}
	eventMap[id] = event
	heap.Push(&pq, event)

	cond.Signal() // Signal the scheduler for new events
}

func main() {
	heap.Init(&pq)
	go scheduler()

	fmt.Println("Starting the event scheduler.")

	sendReminder := func() {
		fmt.Println("Sending reminder for event.")
	}

	// Schedule events
	scheduleEvent("exampleEvent", "Reminder for important meeting", time.Now().Add(1*time.Second), sendReminder)
	scheduleEvent("exampleEvent2", "Daily check-in", time.Now().Add(2*time.Second), sendReminder)

	// Keep the main function running
	select {}
}
