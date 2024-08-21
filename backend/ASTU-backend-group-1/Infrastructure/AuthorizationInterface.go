package infrastructure

import "github.com/golang-jwt/jwt"

type GeneralAuthorizer interface {
	AUTH(tokenString, secretKey string) jwt.Claims
	AdminAuth(anyClaim any) bool
	UserAuth(anyClaim any) bool
}
