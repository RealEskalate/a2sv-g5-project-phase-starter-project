package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	Register(context context.Context, user *User) Error
	Login(context context.Context, username, password string) (string, Error)
	ForgotPassword(context context.Context, email string) Error
	Logout(context context.Context, token string) Error
	PromoteUser(context context.Context, userID string) Error
	DemoteUser(context context.Context, userID string) Error
	UpdateProfile(context context.Context, userID string, user *User) Error
}

type BlogUseCase interface {
	CreateBlog(blog *Blog) Error
	GetBlog(blogID string) (*Blog, Error)
	GetBlogs() ([]*Blog, Error)
	UpdateBlog(blogID string, blog *Blog) Error
	DeleteBlog(blogID string) Error
	SearchBlogs(title, author string) ([]*Blog, Error)
	FilterBlogs(tags []string, dateAfter time.Time, popular bool) ([]*Blog, Error)
	LikeBlog(userID, blogID string) Error

	AddComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	EditComment(c *gin.Context)
	Like(c gin.Context)
	DisLike(c gin.Context)
}

type BlogAssisstantUseCase interface {
	GenerateBlog(c *gin.Context)
	EnhanceBlog(c *gin.Context)
	SuggestBlog(c *gin.Context)
}
