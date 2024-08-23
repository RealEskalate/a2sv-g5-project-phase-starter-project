package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Config() *oauth2.Config {
	var Client_se = ENV.GOOGLE_CLIENT_SECRET
	var client_id = ENV.GOOGLE_CLIENT_ID
	var (
		GoogleOAuthConfig = &oauth2.Config{
			ClientID:     client_id,
			ClientSecret: Client_se,
			RedirectURL:  "http://localhost:8080/auth/google/callback",
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
			Endpoint:     google.Endpoint,
		}
	)
	return GoogleOAuthConfig
}

var GoogleOAuthConfig = Config()
