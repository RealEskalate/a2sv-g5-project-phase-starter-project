package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name" validate:"required,min=2,max=50"`
	LastName           string             `bson:"lastname"`
	Email              string             `bson:"email" validate:"required,email"`
	Password           string             `bson:"password" validate:"required"`
	Role               string             `bson:"role"`
	ImageUrl           string             `bson:"image_url"`
	CreatedAt          time.Time          `bson:"created_at"`
	Verified           bool               `bson:"verified"`
	VerificationToken  string             `bson:"verification_token"`
	PasswordResetToken string             `bson:"password_reset_token"`
}

type Credential struct {
	Email     string `json:"email" bson:"email"`
	Refresher string `json:"refresher" bson:"refresher"`
}
