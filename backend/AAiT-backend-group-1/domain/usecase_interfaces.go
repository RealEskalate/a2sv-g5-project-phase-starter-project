package domain

import (
	"context"
	"time"
)

type UserUseCase interface {
	Register(context context.Context, user *User) Error
	Login(context context.Context, username, password string) (map[string]string, Error)
	ForgotPassword(context context.Context, email string) Error
	Logout(context context.Context, token map[string]string) Error
	PromoteUser(context context.Context, userID string) Error
	DemoteUser(context context.Context, userID string) Error
	UpdateProfile(context context.Context, userID string, user *User) Error
}

type BlogUseCase interface {
<<<<<<< HEAD
	CreateBlog(blog *Blog, authorID string) Error
=======
	CreateBlog(blog *Blog , authorID string) Error
>>>>>>> d260d430 (like and comment crud operations finished)
	GetBlog(blogID string) (*Blog, Error)
	GetBlogs() ([]Blog, Error)
	UpdateBlog(blogID string, blog *Blog) Error
	DeleteBlog(blogID string) Error
	SearchBlogsByTitle(title string) ([]Blog, Error)
	SearchBlogsByAuthor(author string) ([]Blog, Error)
	FilterBlogs(tags []string, dateAfter time.Time, popular bool) ([]Blog, Error)
	LikeBlog(userID, blogID string) Error
<<<<<<< HEAD

	AddComment(blogID string, comment *Comment) Error
	DeleteComment(blogID, commentID string) Error
	EditComment(blogID string, commentID string, comment *Comment) Error
	Like(blogId string, userID string) Error
	DisLike(blogId string, userID string) Error
=======
	
	AddComment(blogID string, comment *Comment) Error
	DeleteComment(blogID, commentID string) Error
	EditComment(blogID string , commentID string, comment *Comment) Error
	Like(blogId string , userID string) Error
	DisLike(blogId string , userID string) Error
>>>>>>> a142a8e6 (blog controller added)
}

type BlogAssistantUseCase interface {
	GenerateBlog(keywords []string, tone, audience string) (string, Error)
	EnhanceBlog(content, command string) (string, Error)
	SuggestBlog(industry string) ([]string, Error)
}
