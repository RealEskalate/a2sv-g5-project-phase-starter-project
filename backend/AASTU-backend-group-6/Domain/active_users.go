package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActiveUser struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Username     string             `json:"username" bson:"username"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
}
type ActiveUserRepository interface {
	CreateActiveUser(au ActiveUser, c context.Context) error
	FindActiveUserById(id primitive.ObjectID, c context.Context) (ActiveUser, error)
	DeleteActiveUserById(id primitive.ObjectID, c context.Context) error
}
