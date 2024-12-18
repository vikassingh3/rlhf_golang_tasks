package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

// Worker processes a task concurrently
func Worker(g *errgroup.Group, task string, delay time.Duration) error {
	g.Go(func() error {
		// Simulate work with a delay
		time.Sleep(delay)
		fmt.Printf("Task %s completed.\n", task)
		return nil
	})
	return nil
}

// Handler for the root route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Handler for the tasks route
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	g, _ := errgroup.WithContext(r.Context())
	for i := 1; i <= 5; i++ {
		task := fmt.Sprintf("Task %d", i)
		g.Go(func() error {
			return Worker(g, task, time.Duration(i)*time.Second)
		})
	}
	if err := g.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "All tasks completed.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/tasks", TaskHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Server error:", err)
	}
}
