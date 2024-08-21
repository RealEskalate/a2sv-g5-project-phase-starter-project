package auth

import (
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge = 86400 * 30
	prod   = false
)

func NewAuth(env *bootstrap.Env) {

	store := sessions.NewCookieStore([]byte(env.SessionSecret))

	store.MaxAge(MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = prod

	gothic.Store = store

	goth.UseProviders(
		google.New(env.GoogleClientID, env.GoogleClientSecret, env.GoogleUrlCallback),
	)
}
