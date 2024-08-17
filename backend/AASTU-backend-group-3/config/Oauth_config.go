package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "http://localhost:8080/oauth/callback",
		ClientID:     "33361312477-ddpanahl6fj6sk82mav0c2kijcpcgvts.apps.googleusercontent.com",
		ClientSecret:"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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}