package domain

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomClaims) Valid() error {
	panic("unimplemented")
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomRefreshClaims) Valid() error {
	panic("unimplemented")
}
