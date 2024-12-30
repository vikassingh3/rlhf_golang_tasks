package main

import (
	"testing"
)

// Event represents an event with a name and a list of registrants.
type Event struct {
	Name        string
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

// VerifyDuplicateRegistrants checks for duplicate registrations in the provided event.
func VerifyDuplicateRegistrants(event Event) bool {
	registrantSet := make(map[string]bool)
	for _, registrant := range event.Registrants {
		if _, found := registrantSet[registrant]; found {
			return true
		}
		registrantSet[registrant] = true
	}
	return false
}

// CompareEvents compares two events for equality, including their registrants.
func CompareEvents(e1, e2 Event) bool {
	if e1.Name != e2.Name {
		return false
	}
	if len(e1.Registrants) != len(e2.Registrants) {
		return false
	}
	for i, registrant := range e1.Registrants {
		if registrant != e2.Registrants[i] {
			return false
		}
	}
	return true
}

func TestAddEvent(t *testing.T) {
	var events Events
	events.AddEvent("Workshop A", []string{"Alice", "Bob"})
	events.AddEvent("Webinar B", []string{"Charlie", "David"})

	expectedEvent1 := Event{Name: "Workshop A", Registrants: []string{"Alice", "Bob"}}
	expectedEvent2 := Event{Name: "Webinar B", Registrants: []string{"Charlie", "David"}}

	if len(events) != 2 {
		t.Errorf("Expected 2 events, got %d", len(events))
	}

	if !CompareEvents(events[0], expectedEvent1) {
		t.Errorf("First event mismatch. Expected %v, got %v", expectedEvent1, events[0])
	}

	if !CompareEvents(events[1], expectedEvent2) {
		t.Errorf("Second event mismatch. Expected %v, got %v", expectedEvent2, events[1])
	}
}

func TestRetrieveEvent(t *testing.T) {
	var events Events
	events.AddEvent("Workshop A", []string{"Alice", "Bob"})
	events.AddEvent("Webinar B", []string{"Charlie", "David"})

	expectedEvent1 := Event{Name: "Workshop A", Registrants: []string{"Alice", "Bob"}}
	retrievedEvent1, found := events.RetrieveEvent("Workshop A")
	if !found {
		t.Error("Event not found when it should be")
	} else if !CompareEvents(retrievedEvent1, expectedEvent1) {
		t.Errorf("Retrieved event mismatch. Expected %v, got %v", expectedEvent1, retrievedEvent1)
	}

	expectedEvent2 := Event{Name: "Webinar B", Registrants: []string{"Charlie", "David"}}
	retrievedEvent2, found := events.RetrieveEvent("Webinar B")
	if !found {
		t.Error("Event not found when it should be")
	} else if !CompareEvents(retrievedEvent2, expectedEvent2) {
		t.Errorf("Retrieved event mismatch. Expected %v, got %v", expectedEvent2, retrievedEvent2)
	}

	_, found = events.RetrieveEvent("Conference C")
	if found {
		t.Error("Event found when it should not be")
	}
}

func TestVerifyDuplicateRegistrants(t *testing.T) {
	var event Event
	event.Name = "Workshop D"
	event.Registrants = []string{"Eve", "Fiona", "Eve"}

	if !VerifyDuplicateRegistrants(event) {
		t.Error("Duplicate registrant not detected")
	}

	event.Registrants = []string{"Grace", "Heidi"}
	if VerifyDuplicateRegistrants(event) {
		t.Error("Duplicate registrant detected when there should be none")
	}
}

