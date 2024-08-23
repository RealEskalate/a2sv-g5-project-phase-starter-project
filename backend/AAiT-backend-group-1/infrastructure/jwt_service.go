package infrastructure

import (
	"context"
	"errors"
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
	VerifySecret  string
	ResetSecret   string
	Collection    *mongo.Collection
}

func NewJWTTokenService(accessSecret, refreshSecret, verifySecret, resetSecret string, collection *mongo.Collection) domain.JwtService {
	return &JWTTokenService{AccessSecret: accessSecret, RefreshSecret: refreshSecret, ResetSecret: resetSecret, Collection: collection, VerifySecret: verifySecret}
}

func (service *JWTTokenService) GenerateAccessTokenWithPayload(user domain.User) (string, domain.Error) {
	claim := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString([]byte(service.AccessSecret))
	if err != nil {
		return "", domain.CustomError{Code: 500, Message: "error generating token"}
	}
	return jwtToken, nil
}

func (service *JWTTokenService) GenerateRefreshTokenWithPayload(user domain.User) (string,  domain.Error) {
	claim := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString([]byte(service.RefreshSecret))
	if err != nil {
		return "", domain.CustomError{Code: 500, Message: "error generating token"}
	}

	return jwtToken, nil
}

func (service *JWTTokenService) GenerateVerificationToken(user domain.User) (string,  domain.Error) {
	claim := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString([]byte(service.VerifySecret))
	if err != nil {
		return "", domain.CustomError{Code: 500, Message: "error generating token"}
	}
	return jwtToken, nil
}

func (service *JWTTokenService) GenerateResetToken(email string) (string,  domain.Error) {
	claim := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString([]byte(service.ResetSecret))
	if err != nil {
		return "", domain.CustomError{Code: 500, Message: "error generating token"}
	}

	return jwtToken, nil
}

func (service *JWTTokenService) ValidateResetToken(token string) (*jwt.Token,  domain.Error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{},  error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(service.ResetSecret), nil
	})
	if errParse != nil {
		return nil, domain.CustomError{Code: 500, Message: "error parsing token"}
	}

	if !parsedToken.Valid {
		return nil, domain.CustomError{Code: 400, Message: "invalid token"}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.CustomError{Code: 400, Message: "invalid token claims"}
	}

	requiredClaims := []string{"email", "exp"}
	for _, claim := range requiredClaims {
		if _, exists := claims[claim]; !exists {
			return nil, domain.CustomError{Code: 400, Message: fmt.Sprintf("missing required claim: %s", claim)}
		}
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, domain.CustomError{Code: 400, Message: "token is expired"}
	}
	return parsedToken, nil
}

func (service *JWTTokenService) ValidateAccessToken(token string) (*jwt.Token,  domain.Error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{},  error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(service.AccessSecret), nil
	})
	if errParse != nil {
		return nil, domain.CustomError{Code: 500, Message: "error parsing token"}
	}
	if !parsedToken.Valid {
		return nil, domain.CustomError{Code: 400, Message: "invalid token"}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.CustomError{Code: 400, Message: "invalid token claims"}
	}

	requiredClaims := []string{"username", "user_id", "role", "iat", "exp"}
	for _, claim := range requiredClaims {
		if _, exists := claims[claim]; !exists {
			return nil, domain.CustomError{Code: 400, Message: fmt.Sprintf("missing required claim: %s", claim)}
		}
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, domain.CustomError{Code: 400, Message: "token is expired"}
	}

	return parsedToken, nil
}

func (service *JWTTokenService) ValidateVerificationToken(token string) (*jwt.Token,  domain.Error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{},  error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(service.VerifySecret), nil
	})
	if errParse != nil {
		return nil, domain.CustomError{Code: 500, Message: "error parsing token"}
	}

	if !parsedToken.Valid {
		return nil, domain.CustomError{Code: 400, Message: "token is invalid"}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.CustomError{Code: 400, Message: "invalid token claims"}
	}

	requiredClaims := []string{"username", "email", "role", "password", "exp"}
	for _, claim := range requiredClaims {
		if _, exists := claims[claim]; !exists {
			return nil, domain.CustomError{Code: 400,  Message: fmt.Sprintf("missing required claim: %s", claim)}
		}
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, domain.CustomError{Code: 400, }
	}

	return parsedToken, nil
}

func (service *JWTTokenService) ValidateRefreshToken(token string) (*jwt.Token,  domain.Error) {
	parsedToken, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{},  error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", t.Header["alg"]))
		}
		return []byte(service.RefreshSecret), nil
	})
	if errParse != nil {
		return nil, domain.CustomError{Code: 500, Message: "error parsing token"}
	}
	if !parsedToken.Valid {
		return nil, domain.CustomError{Code: 400, Message: "invalid token"}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.CustomError{Code: 400, Message: "invalid token claims"}
	}

	requiredClaims := []string{"user_id", "exp"}
	for _, claim := range requiredClaims {
		if _, exists := claims[claim]; !exists {
			return nil, domain.CustomError{Code: 400, Message: fmt.Sprintf("missing required claim: %s", claim)}
		}
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, domain.CustomError{Code: 400, Message: "token is expired"}
	}

	return parsedToken, nil
}

func (service *JWTTokenService) RevokedToken(token string)  domain.Error {
	_, err := service.Collection.DeleteOne(context.TODO(), bson.M{"token": token})
	if err != nil {
		return domain.CustomError{Code: 500, Message: "error revoking token"}
	}

	return nil
}