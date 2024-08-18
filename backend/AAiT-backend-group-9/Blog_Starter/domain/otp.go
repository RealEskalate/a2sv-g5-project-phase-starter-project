package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOTP = "otp"
)

type Otp struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Email              string             `bson:"email,omitempty"`
	Otp                string             `bson:"otp,omitempty"`
	Expiration         time.Time          `bson:"expiration,omitempty"`
	PendingEmailUpdate string             `bson:"pendingEmailUpdate,omitempty"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt          time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}

type OtpRepository interface {
	SaveOtp(c context.Context, otp *Otp) error
	InvalidateOtp(c context.Context, otp *Otp) error
	GetOtpByEmail(c context.Context, email string) (Otp, error)
	GetByID(c context.Context, id string) (Otp, error)
}

type OtpUsecase interface {
	SaveOtp(c context.Context, otp *Otp) error
	InvalidateOtp(c context.Context, otp *Otp) error
	GetOtpByEmail(c context.Context, email string) (Otp, error)
	GetByID(c context.Context, id string) (Otp, error)
}
