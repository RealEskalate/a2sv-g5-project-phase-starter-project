package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOAuthConfig(env Env) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     env.OAUTH_CLIENT_ID,
		RedirectURL:  env.OAUTH_REDIRECT_URL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
		ClientSecret: env.OAUTH_CLIENT_SECRET,
	}
}
