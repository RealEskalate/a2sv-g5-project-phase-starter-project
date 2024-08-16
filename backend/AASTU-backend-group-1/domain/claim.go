package domain

import "github.com/golang-jwt/jwt"

type Claim struct {
  Username string `json:"username"`
  Role     string `json:"role"`
  Type     string `json:"type"`
  jwt.StandardClaims
}
