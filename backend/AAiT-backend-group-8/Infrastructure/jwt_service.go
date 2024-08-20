package infrastructure

import (
	domain "AAiT-backend-group-8/Domain"
    "errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenService struct {
	SecretKey string
}

func NewTokenService(secretKey string) domain.ITokenService {
	return &TokenService{SecretKey: secretKey}
}

func (ts *TokenService) GenerateToken(email string, id primitive.ObjectID, name string, role string, expiryDuration int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"name":  name,
		"role":  role,
		"exp":   expiryDuration,
	})

	jwtToken, err := token.SignedString([]byte(ts.SecretKey))

	return jwtToken, err
}

func (ts *TokenService) ValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if ok && token.Valid {
		return *claims, nil
	}

	return nil, errors.New("invalid token")
}

