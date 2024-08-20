// Package jwt provides JWT generation and validation services.
package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	config "github.com/group13/blog/config"
	usermodel "github.com/group13/blog/domain/models/user"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	er "github.com/group13/blog/domain/errors"
)

// Service implements the ijwt.IService interface for handling JWT operations.
type Service struct {
  secretKey string
  issuer    string
  expTime   time.Duration
  refreshExpTime time.Duration
  jwt.StandardClaims
}

var _ ijwt.Service = &Service{}

// Config holds the configuration for creating a new JWT Service.
type Config struct {
  SecretKey string
  Issuer    string
  ExpTime   time.Duration
  RefreshExpTime time.Duration
}

// New creates a new JWT Service with the given configuration.
func New(config Config) *Service {
  return &Service{
    secretKey: config.SecretKey,
    issuer:    config.Issuer,
    expTime:   config.ExpTime,
	refreshExpTime: config.RefreshExpTime,
  }
}

var _ ijwt.Service = &Service{}

// Config holds the configuration for creating a new JWT Service.
type Config struct {
	SecretKey string
	Issuer    string
	ExpTime   time.Duration
}

// New creates a new JWT Service with the given configuration.
func New(config Config) *Service {
	return &Service{
		secretKey: config.SecretKey,
		issuer:    config.Issuer,
		expTime:   config.ExpTime,
	}
}

// Generate creates a new JWT token for the given user.
func (s *Service) Generate(user *usermodel.User, tokenType string) (string, error) {
	email := user.Email()
	name := user.Username()
	role := user.IsAdmin()
	var claims jwt.Claims

	jwt_secret_key := config.Envs.JWTSecret

	if tokenType == "access" {
		claims = jwt.MapClaims{
			"email":  email, 
			"name": name,
			"role": role,
			"exp": jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			},
		}
		} else if tokenType == "refresh" {

		claims = jwt.MapClaims{
			"email":  email, 
			"name": name,
			"role": role,
			"exp": jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
			},
		}
	} else {
		claims = jwt.MapClaims{
			"email":  email, 
			"name": name,
			"role": role,
			"exp" :jwt.StandardClaims{
				 ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			},
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

// Decode parses and validates a JWT token, returning its claims if valid.
func (s *Service) Decode(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, s.getSigningKey)
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("wrong Credentials")
	}




	return jwt.MapClaims{
		"email":    claims["email"],
		"username": claims["name"],
		"role":     claims["role"],
		"exp":      claims["standardClaims"].(jwt.StandardClaims).ExpiresAt,
	}, nil

}
