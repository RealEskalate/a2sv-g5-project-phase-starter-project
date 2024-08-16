package config

import (
	"log"
	"os"
)

// Global variable to store the JWT secret
var JwtSecret = []byte("your_jwt_secret")
var MONGO_CONNECTION_STRING string

func Envinit() {

	// MONGO_CONNECTION_STRING := os.Getenv("MONGO_CONNECTION_STRING")
	// if uri == "" {
	// 	log.Fatal("Set your 'MONGODB_URI' environment variable. " +
	// 		"See: " +
	// 		"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	// }
	
	// JwtSecretKey := os.Getenv("JWT_SECRETE_KEY")
	// if uri == "" {
	// 	log.Fatal("Set your 'JWT_SECRETE_KEY' environment variable. " +
	// 		"See: " +
	// 		"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	// }

	JwtSecretKey := os.Getenv("JWT_SECRETE_KEY")
	if JwtSecretKey != "" {
		JwtSecret = []byte(JwtSecretKey)
	} else {
		JwtSecret = []byte("JwtSecretKey")
		// log.Fatal("JWT secret key not configured")
	}
	// Read MONGO_CONNECTION_STRING from environment
	MONGO_CONNECTION_STRING = os.Getenv("MONGO_CONNECTION_STRING")
	if MONGO_CONNECTION_STRING == "" {
		MONGO_CONNECTION_STRING = "tst"
		log.Fatal("MONGO_CONNECTION_STRING is not set")
	}

}
