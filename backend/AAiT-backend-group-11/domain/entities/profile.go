package entities

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
    UserID         primitive.ObjectID `bson:"userId"`
    Bio            string             `bson:"bio"`
    ProfilePicture string             `bson:"profilePicture"`
    ContactInfo    ContactInfo        `bson:"contactInfo"`
    UpdatedAt      time.Time          `bson:"updatedAt"`
}

type ContactInfo struct {
    PhoneNumber string `bson:"phoneNumber"`
    Email       string `bson:"email"`
    Address     string `bson:"address"`
}
