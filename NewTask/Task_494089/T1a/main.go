package main

import (
	"fmt"
)

// Event represents an event with a name and a list of registrants.
type Event struct {
	Name    string
	Registrants []string
}

// Events is a slice of Event objects.
type Events []Event

// AddEvent adds a new event to the Events slice.
func (events *Events) AddEvent(name string, registrants []string) {
	*events = append(*events, Event{
		Name:        name,
		Registrants: registrants,
	})
}

// RetrieveEvent finds and returns an event by its name.
func (events *Events) RetrieveEvent(name string) (Event, bool) {
	for _, event := range *events {
		if event.Name == name {
			return event, true
		}
	}
	return Event{}, false
}

// Main function to demonstrate usage.
func main() {
	var events Events

	// Add some events with registrants.
	events.AddEvent("Workshop A", []string{"Alice", "Bob"})
	events.AddEvent("Webinar B", []string{"Charlie", "David"})

	// Retrieve and print an event by its name.
	workshopA, found := events.RetrieveEvent("Workshop A")
	if found {
		fmt.Printf("Event: %s, Registrants: %v\n", workshopA.Name, workshopA.Registrants)
	} else {
		fmt.Println("Event not found.")
	}

	// Retrieve and print another event by its name.
	conferenceC, found := events.RetrieveEvent("Conference C")
	if found {
		fmt.Printf("Event: %s, Registrants: %v\n", conferenceC.Name, conferenceC.Registrants)
	} else {
		fmt.Println("Event not found.")
	}
}