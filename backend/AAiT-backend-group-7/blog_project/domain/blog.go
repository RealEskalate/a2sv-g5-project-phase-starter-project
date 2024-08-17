package domain

import "github.com/gin-gonic/gin"

type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author_id int       `json:"author_id"`
	Content   string    `json:"content"`
	Comments  []Comment `json:"comments"`
	Likes     []Like    `json:"likes"`
	Dislikes  []Dislike `json:"dislikes"`
	Date      string    `json:"date"`
	Tags      []string  `json:"tags"`
	Views     int       `json:"views"`
}

type Comment struct {
	ID      int    `json:"id"`
	User_id int    `json:"user_id"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type Like struct {
	ID      int    `json:"id"`
	User_id int    `json:"user_id"`
	Date    string `json:"date"`
}

type Dislike struct {
	ID      int    `json:"id"`
	User_id int    `json:"user_id"`
	Date    string `json:"date"`
}

type IBlog_Repository interface {
	GetAllBlogs() ([]Blog, error)
	GetBlogByID(id int) (Blog, error)
	CreateBlog(blog Blog) (Blog, error)
	UpdateBlog(id int, blog Blog) (Blog, error)
	DeleteBlog(id int) error
	SearchByTitle(title string) ([]Blog, error)
	SearchByTags(tags []string) ([]Blog, error)
	SearchByAuthor(author_id int) ([]Blog, error)
}

type IBlog_Usecases interface {
	GetAllBlogs() ([]Blog, error)
	GetBlogByID(id int) (Blog, error)
	CreateBlog(blog Blog) (Blog, error)
	UpdateBlog(id int, blog Blog) (Blog, error)
	DeleteBlog(id int) error
	Search(author string, tags []string, title string) ([]Blog, error)
	LikeBlog(blog_id int, author_id int) (Like, error)
	DislikeBlog(blog_id int, author_id int) (Dislike, error)
	CommentBlog(blog_id int, author_id int, content string) (Comment, error)
}

type IBlog_Controller interface {
	GetAllBlogs(c *gin.Context) ([]Blog, error)
	GetBlogsByAuhorID(c *gin.Context) ([]Blog, error)
	CreateBlog(c *gin.Context) (Blog, error)
	UpdateBlog(c *gin.Context) (Blog, error)
	LikeBlog(c *gin.Context) (Blog, error)
	DislikeBlog(c *gin.Context) (Blog, error)
	CommentBlog(c *gin.Context) (Blog, error)
	DeleteBlog(c *gin.Context) error
	Search(c *gin.Context) ([]Blog, error)
}
