// Package jwt provides JWT generation and validation services.
package jwt

import (
  "errors"
  "time"

  "github.com/dgrijalva/jwt-go"
  usermodel "github.com/group13/blog/domain/models/user"
  ijwt "github.com/group13/blog/usecase/common/i_jwt"
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
			 jwt.StandardClaims{
				 ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			},
		}
		} else if tokenType == "refresh" {

		claims = jwt.MapClaims{
			"email":  email, 
			"name": name,
			"role": role,
			 jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
			},
		}
	} else {
		claims = jwt.MapClaims{
			"email":  email, 
			"name": name,
			"role": role,
			 jwt.StandardClaims{
				 ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			},
		}
	}

	log.Println("Generating token with claims:", claims)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwt_secret_key))
	if err != nil {
		log.Println("Error generating token:", err)
		return "", er.NewUnexpected("couldn't generate token")
	}
	log.Println("Generated token:", token)

	return token, nil
}

func (s *Service)Decode(token string) (jwt.MapClaims, error) {
	jwt_secret_key := config.Envs.JWTSecret
	
	parsedToken, err := jwt.ParseWithClaims(token, &Service{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwt_secret_key), nil
	},
	)

	if err != nil {
		return nil, errors.New("wrong Credentails")
	}

	claims, ok := parsedToken.Claims.(*Service)

	if !ok {
		return nil, errors.New("wrong Credentails")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("expired Token")
	}

	return jwt.MapClaims{
		"email":    claims.Email,
		"username": claims.Username,
		"role":     claims.Role,
		"exp":      claims.ExpiresAt,
	}, nil

}
