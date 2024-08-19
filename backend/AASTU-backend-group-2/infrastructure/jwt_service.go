package infrastructure

import (
	"blog_g2/domain"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// a tokenizer for authentication purpose
func TokenGenerator(id primitive.ObjectID, email string, isadmin bool, isAccessToken bool) (string, error) {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SECRET_KEY := os.Getenv("JWT_SECRET")

	var jwtSecret = []byte(SECRET_KEY)

	var expirationTime time.Duration
	if isAccessToken {
		expirationTime = 15 * time.Minute
	} else {
		expirationTime = 7 * 24 * time.Hour
	}

	var claims domain.JWTClaim

	claims.UserID = id.Hex()
	claims.Email = email
	claims.Isadmin = isadmin
	claims.Exp = time.Now().Add(expirationTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

// a token claimer for extracting the necessary datas
func TokenClaimer(tokenstr string) (*jwt.Token, error) {

	SECRET_KEY := DotEnvLoader("JWT_SECRET")

	log.Println("secretkey: ", SECRET_KEY)

	return jwt.ParseWithClaims(tokenstr, &domain.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
}

func RefreshAccessToken(refreshTokenString string) (string, error) {
	token, err := TokenClaimer(refreshTokenString)
	if err != nil || !token.Valid {
		return "", errors.New("Invalid or expired refresh token")
	}

	claims, ok := token.Claims.(*domain.JWTClaim)
	if !ok || claims.Exp < time.Now().Unix() {
		return "", errors.New("Refresh token expired")
	}

	uid, _ := primitive.ObjectIDFromHex(claims.UserID)

	// Issue a new access token
	newAccessToken, err := TokenGenerator(uid, claims.Email, claims.Isadmin, true)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
