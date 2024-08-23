package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type IJWT interface {
	GenerateAccessToken(userId, role string) (string, error)
	GenerateRefreshToken(userId, role string) (string, error)
	ValidateAccessToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (string, error)
	GenerateVerificationToken(token string) (string, error)
	ValidateVerificationToken(token string) (string, error)
	GetClaimsFromToken(tokenString string) (jwt.MapClaims, bool)
}

type JWTService struct {
	secretKey string
}

func NewJWTService(secretKey string) IJWT {
	return &JWTService{
		secretKey: secretKey,
	}
}

func (jwtservice *JWTService) validator(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

func (jwtservice *JWTService) GenerateAccessToken(userId, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}
	accToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accToken.SignedString([]byte(jwtservice.secretKey))
}

func (jwtservice *JWTService) GenerateRefreshToken(userId, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refToken.SignedString([]byte(jwtservice.secretKey))
}

func (jwtservice *JWTService) ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwtservice.validator(token)
}

func (jwtservice *JWTService) ValidateRefreshToken(token string) (string, error) {
	Token, err := jwtservice.validator(token)
	if err != nil {
		return "", err
	}

	claims, ok := Token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", err
	}

	return userId, nil
}

func (jwtservice *JWTService) GenerateVerificationToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	accToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accToken.SignedString([]byte(jwtservice.secretKey))
}

func (jwtservice *JWTService) ValidateVerificationToken(token string) (string, error) {
	Token, err := jwtservice.validator(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", jwt.ErrTokenExpired
		}
		return "", err
	}

	claims, ok := Token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", err
	}
	return userId, nil
}

func (jwtservice *JWTService) GetClaimsFromToken(tokenString string) (jwt.MapClaims, bool) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}