package Config

import (
	"log"
	"os"
)

var Port = ":8080"
var BASE_URL = "http://localhost" + Port

// Global variable to store the Env variables
var JwtSecret = []byte("your_jwt_secret")
var MONGO_CONNECTION_STRING string
var Mail_TRAP_API_KEY string
var GROQ_API_KEY string
var GOOGLE_KEY string
var GOOGLE_SECRET string
var Google_Callback string
var Cloud_api_key string
var Cloud_api_secret string

func Envinit() {

	GOOGLE_KEY = os.Getenv("GOOGLE_KEY")
	if GOOGLE_KEY == "" {
		log.Fatal("GOOGLE_KEY is not set")
	}
	GOOGLE_SECRET = os.Getenv("GOOGLE_SECRET")
	if GOOGLE_SECRET == "" {
		log.Fatal("GOOGLE_SECRET is not set")
	}
	Google_Callback = os.Getenv("Google_Callback")
	if Google_Callback == "" {
		log.Fatal("Google_Callback is not set")
	}

	GROQKEY := os.Getenv("GROQ_API_KEY")
	if GROQKEY == "" {
		log.Fatal("GROQ_API_KEY is not set")
	} else {
		GROQ_API_KEY = GROQKEY
	}
	JwtSecretKey := os.Getenv("JWT_SECRETE_KEY")
	if JwtSecretKey != "" {
		JwtSecret = []byte(JwtSecretKey)
	} else {
		JwtSecret = []byte("JwtSecretKey")
		log.Fatal("JWT secret key not configured")
	}
	// Read MONGO_CONNECTION_STRING from environment
	MONGO_CONNECTION_STRING = os.Getenv("MONGO_CONNECTION_STRING")
	if MONGO_CONNECTION_STRING == "" {
		MONGO_CONNECTION_STRING = "tst"
		log.Fatal("MONGO_CONNECTION_STRING is not set")
	}
	// Read Mail_TRAP_API_KEY from environment
	Mail_TRAP_API_KEY = os.Getenv("Mail_TRAP_API_KEY")
	if Mail_TRAP_API_KEY == "" {
		log.Fatal("Mail_TRAP_API_KEY is not set")
	}
	Cloud_api_key = os.Getenv("Cloud_api_key")
	if Mail_TRAP_API_KEY == "" {
		log.Fatal("Cloud_api_key is not set")
	}
	Cloud_api_secret = os.Getenv("Cloud_api_secret")
	if Mail_TRAP_API_KEY == "" {
		log.Fatal("Cloud_api_secret is not set")
	}

}
