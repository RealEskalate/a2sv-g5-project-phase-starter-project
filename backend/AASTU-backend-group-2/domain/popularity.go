package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Popularity domain model
type Like struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID primitive.ObjectID `json:"post_id" bson:"post_id"`
}

type DisLike struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID primitive.ObjectID `json:"post_id" bson:"post_id"`
}

type Comment struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID  primitive.ObjectID `json:"post_id" bson:"post_id"`
	Comment string             `json:"comment" bson:"comment"`
	Date    time.Time          `json:"date" bson:"date"`
}

// Popularity interfaces
type LikeUsecase interface {
	GetLikes(context context.Context, post_id string) ([]Like, *AppError)
	CreateLike(context context.Context, user_id string, post_id string) *AppError
	DeleteLike(context context.Context, like_id string) *AppError
}

type LikeRepository interface {
	GetLikes(post_id string) ([]Like, *AppError)
	CreateLike(user_id string, post_id string) *AppError
	DeleteLike(like_id string) *AppError
}

type DisLikeUsecase interface {
	GetDisLikes(context context.Context, post_id string) ([]DisLike, *AppError)
	CreateDisLike(context context.Context, user_id string, post_id string) *AppError
	DeleteDisLike(context context.Context, like_id string) *AppError
}

type DisLikeRepository interface {
	GetDisLikes(post_id string) ([]DisLike, *AppError)
	CreateDisLike(user_id string, post_id string) *AppError
	DeleteDisLike(like_id string) *AppError
}

type CommentUsecase interface {
	GetComments(ctx context.Context, BlogID string) ([]Comment, *AppError)
	CreateComment(ctx context.Context, BlogID string, userID string, comment Comment) *AppError
	UpdateComment(ctx context.Context, commentID string, comment Comment) *AppError
	DeleteComment(ctx context.Context, commentID string) *AppError
}

type CommentRepository interface {
	GetComments(BlogID string) ([]Comment, *AppError)
	CreateComment(BlogID, userID string, comment Comment) *AppError
	UpdateComment(commentID string, comment Comment) *AppError
	DeleteComment(commentID string) *AppError
}
