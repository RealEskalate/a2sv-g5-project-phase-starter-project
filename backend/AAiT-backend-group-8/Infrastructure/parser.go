package infrastructure

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Parse(ctx *gin.Context) (*jwt.MapClaims, error) {
	authHeader := ctx.GetHeader("Authorization")
	authParts := strings.Split(authHeader, " ")

	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, errors.New("invalid header")
	}

	token, err := jwt.ParseWithClaims(authParts[1], jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("123456abed"), nil
	})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, errors.New("invalid header")
}
