package config

import (
	"blogs/bootstrap"
	"blogs/domain"
	"context"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"time"

	"golang.org/x/oauth2"
)

var googleOauthConfig oauth2.Config

func InitOauth2() error {
	clientID, err := bootstrap.GetEnv("GOOGLE_CLIENT_ID")
	if err != nil {
		return err
	}

	clientSecret, err := bootstrap.GetEnv("GOOGLE_CLIENT_SECRET")
	if err != nil {
		return err
	}

	redirectURL, err := bootstrap.GetEnv("API_BASE")
	if err != nil {
		return err
	}

	googleOauthConfig = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL + "/oauth2/callback/google",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	return nil
}

func GetGoogleLoginURL(state string) string {
	return googleOauthConfig.AuthCodeURL(state)
}

func HandleGoogleCallback(ctx context.Context, code string) (*domain.User, error) {
	// Exchange the authorization code for an access token
	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		return nil, err
	}

	// Create an HTTP client using the token
	client := googleOauthConfig.Client(ctx, token)

	// Fetch the user's profile information
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the user information
	userInfo, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read user info: %v", err)
		return nil, err
	}

	var googleUser struct {
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Email     string `json:"email"`
		Avatar    string `json:"picture"`
	}

	if err := json.Unmarshal(userInfo, &googleUser); err != nil {
		log.Printf("Failed to unmarshal user info: %v", err)
		return nil, err
	}

	// Map Google user info to your User model
	user := &domain.User{
		FirstName:  googleUser.FirstName,
		LastName:   googleUser.LastName,
		Email:      googleUser.Email,
		Avatar:     googleUser.Avatar,
		JoinedDate: time.Now(),
		Role:       "user",
	}

	return user, nil
}

func GenerateState() string {
	symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	state := make([]byte, 64)

	for i := range state {
		state[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(state)
}
