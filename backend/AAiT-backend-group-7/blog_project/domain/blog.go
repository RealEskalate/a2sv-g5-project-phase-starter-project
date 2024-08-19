package domain

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	AuthorID  int      `json:"author_id"`
	Content   string   `json:"content"`
	Comments  []Comment `json:"comments"`
	Likes     []Like   `json:"likes"`
	Dislikes  []Dislike `json:"dislikes"`
	Date      string   `json:"date"`
	Tags      []string `json:"tags"`
	Views     int      `json:"views"`
}

type Comment struct {
	ID      int    `json:"id" bson:"_id,omitempty"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type Like struct {
	ID     int    `json:"id" bson:"_id,omitempty"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

type Dislike struct {
	ID     int    `json:"id" bson:"_id,omitempty"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

type IBlogRepository interface {
	GetAllBlogs(ctx context.Context) ([]Blog, error)
	GetBlogByID(ctx context.Context, id int) (Blog, error)
	CreateBlog(ctx context.Context, blog Blog) (Blog, error)
	UpdateBlog(ctx context.Context, id int, blog Blog) (Blog, error)
	DeleteBlog(ctx context.Context, id int) error
}

type IBlogUsecase interface {
	GetAllBlogs(ctx context.Context) ([]Blog, error)
	GetBlogByID(ctx context.Context, id int) (Blog, error)
	CreateBlog(ctx context.Context, blog Blog) (Blog, error)
	UpdateBlog(ctx context.Context, id int, blog Blog) (Blog, error)
	DeleteBlog(ctx context.Context, id int) error
	AddComment(ctx context.Context, blogID int, comment Comment) (Blog, error)
	AddLike(ctx context.Context, blogID int, like Like) (Blog, error)
	AddDislike(ctx context.Context, blogID int, dislike Dislike) (Blog, error)
}

type IBlogController interface {
	GetAllBlogs(c *gin.Context)
	GetBlogByID(c *gin.Context)
	CreateBlog(c *gin.Context)
	UpdateBlog(c *gin.Context)
	DeleteBlog(c *gin.Context)
	AddComment(c *gin.Context)
	AddLike(c *gin.Context)
	AddDislike(c *gin.Context)
}
