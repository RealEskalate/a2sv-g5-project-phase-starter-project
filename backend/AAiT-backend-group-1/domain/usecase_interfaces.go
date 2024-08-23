package domain

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	RegisterStart(cxt *gin.Context, user *User) Error
	RegisterEnd(cxt *gin.Context, token string) Error
	Login(context *gin.Context, username, password string) (map[string]string, Error)
	ForgotPassword(context *gin.Context, email string) Error
	ResetPassword(cxt *gin.Context, newPassword, confirmPassword, token string, resetCode int) Error
	Logout(cxt *gin.Context, token map[string]string) Error
	PromoteUser(cxt *gin.Context, userID string) Error
	DemoteUser(cxt *gin.Context, userID string) Error
	UpdateProfile(cxt *gin.Context, userID string, user map[string]interface{}) Error
	ImageUpload(cxt *gin.Context, file *multipart.File, header *multipart.FileHeader) Error
}

type BlogUseCase interface {
	CreateBlog(blog *Blog, authorID string) Error
	GetBlog(blogID string, userID string) (*Blog, Error)
	GetBlogs(page_number string) ([]Blog, Error)
	UpdateBlog(blogID string, blog *Blog, userId string) Error
	DeleteBlog(blogID string , currUserId string) Error
	SearchBlogsByTitle(title string, page_number string) ([]Blog, Error)
	SearchBlogsByAuthor(author string, page_number string) ([]Blog, Error)
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
