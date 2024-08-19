package domain

import (
  "github.com/golang-jwt/jwt"
)

type Claims interface {
  Valid() error
  GetUsername() string
  SetExpiresAt(expiry int64)
}

type LoginClaims struct {
  Username string `json:"username"`
  Role     string `json:"role"`
  jwt.StandardClaims
}

type PasswordResetClaims struct {
  Username string `json:"username"`
  Password string `json:"password"`
  jwt.StandardClaims
}

type RegisterClaims struct {
  User `json:"user"`
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

func (c *LoginClaims) GetUsername() string {
  return c.Username
}

func (c *PasswordResetClaims) GetUsername() string {
  return c.Username
}

func (c *RegisterClaims) GetUsername() string {
  return c.User.Username
}

func (c *LoginClaims) SetExpiresAt(expiry int64) {
  c.StandardClaims.ExpiresAt = expiry
}

func (c *PasswordResetClaims) SetExpiresAt(expiry int64) {
  c.StandardClaims.ExpiresAt = expiry
}

func (c *RegisterClaims) SetExpiresAt(expiry int64) {
  c.StandardClaims.ExpiresAt = expiry
}
