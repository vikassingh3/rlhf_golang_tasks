package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	clientID     = "your_client_id"
	clientSecret = "your_client_secret"
	authURL      = "https://accounts.google.com/o/oauth2/auth"
	tokenURL     = "https://accounts.google.com/o/oauth2/token"
	redirectURI  = "http://localhost:8080/callback"
	scope        = "openid email"
)

func getAuthorizationCode() (string, error) {
	params := url.Values{
		"client_id":     {clientID},
		"redirect_uri":  {redirectURI},
		"response_type": {"code"},
		"scope":         {scope},
	}
	authURLStr := fmt.Sprintf("%s?%s", authURL, params.Encode())

	fmt.Println("Go to the following link in your browser:")
	fmt.Println(authURLStr)

	var code string
	fmt.Print("Enter the authorization code: ")
	_, err := fmt.Scanln(&code)
	return code, err
}

func exchangeCodeForToken(code string) (string, error) {
	params := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"code":          {code},
		"redirect_uri":  {redirectURI},
	}

	resp, err := http.PostForm(tokenURL, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse map[string]interface{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("no access token in response: %v", tokenResponse)
	}

	return accessToken, nil
}

func main() {
	code, err := getAuthorizationCode()
	if err != nil {
		fmt.Println("Error obtaining authorization code:", err)
		os.Exit(1)
	}

	accessToken, err := exchangeCodeForToken(code)
	if err != nil {
		fmt.Println("Error exchanging code for token:", err)
		os.Exit(1)
	}

	fmt.Println("Access Token:", accessToken)
}