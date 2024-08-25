package entities

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
)

type SignupRequest struct {
	FirstName string `json:"first_name" bson:"first_name" binding:"required,min=3,max=30"`
	LastName  string `json:"last_name" bson:"last_name" binding:"max=30"`
	Email     string `json:"email" bson:"email" binding:"required,email"`
	Password  string `json:"password" bson:"password" binding:"required,min=4,max=30,StrongPassword"`
	Bio       string `json:"bio" bson:"bio"`
}
type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) (*User, error)
	ActivateUser(c context.Context, userID string) error
	IsOwner(c context.Context) (bool, error)
	GetUserById(c context.Context, userId string) (*User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateVerificationToken(user *User, secret string, expiry int) (accessToken string, err error)
	SendVerificationEmail(recipientEmail string, encodedToken string, env *bootstrap.Env) (err error)
}
