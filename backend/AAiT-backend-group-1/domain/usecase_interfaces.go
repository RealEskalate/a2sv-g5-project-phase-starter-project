package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	Register(user *User) Error
	Login(username, password string) (string, Error)
	ForgotPassword(email string) Error
	Logout(token string) Error
	PromoteUser(userID string) Error
	DemoteUser(userID string) Error
	UpdateProfile(userID string, user *User) Error
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

type BlogAssistantUseCase interface {
	GenerateBlog(keywords []string, tone, audience string) (string, Error)
	EnhanceBlog(content, command string) (string, Error)
	SuggestBlog(industry string) ([]string, Error)
}