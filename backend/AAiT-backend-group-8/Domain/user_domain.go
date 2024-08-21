package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	Email              string             `bson:"email"`
	Password           string             `bson:"password"`
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
