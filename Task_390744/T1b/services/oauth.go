package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"oauth-client/config"
)

func GetAuthorizationURL(state string) string {
	params := url.Values{
		"client_id":     {config.ClientID},
		"redirect_uri":  {config.RedirectURI},
		"response_type": {"code"},
		"scope":         {config.Scope},
		"state":         {state},
	}
	return fmt.Sprintf("%s?%s", config.AuthURL, params.Encode())
}

func ExchangeCodeForToken(code string) (string, error) {
	params := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {config.ClientID},
		"client_secret": {config.ClientSecret},
		"code":          {code},
		"redirect_uri":  {config.RedirectURI},
	}

	resp, err := http.PostForm(config.TokenURL, params)
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
