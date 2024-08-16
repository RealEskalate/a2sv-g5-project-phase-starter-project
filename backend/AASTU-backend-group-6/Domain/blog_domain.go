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
	Creater_id primitive.ObjectID `bson:"creater_id" validate:"required"`
	CreatedAt  time.Time          `bson:"createdAt" validate:"required"`
	UpdatedAt  time.Time          `bson:"updatedAt" validate:"required"`
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
	CreateBlog(user_id string, blog Blog) error
	GetBlogByID(blog_id string) (Blog, error)
	GetBlogs(pageNo string, pageSize string) ([]Blog, Pagination, error)
	UpdateBlogByID(user_id string, blog_id string, blog Blog) error
	DeleteBlogByID(user_id string, blog_id string) error
	CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment Comment) error

	SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]Blog, Pagination, error)
	FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]Blog, Pagination, error)

	GetMyBlogs(user_id string, pageNo string, pageSize string) ([]Blog, Pagination, error)
	GetMyBlogByID(user_id string, blog_id string) (Blog, error)
}
type BlogRepository interface {
	CreateBlog(user_id string, blog Blog) error
	GetBlogByID(blog_id string) (Blog, error)
	GetBlogs(pageNo string, pageSize string) ([]Blog, Pagination, error)
	UpdateBlogByID(user_id string, blog_id string, blog Blog) error
	DeleteBlogByID(user_id string, blog_id string) error
	CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment Comment) error

	SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]Blog, Pagination, error)
	FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]Blog, Pagination, error)

	GetMyBlogs(user_id string, pageNo string, pageSize string) ([]Blog, Pagination, error)
	GetMyBlogByID(user_id string, blog_id string) (Blog, error)
}
