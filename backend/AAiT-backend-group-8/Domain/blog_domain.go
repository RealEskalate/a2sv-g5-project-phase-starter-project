package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `bson:"title" json:"title" validate:"required,min=3,max=100"`
	Body         string             `bson:"body" json:"body" validate:"required,min=10"`
	Tags         []string           `bson:"tags" json:"tags"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	LastUpdated  time.Time          `bson:"last_updated" json:"last_updated"`
	AuthorName   string             `bson:"author_name" json:"author_name" validate:"max=100"`
	AuthorID     primitive.ObjectID `bson:"author_id" json:"author_id"`
	ViewCount    int                `bson:"view_count" json:"view_count"`
	LikeCount    int                `bson:"like_count" json:"like_count"`
	CommentCount int                `bson:"comment_count" json:"comment_count"`
}

type SearchCriteria struct {
	Title     string    `form:"title" bson:"title" json:"title" validate:"required, min=3,max=100"`
	Author    string    `form:"author" bson:"author" json:"author" validate:"max=100"`
	Tags      []string  `form:"tags" bson:"tags" json:"tags"`
	StartDate time.Time `form:"startDate" bson:"start_date" json:"start_date"`
	EndDate   time.Time `form:"endDate" bson:"end_date" json:"end_date"`
	MinViews  int       `form:"minViews" bson:"min_views" json:"min_views"`
	SortBy    string    `form:"sortBy" bson:"sort_by" json:"sort_by"`
	Page      int       `form:"page" bson:"page" json:"page"`
	PageSize  int       `form:"pageSize" bson:"page_size" json:"page_size"`
}

type IBlogRepository interface {
	Search(criteria *SearchCriteria) ([]Blog, error)
	Create(blog *Blog) error
	UpdateCommentCount(id string, inc bool) error
	UpdateLikeCount(id string, inc bool) error
	FindAll(page int, pageSize int, sortBy string) ([]Blog, error)
	FindByID(ID string) (*Blog, error)
	Delete(ID string) error
	UpdateViewCount(id string) error
	Update(blog *Blog) error
}

type IBlogUseCase interface {
	SearchBlog(criteria *SearchCriteria) ([]Blog, error)
	CreateBlog(blog *Blog) error
	UpdateBlogViewCount(id string) error
	UpdateBlogCommentCount(id string, inc bool) error
	UpdateBlogLikeCount(id string, inc bool) error
	GetAllBlogs(page int, pageSize int, sortBy string) ([]Blog, error)
	GetBlogByID(ID string) (*Blog, error)
	DeleteBlog(ID string) error
	UpdateBlog(blog *Blog) error
}
