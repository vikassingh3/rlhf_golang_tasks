package main

import (
	"errors"
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
func (events *Events) AddEvent(name string, registrants []string) error {
	if len(registrants) == 0 {
		return errors.New("registrants list must not be empty")
	}
	
	// Check for duplicate registrants
	registrantSet := make(map[string]bool)
	for _, registrant := range registrants {
		if _, found := registrantSet[registrant]; found {
			return errors.New("duplicate registrants found")
		}
		registrantSet[registrant] = true
	}

	*events = append(*events, Event{
		Name:        name,
		Registrants: registrants,
	})
	return nil
}

// RemoveEvent removes an event by its name from the Events slice.
func (events *Events) RemoveEvent(name string) {
	// Find the index of the event to remove
	index := -1
	for i, event := range *events {
		if event.Name == name {
			index = i
			break
		}
	}

	// If event is found, remove it
	if index != -1 {
		// Using append and slice slicing to remove the element
		*events = append((*events)[:index], (*events)[index+1:]...)
	}
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

// Unit Tests
func TestAddEvent(t *testing.T) {
	var events Events

	err := events.AddEvent("Workshop", []string{"Alice", "Bob"})
	if err != nil {
		t.Errorf("AddEvent failed: %v", err)
	}

	err = events.AddEvent("Workshop", []string{"Alice", "Bob"})
	if err == nil {
		t.Errorf("AddEvent should have failed for duplicate registrants")
	}
}

func TestRemoveEvent(t *testing.T) {
	var events Events
	events.AddEvent("Workshop", []string{"Alice", "Bob"})
	events.AddEvent("Webinar", []string{"Charlie"})

	// Remove Workshop event
	events.RemoveEvent("Workshop")
	if len(events) != 1 {
		t.Errorf("RemoveEvent failed: expected 1 event, got %d", len(events))
	}

	// Try to remove non-existing event
	events.RemoveEvent("NonExistent")
	if len(events) != 1 {
		t.Errorf("RemoveEvent failed: expected 1 event, got %d", len(events))
	}
}

func TestRetrieveEvent(t *testing.T) {
	var events Events
	events.AddEvent("Workshop", []string{"Alice", "Bob"})

	// Test retrieving existing event
	event, found := events.RetrieveEvent("Workshop")
	if !found || event.Name != "Workshop" {
		t.Errorf("RetrieveEvent failed, got: %v, found: %v", event, found)
	}

	// Test retrieving non-existing event
	_, found = events.RetrieveEvent("NonExistent")
	if found {
		t.Errorf("RetrieveEvent should not find non-existent event")
	}
}