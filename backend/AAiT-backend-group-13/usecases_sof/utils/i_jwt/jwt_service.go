package ijwt

import (
	"github.com/dgrijalva/jwt-go"
	usermodel "github.com/group13/blog/domain/models/user"
)


const (
	Access = "access"
	Refresh = "refresh"
	Reset = "reset"
)

type Services interface {
	Generate(user *usermodel.User, tokenType string )(string,  error )
	Decode(token string)(jwt.MapClaims, error)
	
}

