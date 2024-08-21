package models


import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type User struct {
    ID                      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email                   string             `bson:"email" json:"email"`
    Password                string             `bson:"password" json:"password"`
    Role                    string             `bson:"role" json:"role"`
    Profile                 primitive.ObjectID `bson:"profile,omitempty" json:"profile"`
    CreatedAt               time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt               time.Time          `bson:"updated_at" json:"updated_at"`
    IsVerified              bool               `bson:"is_verified" json:"is_verified"`
    VerificationToken       string             `bson:"verification_token" json:"-"`
    TokenExp                time.Time          `bson:"token_exp" json:"-"`
    RefToken                string             `bson:"ref_token,omitempty" json:"-"`
}