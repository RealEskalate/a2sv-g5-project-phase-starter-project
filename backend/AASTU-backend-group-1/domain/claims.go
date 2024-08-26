package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims interface {
	Valid() error
	SetExpiry()
	GetSecretKey() []byte
}

type LoginClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Type     string `json:"type"`
	jwt.StandardClaims
}

type PasswordResetClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type RegisterClaims struct {
	User `json:"username"`
	jwt.StandardClaims
}

func (c *LoginClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func (c *PasswordResetClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func (c *RegisterClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func (c *LoginClaims) SetExpiry() {
	var expiry time.Duration
	if c.Type == "refresh" {
		expiry = time.Hour * 24 * 7
	} else {
		expiry = time.Minute * 15
	}

	c.ExpiresAt = time.Now().Add(expiry).Unix()
}

func (c *PasswordResetClaims) SetExpiry() {
	c.ExpiresAt = time.Now().Add(time.Hour).Unix()
}

func (c *RegisterClaims) SetExpiry() {
	c.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
}

func (c *LoginClaims) GetSecretKey() []byte {
	if c.Type == "refresh" {
		return []byte("my-refresh-secret-key")
	} else {
		return []byte("my-access-secret-key")
	}
}

func (c *PasswordResetClaims) GetSecretKey() []byte {
	return []byte("my-password-reset-secret-key")
}

func (c *RegisterClaims) GetSecretKey() []byte {
	return []byte("my-register-secret-key")
}

func (c *LoginClaims) ToToken() *Token {
	return &Token{
		Username:  c.Username,
		ExpiresAt: c.ExpiresAt,
	}
}
