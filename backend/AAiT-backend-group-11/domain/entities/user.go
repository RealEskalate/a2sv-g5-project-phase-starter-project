package entities

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID            primitive.ObjectID `bson:"_id,omitempty"`
    Username      string             `bson:"username" binding:"required"`
    Email         string             `bson:"email" binding:"required"`
    Password      string             `bson:"password" binding:"required"`
    Profile       Profile            `bson:"profile"`
    IsVerified    bool               `bson:"isVerified"`
    Role          string             `bson:"role"`
    CreatedAt     time.Time          `bson:"createdAt"`
    UpdatedAt     time.Time          `bson:"updatedAt"`
}



