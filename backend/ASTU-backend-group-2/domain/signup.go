package domain

import "context"


type SignupRequest struct {
	FirstName  string `json:"first_name" bson:"first_name" binding:"required,min=3,max=30"`
	LastName   string `json:"last_name" bson:"last_name" binding:"max=30"`
	Email      string `json:"email" bson:"email" binding:"required,email"`
	Password   string `json:"password" bson:"password" binding:"required,min=4,max=30,StrongPassword"`
	Bio        string `json:"bio" bson:"bio"`
	ProfileImg string `json:"profile_img" bson:"profile_img"`
}
type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
