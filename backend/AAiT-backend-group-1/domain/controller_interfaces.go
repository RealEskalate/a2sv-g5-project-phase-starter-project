package domain

import "github.com/gin-gonic/gin"

type ApiControllers interface {
	BlogController
	UserController
	BlogAssisstantController
}

type BlogController interface {
	CreateBlog(c *gin.Context)
	GetBlog(c *gin.Context)
	GetBlogs(c *gin.Context)
	UpdateBlog(c *gin.Context)
	DeleteBlog(c *gin.Context)
	SearchBlogs(c *gin.Context)
	FilterBlogs(c *gin.Context)
	LikeBlog(c *gin.Context)
	DislikeBlog(c *gin.Context)
	AddComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	EditComment(c *gin.Context)

}

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	ForgotPassword(c *gin.Context)
	Logout(c *gin.Context)
	PromoteUser(c *gin.Context)
	DemoteUser(c *gin.Context)
	UpdateProfile(c *gin.Context)
}

type BlogAssisstantController interface {
	GenerateBlog(c *gin.Context)
	EnhanceBlog(c *gin.Context)
	SuggestBlog(c *gin.Context)
}
