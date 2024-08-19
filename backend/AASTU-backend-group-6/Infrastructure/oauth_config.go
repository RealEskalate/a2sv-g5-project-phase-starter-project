package infrastructure

import (
	domain "blogs/Domain"
	"errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


type Oauth struct {
    config *Config
}

func NewOauthConfig(config *Config) domain.OauthConfig {
    return &Oauth{config: config}
}

func (o *Oauth) InitialConfig() (*oauth2.Config, error) {
	if o.config.ClientID == "" || o.config.ClientSecret == "" || o.config.RedirectURL == "" {
		return nil, errors.New("missing OAuth configuration details")
	}

	return &oauth2.Config{
		ClientID:     o.config.ClientID,
		ClientSecret: o.config.ClientSecret,
		RedirectURL:  o.config.RedirectURL,
		Scopes:       []string{"email", "profile", "openid"},
		Endpoint:     google.Endpoint,
	}, nil
}

