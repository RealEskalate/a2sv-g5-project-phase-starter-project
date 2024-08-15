package domain

import (
	"github.com/gin-gonic/gin"
)

type BlogControllerinterface interface {
	CreateBlog() gin.HandlerFunc
	GetOneBlog() gin.HandlerFunc
	GetBlogs() gin.HandlerFunc
	UpdateBlog() gin.HandlerFunc
	DeleteBlog() gin.HandlerFunc
	FilterBlog() gin.HandlerFunc
}

type BlogUsecase_interface interface {
	CreateBlog(blog Blog) (Blog, error)
	GetOneBlog(id string) ([]Blog, error)
	GetBlogs() ([]Blog, error)
	UpdateBlog(id string, blog Blog) (Blog, error)
	DeleteBlog(id string) error
	FilterBlog(map[string]string) ([]Blog, error)
}
