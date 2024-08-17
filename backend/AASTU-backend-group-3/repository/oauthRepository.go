package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"group3-blogApi/config/db"
	"group3-blogApi/domain"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
)

type OAuthRepository interface {
	GenerateAuthURL() string
	ExchangeCodeForToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (string, error)
}

type oauthRepository struct {
	config *oauth2.Config
}

func NewOAuthRepository(config *oauth2.Config) OAuthRepository {
	return &oauthRepository{config: config}
}

func (r *oauthRepository) GenerateAuthURL() string {
	return r.config.AuthCodeURL("random")
}

func (r *oauthRepository) ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return r.config.Exchange(context.Background(), code)
}


func (r *oauthRepository) GetUserInfo(token *oauth2.Token) (string, error) {
	client := r.config.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return "", fmt.Errorf("failed getting user info: %w", err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body: %w", err)
	}

	// Unmarshal the JSON response into a map
	var userInfo map[string]interface{}
	if err := json.Unmarshal(content, &userInfo); err != nil {
		return "", fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	// Create a new user based on the extracted info
	newUser := domain.User{
		Username:  userInfo["name"].(string),
		Email:     userInfo["email"].(string),
		Role:      "user",
		IsActive:  true,
		
		Image:     userInfo["picture"].(string),
		ActivationToken: "",
		TokenCreatedAt:  time.Now(),
	}
	// cheak if registered
	var user domain.User
	err2 := db.UserCollection.FindOne(context.Background(),  bson.M{"email": newUser.Email}).Decode(&user)
	if err2 != nil  {

		db.UserCollection.InsertOne(context.Background(), newUser)
	}

	return string(content), nil
}


