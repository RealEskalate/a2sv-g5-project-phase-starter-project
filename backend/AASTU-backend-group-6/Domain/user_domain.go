package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   primitive.ObjectID   `bson:"_id,omitempity" json:"id" `
	Full_Name            string               `bson:"full_name"  json:"full_name"`
	Email string `bson:"email" validate:"required,email" json:"email"`
    Username             string               `bson:"username" json:"username" validate:"required"`
    Password             string               `bson:"password" json:"password" validate:"required"`
    Profile_image_url    string               `bson:"profile_image" json:"profile_image"`
    GoogleID             string               `bson:"googleId" json:"googleId"`
    PostsID              []primitive.ObjectID `bson:"posts_id" json:"posts_id"`
    Contact              string               `bson:"contact" json:"contact"`
    Bio                  string               `bson:"bio" json:"bio"`
    Role                 string               `bson:"role" json:"role" validate:"required"`
    CommentsID           []primitive.ObjectID `bson:"comments_id" json:"comments_id"`
    LikedPostsID         []primitive.ObjectID `bson:"liked_posts_id" json:"liked_posts_id"`
    DisLikePostsID       []primitive.ObjectID `bson:"disliked_posts_id" json:"disliked_posts_id"`
    ResetPasswordToken   string               `bson:"reset_password_token" json:"reset_password_token"`
    ResetPasswordExpires time.Time            `bson:"reset_password_expires" json:"reset_password_expires"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	FindUserByEmail(ctx context.Context, email string) (User, error)
	FindUserByUsername(ctx context.Context, username string) (User, error)
	FindUserByID(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, id string) error
	ForgotPassword(ctx context.Context, email string, token string) error
	AllUsers(c context.Context) ([]User, error)
	PromoteandDemoteUser(c context.Context, id string, role string) error
}
type Email struct {
	Email string `json:"email"`
}

type SignupRepository interface {
	Create(c context.Context, user User) (User, error)
	FindUserByEmail(c context.Context, email string) (User, error)
	SetOTP(c context.Context, email string, otp string) error
	VerifyUser(c context.Context, user User) (User, error)
	SetResetToken(c context.Context, email ForgotPasswordRequest, token string, expiration time.Time) (User, error)
	FindUserByResetToken(c context.Context, token string) (User, error)
	UpdateUser(c context.Context, user User) (User, error)
}

type SignupUseCase interface {
	Create(c context.Context, user User) interface{}
	VerifyOTP(c context.Context, otp OtpToken) interface{}
	ForgotPassword(c context.Context, email ForgotPasswordRequest) interface{}
	ResetPassword(c context.Context, password ResetPasswordRequest, token string) interface{}
	HandleUnverifiedUser(c context.Context, email Email) interface{}
	// ResendToken(c context.Context , email User) interface{}
}

type UserUseCase interface {
	// FindUserByID(ctx context.Context, id string) interface{}
	UpdateUser(ctx context.Context, user UserUpdateRequest) interface{}
	// DeleteUser(ctx context.Context, id string) interface{}
	PromoteandDemoteUser(ctx context.Context, id string, promotion UserPromotionRequest, role string) interface{}
}

type UserResponse struct {
	User User `json:"user"`
}

type UserUpdateRequest struct {
	ID                string `json:"id"`
	Full_Name         string `json:"full_name"`
	Username          string `json:"username" validate:"required"`
	Password          string `json:"password" validate:"required"`
	Profile_image_url string `json:"profile_image" `
	Contact           string `json:"contact"`
	Bio               string `json:"bio"`
}

type PromotionandDemotion struct {
	User_id string `json:"user_id"`
	Role    string `json:"role"`
}

type UserPromotionRequest struct {
	Action string `json:"action"`
}

type UserPromotionResponse struct {
	Message string `json:"message"`
}
