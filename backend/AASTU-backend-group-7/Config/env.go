package Config_test

import (
	"log"
	"os"
)

// Global variable to store the JWT secret
var JwtSecret = []byte("your_jwt_secret")
var MONGO_CONNECTION_STRING string

func Envinit() {
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

}
