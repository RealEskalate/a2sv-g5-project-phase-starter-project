package entities

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/forms"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlog = "blogs"
)

type Blog struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Author        User               `json:"author,omitempty" bson:"author"`
	Title         string             `json:"title,omitempty" bson:"title" binding:"required"`
	Tags          []string           `json:"tags" bson:"tags"`
	Content       string             `json:"content,omitempty" bson:"content" binding:"required"`
	ViewCount     int                `json:"view_count,omitempty" bson:"view_count"`
	LikeCount     int                `json:"like_count,omitempty" bson:"like_count"`
	DislikeCount  int                `json:"dislike_count,omitempty" bson:"dislike_count"`
	CommentsCount int                `json:"comments_count,omitempty" bson:"comments_count"`
	Popularity    float64            `json:"popularity,omitempty" bson:"popularity"`
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
	AuthorName    string             `json:"author_name,omitempty" bson:"author_name"`
}

func (blog *Blog) UpdatePopularity() {
	blog.Popularity = float64(blog.ViewCount)*0.5 + float64(blog.LikeCount)*0.3 + float64(blog.CommentsCount)*0.2
}

type BlogUpdate struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	Tags      []string  `json:"tags" bson:"tags"`
	Content   string    `json:"content" bson:"content" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type BlogFilter struct {
	Search         string
	Title          string
	Tags           []string
	DateFrom       time.Time
	DateTo         time.Time
	Limit          int64
	Pages          int64
	PopularityFrom int
	PopularityTo   int
}

// BlogRepository defines the methods required for data access related to blogs and comments.
type BlogRepository interface {
	GetByTags(c context.Context, tags []string, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	GetAllBlogs(c context.Context, filter bson.M, blogFilter BlogFilter) ([]Blog, mongopagination.PaginationData, error)
	GetBlogByID(c context.Context, blogID string, view bool) (Blog, error)
	GetByPopularity(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	CreateBlog(c context.Context, newBlog *Blog) (Blog, error)
	BatchCreateBlog(c context.Context, newBlogs *[]Blog) error
	UpdateBlog(c context.Context, blogID string, updatedBlog *BlogUpdate) (Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SortByDate(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	UpdateLikeCount(c context.Context, blogID string, increment bool) error
	UpdateDislikeCount(c context.Context, blogID string, increment bool) error
	UpdateCommentCount(c context.Context, blogID string, increment bool) error
}

type BlogUsecase interface {
	GetByTags(c context.Context, tags []string, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	GetAllBlogs(c context.Context, filter BlogFilter) ([]Blog, mongopagination.PaginationData, error)
	GetBlogByID(c context.Context, blogID string) (Blog, error)
	GetByPopularity(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	CreateBlog(c context.Context, newBlog *forms.BlogForm, user *User) (Blog, error)
	BatchCreateBlog(c context.Context, newBlogs *[]forms.BlogForm, user *User) error
	UpdateBlog(c context.Context, blogID string, updatedBlog *forms.BlogForm) (Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SortByDate(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
}
