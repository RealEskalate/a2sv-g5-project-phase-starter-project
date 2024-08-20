package google_auth

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func NewAuth(clientId string, clientSecret string, maxAgeDays int, callbackUrl string) {
	store := sessions.NewCookieStore([]byte(clientSecret))
	store.MaxAge(86400 * maxAgeDays)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true

	gothic.Store = store

	goth.UseProviders(
		google.New(clientId, clientSecret, callbackUrl),
	)
}
