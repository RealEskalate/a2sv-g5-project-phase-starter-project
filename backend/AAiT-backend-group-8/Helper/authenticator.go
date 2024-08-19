package Helper

import (
	"AAiT-backend-group-8/Domain"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authenticate(ctx *gin.Context, admin bool, blog *Domain.Blog) error {
	authHeader := ctx.GetHeader("Authorization")
	authParts := strings.Split(authHeader, " ")

	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return errors.New("invalid header")
	}

	token, err := jwt.ParseWithClaims(authParts[1], jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return "secret_key", nil
	})

	if err != nil {
		return errors.New("invalid header")
	}

	if admin {
		if token.Claims.(jwt.MapClaims)["role"].(string) != "admin" {
			if blog.AuthorID.Hex() != token.Claims.(jwt.MapClaims)["id"].(string) {
				return errors.New("invalid token")
			}
		}
	} else if blog.AuthorID.Hex() != token.Claims.(jwt.MapClaims)["id"].(string) {
		return errors.New("you are not the author")
	}

	return nil
}
