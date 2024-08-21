package tokenservice

import (
	"log"
	"os"
	"time"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerifyToken struct{}

func (VerifyToken) GenrateToken(id string , email string) (string, error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env", err.Error())
	}
	var SecretKey = []byte(os.Getenv("SECRETKEY"))
	obJID,_ := primitive.ObjectIDFromHex(id)
	itoken := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.EmailUserClaims{
		ID:    obJID,
		Email: email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()},
	})
	token, err := itoken.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
