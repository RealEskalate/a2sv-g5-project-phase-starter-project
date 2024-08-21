package usecases

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Config() *oauth2.Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var Client_se = os.Getenv("GOOGLE_CLIENT_SECRET")
	var client_id = os.Getenv("GOOGLE_CLIENT_ID")
	var (
		GoogleOAuthConfig = &oauth2.Config{
			ClientID:     client_id,
			ClientSecret: Client_se,
			RedirectURL:  "http://localhost:8080/auth/google/callback",
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
			Endpoint:     google.Endpoint,
		}
	)
	return GoogleOAuthConfig
}

var GoogleOAuthConfig = Config()
