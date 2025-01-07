package handlers

import (
	"fmt"
	"net/http"
	"oauth-client/services"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	if code == "" {
		http.Error(w, "Authorization code missing", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received code: %s, state: %s\n", code, state)

	// Exchange code for access token
	accessToken, err := services.ExchangeCodeForToken(code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the access token
	fmt.Fprintf(w, "Access Token: %s", accessToken)
}
