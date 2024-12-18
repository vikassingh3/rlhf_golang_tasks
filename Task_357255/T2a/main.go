package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(port, nil)
}
