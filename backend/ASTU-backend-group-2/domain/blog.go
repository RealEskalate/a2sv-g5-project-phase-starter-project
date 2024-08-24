package domain

import (
	"context"
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
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
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	AuthorID primitive.ObjectID `json:"author_id" bson:"author_id"`
	Title    string             `json:"title" bson:"title" binding:"required"`
	Tags     []string           `json:"tags" bson:"tags"`
	Content  string             `json:"content" bson:"content" binding:"required"`
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
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// when creating and updating comments
type CommentIn struct {
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CommentUpdate struct {
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	UpdatedAt time.Time          `json:"created_at" bson:"created_at"`
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
type BlogFilter struct {
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
	GetBlogByID(c context.Context, blogID string) (Blog, error)
	GetByPopularity(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	Search(c context.Context, searchTerm string, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	CreateBlog(c context.Context, newBlog *BlogIn) (Blog, error)
	BatchCreateBlog(c context.Context, newBlogs *[]BlogIn) error
	UpdateBlog(c context.Context, blogID string, updatedBlog *BlogIn) (Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SortByDate(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	UpdateLikeCount(c context.Context, blogID string, increment bool) error
	UpdateDislikeCount(c context.Context, blogID string, increment bool) error
	UpdateCommentCount(c context.Context, blogID string, increment bool) error
}

type CommentRepository interface {
	GetComments(c context.Context, blogID string, limit int64, page int64) ([]Comment, mongopagination.PaginationData, error)
	CreateComment(c context.Context, comment *Comment) (Comment, error)
	GetComment(c context.Context, commentID string) (Comment, error)
	UpdateComment(c context.Context, commentID string, updatedComment *CommentUpdate) (Comment, error)
	DeleteComment(c context.Context, commentID string) error
}

type ReactionRepository interface {
	Like(c context.Context, blogID, userID string) error
	Dislike(c context.Context, blogID, userID string) error
	RemoveLike(c context.Context, blogID, userID string) error
	RemoveDislike(c context.Context, blogID, userID string) error
	IsPostLiked(c context.Context, blogID, userID string) (bool, error)
	IsPostDisliked(c context.Context, blogID, userID string) (bool, error)
}

type BlogUsecase interface {
	GetByTags(c context.Context, tags []string, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	GetAllBlogs(c context.Context, filter BlogFilter) ([]Blog, mongopagination.PaginationData, error)
	GetBlogByID(c context.Context, blogID string) (Blog, error)
	GetByPopularity(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	Search(c context.Context, searchTerm string, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
	CreateBlog(c context.Context, userID string, newBlog *BlogIn) (Blog, error)
	BatchCreateBlog(c context.Context, newBlogs *[]BlogIn) error
	UpdateBlog(c context.Context, blogID string, updatedBlog *BlogIn) (Blog, error)
	DeleteBlog(c context.Context, blogID string) error
	SortByDate(c context.Context, limit int64, page int64) ([]Blog, mongopagination.PaginationData, error)
}

type CommentUsecase interface {
	GetComments(c context.Context, blogID string, limit int64, page int64) ([]Comment, mongopagination.PaginationData, error)
	CreateComment(c context.Context, userID string, blogID string, comment *CommentIn) (Comment, error)
	GetComment(c context.Context, commentID string) (Comment, error)
	UpdateComment(c context.Context, commentID string, updatedComment *CommentUpdate) (Comment, error)
	DeleteComment(c context.Context, commentID string) error
}

type ReactionUsecase interface {
	ToggleLike(c context.Context, blogID, userID string) error
	ToggleDislike(c context.Context, blogID, userID string) error
}
