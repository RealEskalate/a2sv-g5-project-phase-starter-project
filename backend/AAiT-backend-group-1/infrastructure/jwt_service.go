package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JWTTokenService struct {
	AccessSecret  string
	RefreshSecret string
	Collection    *mongo.Collection
}

func NewJWTTokenService(accessSecret, refreshSecret string, collection *mongo.Collection) domain.JwtService {
	return &JWTTokenService{ AccessSecret: accessSecret, RefreshSecret: refreshSecret, Collection: collection }
}

func (service *JWTTokenService) GenerateAccessTokenWithPayload(user domain.User) (string, error) {
	claim := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString(service.AccessSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (service *JWTTokenService) GenerateRefreshTokenWithPayload(user domain.User) (string, error) {
	claim := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString(service.RefreshSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (service *JWTTokenService) GenerateVerificationToken(user domain.User) (string, error) {
	claim := jwt.MapClaims{
		"username": user.ID,
		"email":    user.Email,
		"password": user.Password,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString(service.RefreshSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (service *JWTTokenService) GenerateResetToken(email string) (string, error) {
	claim := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString(service.RefreshSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (service *JWTTokenService) ValidateResetToken(token string) (*jwt.Token, error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return service.AccessSecret, nil
	})
	if errParse != nil {
		return nil, errParse
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	requiredClaims := []string{"email", "exp"}
	for _, claim := range requiredClaims {
		if _, exists := claims[claim]; !exists {
			return nil, fmt.Errorf("missing required claim: %s", claim)
		}
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, fmt.Errorf("token has expired")
	}
	return parsedToken, nil
}

func (service *JWTTokenService) ValidateAccessToken(token string) (*jwt.Token, error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return service.AccessSecret, nil
	})
	if errParse != nil {
		return nil, errParse
	}
	if !parsedToken.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	return parsedToken, nil
}

func (service *JWTTokenService) ValidateVerificationToken(token string) (*jwt.Token, error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return service.AccessSecret, nil
	})
	if errParse != nil {
		return nil, errParse
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	requiredClaims := []string{"username", "email", "role", "password", "exp"}
	for _, claim := range requiredClaims {
		if _, exists := claims[claim]; !exists {
			return nil, fmt.Errorf("missing required claim: %s", claim)
		}
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, fmt.Errorf("token has expired")
	}

	return parsedToken, nil
}

func (service *JWTTokenService) ValidateRefreshToken(token string) (*jwt.Token, error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return service.AccessSecret, nil
	})
	if errParse != nil {
		return nil, errParse
	}
	if !parsedToken.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	return parsedToken, nil
}

func (service *JWTTokenService) RevokedToken(token string) error {
	_, err := service.Collection.DeleteOne(context.TODO(), bson.M{"token": token})
	return err
}
