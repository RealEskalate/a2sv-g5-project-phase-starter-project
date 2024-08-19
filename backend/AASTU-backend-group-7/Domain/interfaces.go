package Domain

import (
	"blogapp/Dtos"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository interface {
	Login(ctx context.Context, user *User) (Tokens, error, int)
	Register(ctx context.Context, user *Dtos.RegisterUserDto) (*OmitedUser, error, int)
	Logout(ctx context.Context, user_id primitive.ObjectID) (error, int)
}

type AuthUseCase interface {
	Login(c *gin.Context, user *User) (Tokens, error, int)
	Register(c *gin.Context, user *Dtos.RegisterUserDto) (*OmitedUser, error, int)
	Logout(c *gin.Context, user_id primitive.ObjectID) (error, int)
}

type RefreshRepository interface {
	UpdateToken(ctx context.Context, refreshToken string, userid primitive.ObjectID) (error, int)
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
	CreateBlog(ctx context.Context, post *Post) (error, int)
	GetPostBySlug(ctx context.Context, slug string) ([]*Post, error, int)
	GetPostByAuthorID(ctx context.Context, authorID primitive.ObjectID) ([]*Post, error, int)
	GetPostByID(ctx context.Context, id primitive.ObjectID) (*Post, error, int)
	UpdatePostByID(ctx context.Context, id primitive.ObjectID, post *Post) (error, int)
}

type BlogUseCase interface {
	CreateBlog(c *gin.Context, post *Post) (error, int)
	GetPostBySlug(c *gin.Context, slug string) ([]*Post, error, int)
	GetPostByAuthorID(c *gin.Context, authorID primitive.ObjectID) ([]*Post, error, int)
	GetPostByID(c *gin.Context, id primitive.ObjectID) (*Post, error, int)
	UpdatePostByID(c *gin.Context, id primitive.ObjectID, post *Post) (error, int)
}

type CommentRepository interface {
	CommentOnPost(ctx context.Context, comment *Comment, objID primitive.ObjectID) (error, int)
	GetCommentByID(ctx context.Context, id primitive.ObjectID) (*Comment, error, int)
	EditComment(ctx context.Context, id primitive.ObjectID, comment *Comment) (error, int)
	GetUserComments(ctx context.Context, id primitive.ObjectID) ([]*Comment, error, int)
}

type CommentUseCase interface {
	CommentOnPost(c *gin.Context, comment *Comment, objID primitive.ObjectID) (error, int)
	GetCommentByID(c *gin.Context, id primitive.ObjectID) (*Comment, error, int)
	EditComment(c *gin.Context, id primitive.ObjectID, comment *Comment) (error, int)
	GetUserComments(c *gin.Context, id primitive.ObjectID) ([]*Comment, error, int)
}

type TagRepository interface {
	CreateTag(ctx context.Context, tag *Tag) (error, int)
	DeleteTag(ctx context.Context, id primitive.ObjectID) (error, int)
}

type TagUseCase interface {
	CreateTag(c *gin.Context, tag *Tag) (error, int)
	DeleteTag(c *gin.Context, id primitive.ObjectID) (error,int)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (OmitedUser, error, int)
	GetUsers(ctx context.Context) ([]*OmitedUser, error, int)
	GetUsersById(ctx context.Context, id primitive.ObjectID, user AccessClaims) (OmitedUser, error, int)
	UpdateUsersById(ctx context.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteUsersById(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
	PromoteUser(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	DemoteUser(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
}

type UserUseCases interface {
	CreateUser(c *gin.Context, user *User) (OmitedUser, error, int)
	GetUsers(c *gin.Context) ([]*OmitedUser, error, int)
	GetUsersById(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	UpdateUsersById(c *gin.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteUsersById(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
	PromoteUser(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	DemoteUser(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
}

type ProfileUseCases interface {
	GetProfile(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	UpdateProfile(c *gin.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteProfile(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
}

type ProfileRepository interface {
	GetProfile(ctx context.Context, id primitive.ObjectID, user AccessClaims) (OmitedUser, error, int)
	UpdateProfile(ctx context.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteProfile(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
}
