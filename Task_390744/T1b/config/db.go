package config

const (
	ClientID     = "your_client_id"
	ClientSecret = "your_client_secret"
	AuthURL      = "https://accounts.google.com/o/oauth2/auth"
	TokenURL     = "https://accounts.google.com/o/oauth2/token"
	RedirectURI  = "http://localhost:8080/callback"
	Scope        = "openid email"
)
