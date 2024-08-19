package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID         primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Author     string             `bson:"author" validate:"required"`
	Title      string             `bson:"title" validate:"required,min=1,max=255"`
	Content    string             `bson:"content" validate:"required,min=9"`
	Tags       []string           `bson:"tags"`
	Creater_id primitive.ObjectID `bson:"creater_id"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
	Comments   []Comment          `bson:"comments"`
	Blog_image string             `bson:"blog_image"`
}

type Comment struct {
	ID                 primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Blog_ID            primitive.ObjectID `bson:"blog_id"`
	Commentor_ID       primitive.ObjectID `bson:"commentor_id" validate:"required"`
	Commentor_username string             `bson:"commentor_username" validate:"required"`
	Content            string             `bson:"content" validate:"required,min=1,max=255"`
}

type BlogUsecase interface {
	CreateBlog(user_id string, blog Blog, role string) (Blog, error)
	GetBlogByID(blog_id string) (Blog, error)
	GetBlogs(pageNo string, pageSize string) ([]Blog, Pagination, error)
	UpdateBlogByID(user_id string, blog_id string, blog Blog) (Blog, error)
	DeleteBlogByID(user_id string, blog_id string, role string) ErrorResponse
	CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment Comment) error

	SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]Blog, Pagination, ErrorResponse)
	FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]Blog, Pagination, error)

	GetMyBlogs(user_id string, pageNo string, pageSize string) ([]Blog, Pagination, error)
	GetMyBlogByID(user_id string, blog_id string) (Blog, error)
}
type BlogRepository interface {
	CreateBlog(user_id string, blog Blog, role string) (Blog, error)
	GetBlogByID(blog_id string) (Blog, error)
	GetBlogs(pageNo int64, pageSize int64) ([]Blog, Pagination, error)
	UpdateBlogByID(user_id string, blog_id string, blog Blog) (Blog, error)
	DeleteBlogByID(user_id string, blog_id string) ErrorResponse
	CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment Comment) error

	SearchBlogByTitleAndAuthor(title string, author string, pageNo int64, pageSize int64) ([]Blog, Pagination, error)
	FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]Blog, Pagination, error)

	GetMyBlogs(user_id string, pageNo int64, pageSize int64) ([]Blog, Pagination, error)
	GetMyBlogByID(user_id string, blog_id string) (Blog, error)
}
