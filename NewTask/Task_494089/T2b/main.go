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
// It returns an error if the event name is empty or there are duplicate registrants.
func (events *Events) AddEvent(name string, registrants []string) error {
	if name == "" {
		return errors.New("event name cannot be empty")
	}

	registrantSet := make(map[string]bool)
	for _, registrant := range registrants {
		if registrantSet[registrant] {
			return errors.New("duplicate registrant found")
		}
		registrantSet[registrant] = true
	}

	*events = append(*events, Event{
		Name:        name,
		Registrants: registrants,
	})
	return nil
}

// RetrieveEvent finds and returns an event by its name.
// It returns an error if the event is not found.
func (events *Events) RetrieveEvent(name string) (Event, error) {
	for _, event := range *events {
		if event.Name == name {
			return event, nil
		}
	}
	return Event{}, errors.New("event not found")
}



func TestAddEvent(t *testing.T) {
	var events Events
	testCases := []struct {
		name        string
		registrants []string
		wantErr     bool
	}{
		{
			name:        "Add a new event",
			registrants: []string{"Alice", "Bob"},
			wantErr:     false,
		},
		{
			name:        "Add an event with duplicate registration",
			registrants: []string{"Alice", "Bob", "Alice"},
			wantErr:     true,
		},
		{
			name:        "Add an event with empty name",
			registrants: []string{"Charlie"},
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := events.AddEvent(tc.name, tc.registrants)
			if (err != nil) != tc.wantErr {
				t.Errorf("AddEvent() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestRetrieveEvent(t *testing.T) {
	var events Events
	// Add some events for retrieval testing
	events.AddEvent("Workshop A", []string{"Alice", "Bob"})
	events.AddEvent("Webinar B", []string{"Charlie", "David"})

	testCases := []struct {
		searchName string
		wantName   string
		wantErr    bool
	}{
		{
			searchName: "Workshop A",
			wantName:   "Workshop A",
			wantErr:    false,
		},
		{
			searchName: "Conference C",
			wantName:   "",
			wantErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.searchName, func(t *testing.T) {
			event, err := events.RetrieveEvent(tc.searchName)
			if (err != nil) != tc.wantErr {
				t.Errorf("RetrieveEvent() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if event.Name != tc.wantName {
				t.Errorf("RetrieveEvent() = %v, want %v", event.Name, tc.wantName)
			}
		})
	}
}
