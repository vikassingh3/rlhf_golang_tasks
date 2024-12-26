package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/gofuzz"
)

func FuzzHandler(f *gofuzz.Fzz) {
	// Define your valid parameter names
	paramNames := []string{"userID", "email"}

	for i := 0; f.Next(); i++ {
		var q url.Values
		// Fuzz each parameter name
		for _, name := range paramNames {
			var value string
			f.Fuzz(&value)
			q.Add(name, value)
		}

		// Create the URL with fuzzed query parameters
		u := url.URL{Path: "/api/parameters", RawQuery: q.Encode()}
		req, err := http.NewRequest("GET", u.String(), bytes.NewReader(nil))
		if err != nil {
			panic(err)
		}

		// Your actual handler here, let's use the previous parameterHandler for simplicity
		rw := http.ResponseRecorder{Body: &bytes.Buffer{}}
		parameterHandler(rw, req)

		// Analyze the response for anomalies
		response := rw.Body.String()
		if strings.Contains(response, "Invalid Parameter") {
			fmt.Printf("Fuzz Iteration: %d, Invalid Query: %s\n", i, u.String())
			// Further analysis on the response content could help identify anomalies
		}
	}
}