package controller

import (
	"github.com/gin-gonic/gin"
)

// interface for blog controllers
type BlogController interface {
	GetBlogs() gin.HandlerFunc
	GetBlog() gin.HandlerFunc
	CreateBlog() gin.HandlerFunc
	UpdateBlog() gin.HandlerFunc
	DeleteBlog() gin.HandlerFunc
	GetComments() gin.HandlerFunc
	CreateComment() gin.HandlerFunc
	GetComment() gin.HandlerFunc
	UpdateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
	CreateLike() gin.HandlerFunc
}
