package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username           string             `json:"username" required:"true" bson:"username"`
	RefreshToken       string             `json:"refresh_token" bson:"refresh_token"`
	VerificationToken  string             `json:"verification_token" bson:"verification_token"`
	ResetPasswordToken int                `json:"reset_password_token" bson:"reset_password_token"`
	PasswordResetToken string             `json:"password_reset_token" bson:"password_reset_token"`
}
