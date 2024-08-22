package domain

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID    string             `json:"userid"`
	Bio       string             `json:"bio"`
	AvatarURL string             `json:"avatar_url"`
}

type ProfileGet struct {
	ID     primitive.ObjectID   `form:"_id" bson:"_id,omitempty"`
	UserID string               `form:"userid"`
	Bio    string               `form:"bio"`
	Avatar multipart.FileHeader `form:"avatar"`
}

type ProfileRepository interface {
	SaveProfile(*Profile) error
	FindProfile(userID string) (*Profile, error)
	DeleteProfile(userID string) error
	UpdateProfile(*Profile) error
}

type ProfileUsecase interface {
	SaveProfile(*ProfileGet) error
	FindProfile(userID string) (*Profile, error)
	DeleteProfile(userID string) error
	UpdateProfile(*ProfileGet) error
}

type ProfileHandler interface {
	SaveProfile(*gin.Context)
	FindProfile(*gin.Context)
	DeleteProfile(*gin.Context)
	UpdateProfile(*gin.Context)
}

type ProfileRouter interface {
	InitProfileRoutes(*gin.RouterGroup)
}
