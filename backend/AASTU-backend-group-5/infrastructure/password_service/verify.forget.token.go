package passwordservice

import (
	"errors"
	"log"
	"os"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func IsValidForgetToken(strtoken string , id string) error {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env", err.Error())
		return err
	}
	var SecretKey = []byte(os.Getenv("SECRETKEY"))
	token, err := jwt.ParseWithClaims(strtoken, &domain.EmailUserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token not valid")
	}

	payload, ok := token.Claims.(*domain.EmailUserClaims)
	if !ok {
		return errors.New("token payload not valid")
	}

	if id != payload.ID.Hex() {
		return errors.New("user-Id not matching the token-id")
	}

	return nil
}
