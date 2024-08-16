package Domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository interface {
	Login(ctx context.Context,user *User) (string, error, int)
	Register(ctx context.Context,user *User) (*OmitedUser, error, int)
}

type AuthUseCase interface {
	Login(c *gin.Context, user *User) (string, error, int)
	Register(c *gin.Context, user *User) (*OmitedUser, error, int)
}

type UserRepository interface {
	CreateUsers(ctx context.Context, user *User) (OmitedUser, error, int)
	GetUsers(ctx context.Context) ([]*OmitedUser, error, int)
	GetUsersById(ctx context.Context, id primitive.ObjectID, user OmitedUser) (OmitedUser, error, int)
	UpdateUsersById(ctx context.Context, id primitive.ObjectID, user User, curentuser OmitedUser) (OmitedUser, error, int)
	DeleteUsersById(ctx context.Context, id primitive.ObjectID, user OmitedUser) (error, int)
}

type UserUseCases interface {
	CreateUsers(c *gin.Context, user *User) (OmitedUser, error, int)
	GetUsers(c *gin.Context) ([]*OmitedUser, error, int)
	GetUsersById(c *gin.Context, id primitive.ObjectID, user OmitedUser) (OmitedUser, error, int)
	UpdateUsersById(c *gin.Context, id primitive.ObjectID, user User, curentuser OmitedUser) (OmitedUser, error, int)
	DeleteUsersById(c *gin.Context, id primitive.ObjectID, user OmitedUser) (error, int)
}
