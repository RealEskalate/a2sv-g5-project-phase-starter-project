package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
	Likes    []Like    `json:"likes"`
	Dislikes []Dislike `json:"dislikes"`
	Date     time.Time `json:"date"`
	Tags     []string  `json:"tags"`
	Views    int       `json:"views"`
}

type Comment struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}

type Like struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

type Dislike struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

type IBlogRepository interface {
	GetAllBlogs(ctx context.Context) ([]Blog, error)
	GetBlogByID(ctx context.Context, id int) (Blog, error)
	CreateBlog(ctx context.Context, blog Blog) (Blog, error)
	UpdateBlog(ctx context.Context, id int, blog Blog) (Blog, error)
	DeleteBlog(ctx context.Context, id int) error
	SearchByTitle(ctx context.Context, title string) ([]Blog, error)
	SearchByTags(ctx context.Context, tags []string) ([]Blog, error)
	SearchByAuthor(ctx context.Context, author string) ([]Blog, error)
	GetBlogsByPage(ctx context.Context, offset, limit int) ([]Blog, error)
	UpdateAuthorName(ctx context.Context, oldName, newName string) error
}

type IBlogUsecase interface {
	GetAllBlogs(ctx context.Context, sortOrder string, page, limit int) ([]Blog, error)
	GetBlogByID(ctx context.Context, id int) (Blog, error)
	CreateBlog(ctx context.Context, blog Blog) (Blog, error)
	UpdateBlog(ctx context.Context, id int, blog Blog) (Blog, error)
	DeleteBlog(ctx context.Context, id int) error
	AddComent(ctx context.Context, blogID int, authorID int, content string) (Blog, error)
	LikeBlog(ctx context.Context, blogID int, authorID int) (Blog, error)
	DislikeBlog(ctx context.Context, blogID int, authorID int) (Blog, error)
	Search(ctx context.Context, author string, tags []string, title string) ([]Blog, error)
	AiRecommendation(ctx context.Context, content string) (string, error)
}

type IBlogController interface {
	GetAllBlogs(c *gin.Context)
	CreateBlog(c *gin.Context)
	UpdateBlog(c *gin.Context)
	DeleteBlog(c *gin.Context)
	AddComment(c *gin.Context)
	LikeBlog(c *gin.Context)
	DislikeBlog(c *gin.Context)
	Search(c *gin.Context)
	AiRecommendation(c *gin.Context)
}

type AiService interface {
	GenerateContent(ctx context.Context, content string) (string, error)
}
