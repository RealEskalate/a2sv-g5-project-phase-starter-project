package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func InitOauth() (*oauth2.Config, error){
	err := godotenv.Load()
    if err != nil {
        log.Printf("Error loading .env file: %v", err)
        return nil, err
    }
	clientId := os.Getenv("CLIENTID")
	if clientId == ""{
		log.Printf("Error loading client ID: %v", err)
        return nil, err
	}
	clientSecret := os.Getenv("CLIENTSECRET")
	if clientSecret == ""{
		log.Printf("Error loading client secret: %v", err)
        return nil, err
	}
	redirectUrl := os.Getenv("REDIRECTURL")
	if redirectUrl == ""{
		log.Printf("Error loading redirect url: %v", err)
        return nil, err
	}

	googleOauthConfig := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	return &googleOauthConfig, nil
}