package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google" // Change to your OAuth2 provider
)

var (
	oauth2Config = &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint, // Change based on the provider
	}
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/callback", callbackHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusFound)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != "state" {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "could not get token", http.StatusInternalServerError)
		return
	}

	// Use the token to make requests to the API
	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/userinfo/v2/me") // Example API for user info
	if err != nil {
		http.Error(w, "could not get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Process the response data (omitted for brevity)
	// ...

	fmt.Fprintln(w, "Login Successful! Access Token:", token.AccessToken)
}
