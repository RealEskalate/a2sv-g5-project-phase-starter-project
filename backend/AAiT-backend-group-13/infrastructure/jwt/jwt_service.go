// Package jwt provides JWT generation and validation services.
package jwt

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
)

// Service implements the ijwt.IService interface for handling JWT operations.
type Service struct {
	secretKey      string
	issuer         string
	expTime        time.Duration
	refreshExpTime time.Duration
}

// Ensure Service implements the ijwt.Service interface.
var _ ijwt.Service = &Service{}

// Config holds the configuration for creating a new JWT Service.
type Config struct {
	SecretKey      string
	Issuer         string
	ExpTime        time.Duration
	RefreshExpTime time.Duration
}

// New creates a new JWT Service with the given configuration.
func New(config Config) *Service {
	return &Service{
		secretKey:      config.SecretKey,
		issuer:         config.Issuer,
		expTime:        config.ExpTime,
		refreshExpTime: config.RefreshExpTime,
	}
}

// Generate creates a new JWT token based on the provided user and token type.
func (s *Service) Generate(user *models.User, tokenType string) (string, error) {
	email := user.Email()
	name := user.Username()
	isAdmin := user.IsAdmin()

	var expTime time.Duration

	switch tokenType {
	case ijwt.Access:
		expTime = s.expTime
	case ijwt.Refresh:
		expTime = s.refreshExpTime
	default:
		expTime = time.Minute * 15
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expTime).Unix(),
		Issuer:    s.issuer,
	}

	tokenClaims := jwt.MapClaims{
		"email":    email,
		"name":     name,
		"is_admin": isAdmin,
		"exp":      claims.ExpiresAt,
		"issuer":   claims.Issuer,
	}
	if tokenType == ijwt.Reset {
		tokenClaims["is_for_reset"] = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Println("Error generating token:", err)
		return "", er.NewUnexpected("couldn't generate token")
	}

	log.Println("Generated token:", signedToken)
	return signedToken, nil
}

// Decode parses and validates the JWT token and returns its claims.
func (s *Service) Decode(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid token: " + err.Error())
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return *claims, nil
}
