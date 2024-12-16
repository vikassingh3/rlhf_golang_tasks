package main

import (
	"fmt"
	"net/http"

	"go_packages/localize"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Detect language from HTTP headers
		language := r.Header.Get("Accept-Language")
		if language == "" {
			language = "en" // Default language
		}

		// Print a hello message in the detected language
		fmt.Fprintln(w, localize.GetText(language, "hello"))
	})

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}