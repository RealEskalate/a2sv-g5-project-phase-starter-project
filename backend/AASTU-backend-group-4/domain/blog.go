package domain

import (
	"context"
	"time"

	"github.com/google/generative-ai-go/genai"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author"`
	Content   string             `json:"content" bson:"content"`
	AuthorID  primitive.ObjectID `json:"author_id" bson:"author_id"`
	Tags      []string           `json:"tags" bson:"tags"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	ViewCount int                `bson:"view_count" json:"view_count"`
	Likes     int                `bson:"-" json:"likes"`    // Field for likes count
	Comments  int                `bson:"-" json:"comments"` // Field for comments count
}

type BlogUsecase interface {
	GenerateAIContent(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error)
	CreateBlog(ctx context.Context, blog *Blog) (*Blog, error)
	CreateComment(ctx context.Context, comment *CommentRequest) (*Blog, error)
	AddLike(ctx context.Context, like *LikeRequest) error
	GetBlogs(ctx context.Context, page, limit int, sortBy string) ([]Blog, int, error)
	UpdateBlog(ctx context.Context, blogID primitive.ObjectID, updatedBlog *Blog) error
	GetBlogByID(ctx context.Context, blogID primitive.ObjectID) (*Blog, error)
	DeleteBlog(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID, isAdmin bool) error
	RemoveComment(ctx context.Context, userID primitive.ObjectID, commentID primitive.ObjectID, isAdmin bool) error
	RemoveLike(ctx context.Context, userID primitive.ObjectID, likeID primitive.ObjectID, isAdmin bool) error
	SearchBlog(ctx context.Context, filter map[string]string) ([]*Blog, error)
	FilterBlog(ctx context.Context, filter FilterRequest) ([]*Blog, error)
}

type BlogRepository interface {
	CreateBlog(ctx context.Context, blog *Blog) error
	GetBlogByID(ctx context.Context, blogID primitive.ObjectID) (*Blog, error)
	GetPaginatedBlogs(ctx context.Context, page, limit int, sortBy string) ([]Blog, error)
	GetTotalBlogs(ctx context.Context) (int, error)
	UpdateBlog(ctx context.Context, authorID primitive.ObjectID, updatedBlog *Blog) error
	DeleteBlog(ctx context.Context, blogID primitive.ObjectID) error
	SearchBlog(ctx context.Context, filter map[string]string) ([]*Blog, error)
	FilterBlog(ctx context.Context, filter FilterRequest) ([]*Blog, error)
	// GetAllBlogs(ctx context.Context) ([]*Blog, error)
}

type BlogResponse struct {
	Blog     Blog
	Comments []Comment
	Likes    []Like
}
