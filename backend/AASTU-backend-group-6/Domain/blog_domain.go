package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID            primitive.ObjectID   `bson:"_id,omitempity" json:"id" `
	Author        string               `bson:"author" validate:"required"`
	Title         string               `bson:"title" validate:"required,min=1,max=255"`
	Content       string               `bson:"content" validate:"required,min=9"`
	Tags          []string             `bson:"tags"`
	Creator_id    primitive.ObjectID   `bson:"creater_id"`
	CreatedAt     time.Time            `bson:"createdAt"`
	UpdatedAt     time.Time            `bson:"updatedAt"`
	ViewCount     int64                `bson:"view_count"`
	Commenters_ID []primitive.ObjectID `bson:"comment_ids"`
	LikeCount     int64                `bson:"like_count"`
	DisLikeCount  int64                `bson:"dislike_count"`
	Popularity    int64                `bson:"popularity"`
	Blog_image    string               `bson:"blog_image"`
	Deleted       bool                 `bson:"deleted"`
	DeletedAt     time.Time            `bson:"deletedAt"`
	Comments      []Comment   `bson:"comments"`
}

type Comment struct {
	ID                 primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Blog_ID            primitive.ObjectID `bson:"blog_id"`
	Commentor_ID       primitive.ObjectID `bson:"commentor_id"`
	Commentor_username string             `bson:"commentor_username"`
	Content            string             `bson:"content" validate:"required,min=1,max=255"`
	Deleted            bool               `bson:"deleted"`
	DeletedAt          time.Time          `bson:"deletedAt"`
}

type BlogUsecase interface {
	CreateBlog(user_id string, blog Blog, creator_id string) (Blog, error)
	GetBlogByID(blog_id string, isCalled bool) (Blog, error)
	GetBlogs(pageNo string, pageSize string, popularity string) ([]Blog, Pagination, error)
	UpdateBlogByID(user_id string, blog_id string, blog Blog) (Blog, error)
	DeleteBlogByID(user_id string, blog_id string) ErrorResponse
	CommentOnBlog(user_id string, comment Comment) error
	ReactOnBlog(user_id string, reactionType string, blog_id string) ErrorResponse

	SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string, popularity string) ([]Blog, Pagination, ErrorResponse)
	FilterBlogsByTag(tags []string, pageNo string, pageSize string, startDate string, endDate string, popularity string) ([]Blog, Pagination, error)

	GetMyBlogs(user_id string, pageNo string, pageSize string, popularity string) ([]Blog, Pagination, error)
	GetMyBlogByID(user_id string, blog_id string) (Blog, error)
}
type BlogRepository interface {
	CreateBlog(user_id string, blog Blog, creator_id string) (Blog, error)
	GetBlogByID(blog_id string, isCalled bool) (Blog, error)
	GetBlogs(pageNo int64, pageSize int64, popularity string) ([]Blog, Pagination, error)
	UpdateBlogByID(user_id string, blog_id string, blog Blog) (Blog, error)
	DeleteBlogByID(user_id string, blog_id string) ErrorResponse
	CommentOnBlog(user_id string, comment Comment) error
	ReactOnBlog(user_id string, reactionType bool, blog_id string) ErrorResponse

	SearchBlogByTitleAndAuthor(title string, author string, pageNo int64, pageSize int64, popularity string) ([]Blog, Pagination, error)
	FilterBlogsByTag(tags []string, pageNo int64, pageSize int64, startDate time.Time, endDate time.Time, popularity string) ([]Blog, Pagination, error)

	GetMyBlogs(user_id string, pageNo int64, pageSize int64, popularity string) ([]Blog, Pagination, error)
	GetMyBlogByID(user_id string, blog_id string) (Blog, error)
	GetUserRoleByID(id string) (string, error)
}
