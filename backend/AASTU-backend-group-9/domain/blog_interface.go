package domain

import (
	// "blog/domain"
	"context"
	// "strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogRepository defines the methods to interact with the data layer for blogs.
type BlogRepository interface {
	CreateBlog(ctx context.Context, blog *Blog) *Error
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*Blog, *Error)
	GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*Blog, *Error)
	UpdateBlog(ctx context.Context, blog *Blog) *Error
	DeleteBlog(ctx context.Context, id primitive.ObjectID) *Error
	SearchBlogs(ctx context.Context, title string, author string) (*[]Blog, *Error)
	FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*Blog, *Error)
	AddComment(ctx context.Context, id primitive.ObjectID, comment *Comment) *Error
	HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID string) (bool, *Error)
	IncrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) *Error
	DecrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) *Error
}

// BlogUsecase defines the business logic methods for blogs.
type BlogUsecase interface {
	CreateBlog(ctx context.Context, blog *BlogCreationRequest, claims *JwtCustomClaims) (*BlogResponse, *Error)
	GetBlogByID(ctx context.Context, id primitive.ObjectID) (*BlogResponse, *Error)
	GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*BlogResponse, *Error)
	UpdateBlog(ctx context.Context, id primitive.ObjectID, blog *BlogUpdateRequest) (*BlogResponse, *Error)
	DeleteBlog(ctx context.Context, id primitive.ObjectID) *Error
	SearchBlogs(ctx context.Context, title string, author string) (*[]Blog, *Error)
	FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*Blog, *Error)
	TrackView(ctx context.Context, id primitive.ObjectID) *Error
	TrackLike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) *Error
	TrackDislike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) *Error
	AddComment(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) *Error
	GetComments(ctx context.Context, post_id primitive.ObjectID) ([]Comment, *Error)
	DeleteComment(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID, userID primitive.ObjectID) *Error
	UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) *Error
	AddReply(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, reply *Comment) *Error
	TrackCommentPopularity(ctx context.Context, postID, commentID, userID primitive.ObjectID, metric string) *Error
}

type PopularityRepository interface {
	HasUserLiked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, *Error)
	HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, *Error)
	UserInteractionsAdder(ctx context.Context, user UserInteraction) *Error
	UserInteractionsDelete(ctx context.Context, user UserInteraction) *Error
}
