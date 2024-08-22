package domain

import (
	"github.com/gin-gonic/gin"
)

type Blog_Controller_interface interface {
	CreateBlog() gin.HandlerFunc
	GetOneBlog() gin.HandlerFunc
	GetBlogs() gin.HandlerFunc
	UpdateBlog() gin.HandlerFunc
	DeleteBlog() gin.HandlerFunc
	FilterBlog() gin.HandlerFunc
}

type Blog_Usecase_interface interface {
	CreateBlog(iblog PostBlog) (Blog, error)
	GetOneBlog(id string) (Blog, error)
	GetBlogs(limit int, page_number int) ([]Blog, error)
	UpdateBlog(id string, blog Blog) (Blog, error)
	DeleteBlog(id string) error
	FilterBlog(filters map[string]interface{}) ([]Blog, error)
}

type Blog_Repository_interface interface {
	CreateBlogDocument(blog Blog) (Blog, error)
	GetOneBlogDocument(id string) (Blog, error)
	GetBlogDocuments(offset int, limit int) ([]Blog, error)
	UpdateBlogDocument(id string, blog Blog) (Blog, error)
	DeleteBlogDocument(id string) error
	FilterBlogDocument(filters map[string]interface{}) ([]Blog, error)
}
