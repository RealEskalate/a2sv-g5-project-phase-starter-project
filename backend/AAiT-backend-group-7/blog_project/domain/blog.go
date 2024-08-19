package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	AuthorID int       `json:"author_id"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
	Likes    []Like    `json:"likes"`
	Dislikes []Dislike `json:"dislikes"`
	Date     string    `json:"date"`
	Tags     []string  `json:"tags"`
	Views    int       `json:"views"`
}

type Comment struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type Like struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

type Dislike struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

type IBlogRepository interface {
	GetAllBlogs(ctx context.Context) ([]Blog, error)
	GetBlogByID(ctx context.Context, id int) (Blog, error)
	CreateBlog(ctx context.Context, blog Blog) (Blog, error)
	UpdateBlog(ctx context.Context, id int, blog Blog) (Blog, error)
	DeleteBlog(ctx context.Context, id int) error
	SearchByTitle(ctx context.Context, title string) ([]Blog, error)
	SearchByTags(ctx context.Context, tags []string) ([]Blog, error)
	SearchByAuthor(ctx context.Context, authorID int) ([]Blog, error)
}

type IBlogUsecases interface {
	GetAllBlogs(ctx context.Context) ([]Blog, error)
	GetBlogByID(ctx context.Context, id int) (Blog, error)
	CreateBlog(ctx context.Context, blog Blog) (Blog, error)
	UpdateBlog(ctx context.Context, id int, blog Blog) (Blog, error)
	DeleteBlog(ctx context.Context, id int) error
	Search(ctx context.Context, author string, tags []string, title string) ([]Blog, error)
	LikeBlog(ctx context.Context, blogID int, authorID int) (Like, error)
	DislikeBlog(ctx context.Context, blogID int, authorID int) (Dislike, error)
	CommentBlog(ctx context.Context, blogID int, authorID int, content string) (Comment, error)
}

type IBlogController interface {
	GetAllBlogs(c *gin.Context)
	// GetBlogsByAuthorID(c *gin.Context)
	// GetBlogByID(c *gin.Context)
	CreateBlog(c *gin.Context)
	UpdateBlog(c *gin.Context)
	DeleteBlog(c *gin.Context)
	LikeBlog(c *gin.Context)
	DislikeBlog(c *gin.Context)
	CommentBlog(c *gin.Context)
	Search(c *gin.Context)
}
