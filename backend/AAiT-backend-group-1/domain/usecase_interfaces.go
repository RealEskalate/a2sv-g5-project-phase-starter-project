package domain

import (
	"time"
	"github.com/google/uuid"
)

type UserUseCase interface {
	Register(user *User) Error
	Login(email, password string) (string, string, Error)
	Authenticate(token string) (string, Error)
	ForgotPassword(email string) Error
	Logout(token string) Error
	PromoteUser(userID uuid.UUID) Error
	DemoteUser(userID uuid.UUID) Error
	UpdateProfile(userID uuid.UUID, user *User) Error
}

type BlogUseCase interface {
	CreateBlog(blog *Blog) Error
	GetBlog(blogID uuid.UUID) (*Blog, Error)
	GetBlogs() ([]*Blog, Error)
	UpdateBlog(blogID uuid.UUID, blog *Blog) Error
	DeleteBlog(blogID uuid.UUID) Error
	SearchBlogs(title, author string) ([]*Blog, Error)
	FiltersBlogs(tags []string, minimumDate time.Time, popular bool) ([]*Blog, Error)
	GenerateBlog(keywords []string) (*Blog, Error)
	LikeBlog(userID, blogID uuid.UUID) Error
}

type CommentUseCase interface {
	CreateComment(comment *Comment) Error
	// GetComment(commentID uuid.UUID) (*Comment, Error)
	GetComments(blogID uuid.UUID) ([]*Comment, Error)
	UpdateComment(blogID, commentID uuid.UUID, comment *Comment) Error
	DeleteComment(blogID, commentID uuid.UUID) Error
}