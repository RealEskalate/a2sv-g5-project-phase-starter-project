package Domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository interface {
	Login(ctx context.Context, user *User) (Tokens, error, int)
	Register(ctx context.Context, user *User) (*OmitedUser, error, int)
	Logout(ctx context.Context, user_id primitive.ObjectID) (error, int)
}

type AuthUseCase interface {
	Login(c *gin.Context, user *User) (Tokens, error, int)
	Register(c *gin.Context, user *User) (*OmitedUser, error, int)
	Logout(c *gin.Context, user_id primitive.ObjectID) (error, int)
}

type RefreshRepository interface {
	// UpdateToken(ctx context.Context, refreshToken string, userid primitive.ObjectID) (error, int)
	DeleteToken(ctx context.Context, userid primitive.ObjectID) (error, int)
	FindToken(ctx context.Context, userid primitive.ObjectID) (string, error, int)
	StoreToken(ctx context.Context, userid primitive.ObjectID, refreshToken string) (error, int)
}

type RefreshUseCase interface {
	// UpdateToken(c *gin.Context, refreshToken string, userid primitive.ObjectID) (error, int)
	DeleteToken(c *gin.Context, userid primitive.ObjectID) (error, int)
	FindToken(c *gin.Context, userid primitive.ObjectID) (string, error, int)
	StoreToken(c *gin.Context, userid primitive.ObjectID, refreshToken string) (error, int)
}

type BlogRepository interface {
}

type BlogUseCase interface {
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
