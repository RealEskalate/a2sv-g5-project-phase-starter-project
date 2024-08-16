package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// a tokenizer for authentication purpose
func TokenGenerator(id primitive.ObjectID, email string, isadmin bool) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SECRET_KEY := os.Getenv("JWT_SECRET")

	var jwtSecret = []byte(SECRET_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":     id.Hex(),
		"email":   email,
		"isadmin": isadmin,
	})

	return token.SignedString(jwtSecret)
}

// a token claimer for extracting the necessary datas
func TokenClaimer(tokenstr string) (*jwt.Token, error) {

	SECRET_KEY := DotEnvLoader("JWT_SECRET")

	log.Println("secretkey: ", SECRET_KEY)

	return jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})
}
