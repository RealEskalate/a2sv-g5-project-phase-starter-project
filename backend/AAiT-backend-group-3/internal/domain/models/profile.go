package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	PhoneNum   string             `bson:"phone_num" json:"phone_num"`
	Bio        string             `bson:"bio" json:"bio"`
	ProfilePic string             `bson:"profile_pic" json:"profile_pic"`
}
