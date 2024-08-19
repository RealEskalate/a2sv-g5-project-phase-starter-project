package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name" json:"name"`
	ProfilePictureUrl string             `bson:"profile_picture_url" json:"profile_picture_url"`
}