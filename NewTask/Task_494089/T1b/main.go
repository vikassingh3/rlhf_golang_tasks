package main
import ("fmt")

type Event struct {
    Name string
    Registrations []string
}

func main() {
    // Initialize an empty slice of events
    var events []Event
    
    // Adding events
    event1 := Event{Name: "Tech Conference"}
    events = append(events, event1)
    
    event2 := Event{Name: "Cooking Workshop"}
    events = append(events, event2)
    
    // Adding registrations for events
    events[0].Registrations = append(events[0].Registrations, "Alice", "Bob")
    events[1].Registrations = append(events[1].Registrations, "Charlie", "David", "Eve")
    
    // Retrieving events by name
    fmt.Println("Events:")
    for _, event := range events {
        fmt.Printf("Event: %s\n", event.Name)
        fmt.Println("Registrations:")
        for _, reg := range event.Registrations {
            fmt.Println("  ", reg)
        }
        fmt.Println()
    }
    
    // Retrieving event by index
    fmt.Println("\nEvent at index 1:")
    eventIndex1 := events[1]
    fmt.Printf("Event: %s\n", eventIndex1.Name)
    fmt.Println("Registrations:")
    for _, reg := range eventIndex1.Registrations {  
        fmt.Println("  ", reg)
    }
}