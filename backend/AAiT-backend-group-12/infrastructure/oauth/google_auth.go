package google_auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type GoogleUser struct {
	ISS           string `json:"iss"`
	AZP           string `json:"azp"`
	AUD           string `json:"aud"`
	SUB           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Iat           string `json:"iat"`
	Exp           string `json:"exp"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
	Typ           string `json:"typ"`
}

// NewAuth initializes the Google OAuth provider
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

/*
VerifyIdToken verifies the Google ID token by making a request to the Google OAuth2 API and
checks if the issuer is Google, the audience is the client ID, the token has not expired and
the email in the token matches the email in the tokenEmail parameter
*/
func VerifyIdToken(idToken string, tokenEmail string, googleClientID string) error {
	// call to API
	url := "https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var USER GoogleUser

	err = json.NewDecoder(resp.Body).Decode(&USER)
	if err != nil {
		return err
	}

	return VerifyResponseContents(USER, googleClientID, tokenEmail)
}

// VerifyResponseContents verifies the contents of the Google response
func VerifyResponseContents(user GoogleUser, googleClientID string, tokenEmail string) error {
	if user.ISS != "accounts.google.com" && user.ISS != "https://accounts.google.com" {
		return fmt.Errorf("invalid issuer")
	}

	if user.AUD != googleClientID {
		return fmt.Errorf("invalid audience")
	}

	exp, err := strconv.ParseInt(user.Exp, 10, 64)
	if exp == 0 || err != nil {
		return fmt.Errorf("invalid expiration date")
	}

	expiresAt := time.Unix(int64(exp), 0)
	if expiresAt.Before(time.Now()) {
		return fmt.Errorf("token expired")
	}

	if user.Email != tokenEmail {
		return fmt.Errorf("email does not match")
	}

	return nil
}
