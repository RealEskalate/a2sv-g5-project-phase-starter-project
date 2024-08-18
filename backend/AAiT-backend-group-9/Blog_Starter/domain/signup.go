package domain

import "context"

type UserSignUp struct {
	Username       string `json:"username" bson:"username" binding:"required" `
	Email          string `json:"email" bson:"email" binding:"required"`
	Password       string `json:"password" bson:"password" binding:"required"`
	Name           string `json:"name" bson:"name"`
	Bio            string `json:"bio" bson:"bio"`
	ContactInfo    ContactInfo `json:"contact_info" bson:"contact_info"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
}

type VerifyEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
	OTP   string `form:"otp" binding:"required"`
}

type ResendOTPRequest struct {
	Email string `form:"email" binding:"required,email"`
}

type SignupUsecase interface {
	CreateUser(c context.Context, user *UserSignUp) error // do validation here
	VerifyEmail(c context.Context, req *VerifyEmailRequest) (*UserResponse, error)
	//TODO:// ResendOTP(c context.Context, req *ResendOTPRequest) error
}