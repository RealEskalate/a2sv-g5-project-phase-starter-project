package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlog = "blogs"
)

// this structure will be used when responding for blog request
type Blog struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	AuthorID      primitive.ObjectID `json:"author_id" bson:"author_id"`
	Title         string             `json:"title" bson:"title" binding:"required"`
	Tags          []string           `json:"tags" bson:"tags"`
	Content       string             `json:"content" bson:"content" binding:"required"`
	ViewCount     int                `json:"view_count" bson:"view_count"`
	LikeCount     int                `json:"like_count" bson:"like_count"`
	DislikeCount  int                `json:"dislike_count" bson:"dislike_count"`
	CommentsCount int                `json:"comments_count" bson:"comments_count"`
	Popularity    float64            `json:"popularity" bson:"popularity"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}

func (blog *Blog) UpdatePopularity() {
	blog.Popularity = float64(blog.ViewCount)*0.5 + float64(blog.LikeCount)*0.3 + float64(blog.CommentsCount)*0.2
}

// defines the structure for the blogs that will be  received from the request when creating and updating
type BlogIn struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	Tags      []string  `json:"tags" bson:"tags"`
	Content   string    `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type BlogUpdate struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	Tags      []string  `json:"tags" bson:"tags"`
	Content   string    `json:"content" bson:"content" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

const (
	CollectionComment = "comments"
)

// this structure for comments
type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// when creating and updating comments
type CommentIn struct {
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

const (
	CollectionReaction = "reactions"
)

// user reaction to the blog if liked or disliked
type Reaction struct {
	BlogID   primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Liked    bool               `json:"liked" bson:"liked"`
	Disliked bool               `json:"disliked" bson:"disliked"`
	Date     time.Time          `json:"date" bson:"date"`
}

// BlogRepository defines the methods required for data access related to blogs and comments.
type BlogRepository interface {
	GetByTags(c context.Context, tags []string, limit int64, page int64) ([]Blog, error)
	GetAllBlogs(c context.Context, limit int64, page int64) ([]Blog, error)
	GetBlogByID(c context.Context, blogID string) (Blog, error)
	GetByPopularity(c context.Context, limit int64, page int64) ([]Blog, error)
	Search(c context.Context, searchTerm string, limit int64, page int64) ([]Blog, error)
	CreateBlog(c context.Context, newBlog *Blog) (Blog, error)
	UpdateBlog(c context.Context, blogID string, updatedBlog *BlogUpdate) (Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SortByDate(c context.Context, limit int64, page int64) ([]Blog, error)
}

type CommentRepository interface {
	GetComments(c context.Context, blogID string) ([]Comment, error)
	CreateComment(c context.Context, blogID string, comment *Comment) (Comment, error)
	GetComment(c context.Context, blogID, commentID string) (Comment, error)
	UpdateComment(c context.Context, blogID, commentID string, updatedComment *Comment) (Comment, error)
	DeleteComment(c context.Context, blogID, commentID string) error
}

type LikeRepository interface {
	LikeBlog(c context.Context, blogID, userID string) error
	DislikeBlog(c context.Context, blogID, userID string) error
}

type BlogUsecase interface {
	GetByTags(c context.Context, tags []string) ([]Blog, error)
	GetAllBlogs(c context.Context) ([]Blog, error)
	GetBlogByID(c context.Context, blogID string) (Blog, error)
	GetByPopularity(c context.Context) ([]Blog, error)
	Search(c context.Context, searchTerm string) ([]Blog, error)
	CreateBlog(c context.Context, newBlog *Blog) (Blog, error)
	UpdateBlog(c context.Context, blogID string, updatedBlog *BlogUpdate) (Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SortByDate(c context.Context) ([]Blog, error)
}

type CommentUsecase interface {
	GetComments(c context.Context, blogID string) ([]Comment, error)
	CreateComment(c context.Context, blogID string, comment *Comment) (Comment, error)
	GetComment(c context.Context, blogID, commentID string) (Comment, error)
	UpdateComment(c context.Context, blogID, commentID string, updatedComment *Comment) (Comment, error)
	DeleteComment(c context.Context, blogID, commentID string) error
}

type LikeUsecase interface {
	LikeBlog(c context.Context, blogID, userID string) error
	DislikeBlog(c context.Context, blogID, userID string) error
}
