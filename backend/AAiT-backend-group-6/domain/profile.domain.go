package domain

import (
	"context"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UpdateProfileDto struct {
	Avatar      *multipart.FileHeader `json:"avatar" form:"avatar"`
	UserProfile Profile       `json:"userProfile" form:"userProfile"`
}

type UpdateData map[string]interface{}
type Profile struct {
	ID primitive.ObjectID  `bson:"_id"`
	UserID  primitive.ObjectID `bson:"user_id"`
	Url      string `json:"url" bson:"url"`
	Name  	 string `json:"name" bson:"name"`
	Email 	 string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Bio 	 string `json:"bio" bson:"bio"`
}

type ProfileRepository interface{
	GetProfileByID(c context.Context, userID string) (*Profile, error)
	UpdateProfile(c context.Context, userID string, profile *Profile) error
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
	UpdateProfile(c context.Context, userID string, profile *UpdateProfileDto) error
}