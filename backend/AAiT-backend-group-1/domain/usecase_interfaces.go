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
	CreateBlog(blog *Blog , authorID string) Error
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
	EditComment(blogID string , commentID string, comment *Comment) Error
	Like(blogId string , userID string) Error
	DisLike(blogId string , userID string) Error
}

type BlogAssisstantUseCase interface {
	GenerateBlog(c *gin.Context)
	EnhanceBlog(c *gin.Context)
	SuggestBlog(c *gin.Context)
}
