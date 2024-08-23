package domain

import (
	// "blog/domain"
	"context"
	// "strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogRepository defines the methods to interact with the data layer for blogs.
type BlogRepository interface {
	CreateBlog(ctx context.Context, blog *Blog) error
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*Blog, error)
	GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*Blog, error)
	UpdateBlog(ctx context.Context, blog *Blog) error
	DeleteBlog(ctx context.Context, id primitive.ObjectID) error
	SearchBlogs(ctx context.Context, title string, author string) (*[]Blog, error)
	FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*Blog, error)
	AddComment(ctx context.Context, id primitive.ObjectID, comment *Comment) error
	HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID string) (bool, error)
	IncrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error
	DecrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error
}

// BlogUsecase defines the business logic methods for blogs.
type BlogUsecase interface {
	CreateBlog(ctx context.Context, blog *BlogCreationRequest, claims *JwtCustomClaims) (*BlogResponse, error)
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*BlogResponse, error)
	GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*BlogResponse, error)
	UpdateBlog(ctx context.Context, id primitive.ObjectID, blog *BlogUpdateRequest) (*BlogResponse, error)
	DeleteBlog(ctx context.Context, id primitive.ObjectID) error
	SearchBlogs(ctx context.Context, title string, author string) (*[]Blog, error)
	FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*Blog, error)
	TrackView(ctx context.Context, id primitive.ObjectID) error
	TrackLike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error
	TrackDislike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error
	AddComment(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) error
	GetComments(ctx context.Context, post_id primitive.ObjectID) ([]Comment, error)
	DeleteComment(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID, userID primitive.ObjectID) error
	UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) error
}

type PopularityRepository interface {
	HasUserLiked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, error)
	HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, error)
	UserInteractionsAdder(ctx context.Context, user UserInteraction) error
	UserInteractionsDelete(ctx context.Context, user UserInteraction) error
}
