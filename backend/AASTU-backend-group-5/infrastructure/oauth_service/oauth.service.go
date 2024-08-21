package oauthservice

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/RealEskalate/blogpost/domain"
	"golang.org/x/oauth2"
)

type GoogleOAuthService struct {
	GoogleOauthConfig *oauth2.Config
}

func NewGoogleOAuthService(GOC *oauth2.Config) *GoogleOAuthService {
	return &GoogleOAuthService{
		GoogleOauthConfig: GOC,
	}
}

func generateState() string {
	symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	state := make([]byte, 64)

	for i := range state {
		state[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(state)
}

func (GOS *GoogleOAuthService) GetGoogleLoginURL(state string) string{
	return GOS.GoogleOauthConfig.AuthCodeURL(state)
}
func (GOS *GoogleOAuthService) HandleGoogleCallback(code string) (*domain.User, error){
	token, err := GOS.GoogleOauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		return nil, err
	}

	client := GOS.GoogleOauthConfig.Client(context.TODO(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	userInfo, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read user info: %v", err)
		return nil, err
	}

	var googleUser struct {
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Email     string `json:"email"`
	}

	if err := json.Unmarshal(userInfo, &googleUser); err != nil {
		log.Printf("Failed to unmarshal user info: %v", err)
		return nil, err
	}

	user := &domain.User{
		UserName:  googleUser.FirstName + "_" + googleUser.LastName,
		Email:      googleUser.Email,
		IsVerified: true,
		Is_Admin:   false,
	}

	return user, nil
}

func  (GOS *GoogleOAuthService) GetState() domain.State{
	stateString := generateState()

	return domain.State{
		StateID:        stateString,
		ExpiresAT: 		time.Now().Add(time.Minute * 10),
	}
}
