package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username           string             `json:"username" required:"true" bson:"username"`
	RefreshToken       string             `json:"refresh_token" bson:"refresh_token"`
	VerificationToken  string             `json:"verification_token" bson:"verification_token"`
	ResetPasswordCode  int                `json:"reset_password_code" bson:"reset_password_code"`
	PasswordResetToken string             `json:"password_reset_token" bson:"password_reset_token"`
}
