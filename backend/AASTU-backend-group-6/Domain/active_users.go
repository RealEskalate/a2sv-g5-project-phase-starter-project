package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActiveUser struct {
	ID           primitive.ObjectID `json:"id"`
	Username     string `json:"username" bson:"username"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	UserAgent    string `json:"user_agent" bson:"user_agent"`
}
type ActiveUserRepository interface {
	CreateActiveUser(au ActiveUser, c context.Context) error
	FindActiveUserById(id string, c context.Context) (ActiveUser, error)
	DeleteActiveUser(id string, user_agent string, c context.Context) error
}
