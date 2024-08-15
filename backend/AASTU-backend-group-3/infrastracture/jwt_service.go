package infrastracture

import (
	"group3-blogApi/config"
	"group3-blogApi/domain"
	"fmt"

	"github.com/golang-jwt/jwt"
)


func GenerateToken(user domain.User) (string, error) {
	var JwtSecret = []byte(config.EnvConfigs.JwtSecret)


	claims := domain.JwtCustomClaims{
		Authorized: true,
		UserID:     user.ID.Hex(),
		Role:       user.Role,
		Username:   user.Username,
	}

	fmt.Println(JwtSecret, "GErator ******************")


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}


