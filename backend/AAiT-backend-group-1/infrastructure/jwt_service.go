package infrastructure

import (
	"errors"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService(secretKey, issuer string) domain.JwtService {
	return &jwtService{secretKey: secretKey, issuer: issuer}
}

func (j *jwtService) GenerateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iss"] = j.issuer 

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(j.secretKey), nil
	})
}