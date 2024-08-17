package domain

import "github.com/gin-gonic/gin"

type BlogController interface {
	CreateBlog(c *gin.Context)
	GetBlog(c *gin.Context)
	GetBlogs(c *gin.Context)
	UpdateBlog(c *gin.Context)
	DeleteBlog(c *gin.Context)
	SearchBlogs(c *gin.Context)
	FiltersBlogs(c *gin.Context)
	GenerateBlog(c *gin.Context)
	LikeBlog(c *gin.Context)
}

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Authenticate(c *gin.Context)
	ForgotPassword(c *gin.Context)
	Logout(c *gin.Context)
	PromoteUser(c *gin.Context)
	DemoteUser(c *gin.Context)
	UpdateProfile(c *gin.Context)
}

type CommentController interface {
	CreateComment(c *gin.Context)
	GetComments(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}