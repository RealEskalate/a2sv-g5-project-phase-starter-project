package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	RegisterStart(cxt *gin.Context, user *User) Error
	RegisterEnd(cxt *gin.Context, token string) Error
	Login(context context.Context, username, password string) (map[string]string, Error)
	ForgotPassword(context context.Context, email string) Error
	ResetPassword(cxt *gin.Context, newPassword, confirmPassword, token string) Error
	Logout(context context.Context, token map[string]string) Error
	PromoteUser(context context.Context, userID string) Error
	DemoteUser(cxt context.Context, userID string) Error
	UpdateProfile(context context.Context, userID string, user *User) Error
}

type BlogUseCase interface {
	CreateBlog(blog *Blog, authorID string) Error
	GetBlog(blogID string) (*Blog, Error)
	GetBlogs() ([]Blog, Error)
	UpdateBlog(blogID string, blog *Blog) Error
	DeleteBlog(blogID string) Error
	SearchBlogsByTitle(title string) ([]Blog, Error)
	SearchBlogsByAuthor(author string) ([]Blog, Error)
	FilterBlogs(tags []string, dateAfter time.Time, popular bool) ([]Blog, Error)
	LikeBlog(userID, blogID string) Error
	AddComment(blogID string, comment *Comment) Error
	DeleteComment(blogID, commentID string) Error
	EditComment(blogID string, commentID string, comment *Comment) Error
	Like(blogId string, userID string) Error
	DisLike(blogId string, userID string) Error
}

type BlogAssistantUseCase interface {
	GenerateBlog(keywords []string, tone, audience string) (map[string]interface{}, Error)
	EnhanceBlog(content, command string) (map[string]interface{}, Error)
	SuggestBlog(industry string) ([]map[string]interface{}, Error)
}
