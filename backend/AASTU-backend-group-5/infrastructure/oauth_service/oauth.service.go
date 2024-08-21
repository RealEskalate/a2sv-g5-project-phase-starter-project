package oauthservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
)

type GoogleOAuthService struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func NewGoogleOAuthService(clientID, clientSecret, redirectURL string) *GoogleOAuthService {
	return &GoogleOAuthService{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
	}
}

func (g *GoogleOAuthService) VerifyOAuthToken(provider, token string) (string, error) {
	if provider != "google" {
		return "", errors.New("unsupported provider")
	}

	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", token))
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("invalid OAuth token")
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	oauthID, ok := result["sub"].(string)
	if !ok {
		return "", errors.New("invalid OAuth response")
	}
	return oauthID, nil
}

func (g *GoogleOAuthService) GetUserDataFromProvider(provider, token string) (domain.User, error) {
	if provider != "google" {
		return domain.User{}, errors.New("unsupported provider")
	}

	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token))
	if err != nil || resp.StatusCode != 200 {
		return domain.User{}, errors.New("failed to get user data")
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	user := domain.User{
		OAuthID:       result["id"].(string),
		UserName:      result["name"].(string),
		Email:         result["email"].(string),
		OAuthProvider: "google",
		IsVerified:    true, 
	}
	return user, nil
}