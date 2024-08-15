package domain

import "github.com/golang-jwt/jwt/v5"

// Claims struct that will be encoded to a JWT.
// This can contain any information you want to include in your JWT.
type Claims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role	 string `json:"role"`
	jwt.RegisteredClaims
}

// Token struct that will be returned to the user
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
