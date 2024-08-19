package services


import (
	"github.com/golang-jwt/jwt/v4"
	"errors"
	"time"
)

type IJWT interface {
	GenerateAccessToken(userId string) (string, error)
	GenerateRefreshToken(userId string) (string, error)
	ValidateAccessToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
}

type JWTService struct{
	secretKey string
}

func NewJWTService(secretKey string) IJWT {
	return &JWTService{
		secretKey: secretKey,
	}
}

func (jwtservice *JWTService) validator(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(jwtservice.secretKey), nil
	})

	if err != nil {
		return nil, err
	}
	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return token, err
	}
	return nil, errors.New("invalid token")
}

func (jwtservice *JWTService) GenerateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId" : userId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	} 
	accToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accToken.SignedString([]byte(jwtservice.secretKey))
}

func (jwtservice *JWTService) GenerateRefreshToken(userId string) (string, error){
	claims := jwt.MapClaims{
		"userId":userId,
		"exp": time.Now().Add(time.Hour * 24 *7).Unix(),
	}

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refToken.SignedString([]byte(jwtservice.secretKey))
}	

func (jwtservice *JWTService) ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwtservice.validator(token)
}

func (jwtservice *JWTService) ValidateRefreshToken(token string) (*jwt.Token, error) {
	return jwtservice.validator(token)
}