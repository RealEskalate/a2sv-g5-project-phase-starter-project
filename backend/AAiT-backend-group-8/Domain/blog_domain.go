package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Blog struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	Title        string             `bson:"title" json:"title"`
	Body         string             `bson:"body" json:"body"`
	Tags         []string           `bson:"tags" json:"tags"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	LastUpdated  time.Time          `bson:"last_updated" json:"last_updated"`
	AuthorName   string             `bson:"author_name" json:"author_name"`
	AuthorID     primitive.ObjectID `bson:"author_id" json:"author_id"`
	ViewCount    int                `bson:"view_count" json:"view_count"`
	LikeCount    int                `bson:"like_count" json:"like_count"`
	CommentCount int                `bson:"comment_count" json:"comment_count"`
}

type SearchCriteria struct {
	Title     string    `form:"title"`
	Author    string    `form:"author"`
	Tags      []string  `form:"tags"`
	StartDate time.Time `form:"startDate"`
	EndDate   time.Time `form:"endDate"`
	MinViews  int       `form:"minViews"`
	SortBy    string    `form:"sortBy"`
	Order     string    `form:"order"`
	Page      int       `form:"page"`
	PageSize  int       `form:"pageSize"`
}

type IBlogRepository interface {
	Search(criteria *SearchCriteria) (*[]Blog, error)
	Create(blog *Blog) error
	UpdateCommentCount(id string, inc bool) error
	UpdateLikeCount(id string, inc bool) error
	FindAll(page int, pageSize int, sortBy string) (*[]Blog, error)
	FindByID(ID string) (*Blog, error)
	Delete(ID string) error
	UpdateViewCount(id string) error
	Update(blog *Blog) error
}

type IBlogUseCase interface {
	SearchBlog(criteria *SearchCriteria) (*[]Blog, error)
	CreateBlog(blog *Blog) error
	UpdateBlogViewCount(id string) error
	UpdateBlogCommentCount(id string, inc bool) error
	UpdateBlogLikeCount(id string, inc bool) error
	GetAllBlogs(page int, pageSize int, sortBy string) (*[]Blog, error)
	GetBlogByID(ID string) (*Blog, error)
	DeleteBlog(ID string) error
	UpdateBlog(blog *Blog) error
}
