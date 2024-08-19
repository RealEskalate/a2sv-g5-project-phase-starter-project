package jwt

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/group13/blog/config"
	"github.com/group13/blog/domain/errors"
	usermodel "github.com/group13/blog/domain/models/user"
	ijwt "github.com/group13/blog/usecases_sof/utils/i_jwt"
)

type Service struct {
	Email    string
	Username string
	Role     bool
	jwt.StandardClaims
}



var _ ijwt.Services = &Service{}

func (s *Service) Generate(user *usermodel.User, tokenType string) (string, error) {
	email := user.Email()
	name := user.Username()
	role := user.IsAdmin()
	var claims jwt.Claims

	jwt_secret_key := config.Envs.JWTSecret

	if tokenType == "access" {

		claims = &Service{
			email, name, role, jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			},
		}
	} else if tokenType == "refresh" {

		claims = &Service{
			email, name, role, jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
			},
		}
	} else {
		claims = &Service{
			email, name, role, jwt.StandardClaims{
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
