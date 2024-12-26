package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Event represents a scheduled task
type Event struct {
	ID          string
	Description string
	NextRun     time.Time
	Run         func()
	index       int // Required for heap.Interface
}

// PriorityQueue implements heap.Interface for a priority queue
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
	pq        PriorityQueue
	eventMap  = make(map[string]*Event)
	workerCount int
	wg        sync.WaitGroup
	taskQueue chan *Event
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
}

func worker(id int) {
	defer wg.Done()
	for event := range taskQueue {
		if event == nil {
			break
		}
		event.Run()
	}
	fmt.Printf("Worker %d exiting\n", id)
}

func main() {
	heap.Init(&pq)
	workerCount = 4 // Number of worker goroutines
	taskQueue = make(chan *Event, 1000) // Buffered channel to hold tasks

	// Start workers
	for i := 0; i < workerCount; i++ {
		go worker(i)
	}

	fmt.Println("Starting the event scheduler.")

	sendReminder := func() {
		fmt.Println("Sending reminder for event.")
	}

	// Schedule events
	scheduleEvent("exampleEvent1", "Reminder for meeting", time.Now().Add(1*time.Second), sendReminder)
	scheduleEvent("exampleEvent2", "Daily check-in", time.Now().Add(2*time.Second), sendReminder)
	// Schedule many more events
	for i := 3; i <= 10; i++ {
		eventID := fmt.Sprintf("exampleEvent%d", i)
		scheduleEvent(eventID, fmt.Sprintf("Reminder %d", i), time.Now().Add(time.Duration(i)*time.Second), sendReminder)
	}

	// Process events in the priority queue
	go func() {
		for pq.Len() > 0 {
			nextEvent := heap.Pop(&pq).(*Event)
			// Send event to the task queue for workers to process
			taskQueue <- nextEvent
			time.Sleep(time.Until(nextEvent.NextRun)) // Sleep until next event time
			// Reschedule the event (example: adding 5 minutes)
			nextEvent.NextRun = nextEvent.NextRun.Add(5 * time.Minute)
			heap.Push(&pq, nextEvent) // Push the event back into the priority queue
		}
		close(taskQueue) // Close the taskQueue to signal workers to stop
	}()

	// Wait for workers to finish
	wg.Wait()
}
