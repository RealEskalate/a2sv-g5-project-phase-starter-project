package domain

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID    string `json:"id" bson:"_id"`
	Username       string             `json:"username"`
	Role           string             `json:"role"`
	jwt.StandardClaims
}

type User struct {
    ID            string    `json:"id" ` //bson:"_id,omitempty
    Name          string    `json:"name" validate:"required,min=2,max=100"`
    Email         string    `json:"email" validate:"required,email"`
    Password      string    `json:"-"`
    Role          string    `json:"role"` 
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Profile struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"user_id"`
    Bio       string    `json:"bio"`
    AvatarURL string    `json:"avatar_url"`
}
