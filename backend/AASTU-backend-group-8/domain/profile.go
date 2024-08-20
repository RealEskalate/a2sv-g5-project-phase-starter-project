package domain

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID    string             `json:"userid"`
	Bio       string             `json:"bio"`
	AvatarURL string             `json:"avatar_url"`
}

type ProfileRepository interface {
	SaveProfile(*Profile) error
	FindProfile(userID string) (*Profile, error)
	DeleteProfile(userID string) error
	UpdateProfile(*Profile) error
}

type ProfileUsecase interface {
	SaveProfile(*Profile) error
	FindProfile(userID string) (*Profile, error)
	DeleteProfile(userID string) error
	UpdateProfile(*Profile) error
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
