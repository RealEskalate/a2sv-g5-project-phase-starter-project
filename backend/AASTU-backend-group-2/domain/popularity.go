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
	UserID  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BlogID  primitive.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	Comment string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Date    time.Time          `json:"date,omitempty" bson:"date,omitempty"`
}

// Popularity interfaces
type LikeUsecase interface {
	GetLikes(context context.Context, post_id string) ([]Like, error)
	CreateLike(context context.Context, user_id string, post_id string) error
	DeleteLike(context context.Context, like_id string) error
}

type LikeRepository interface {
	GetLikes(post_id string) ([]Like, error)
	CreateLike(user_id string, post_id string) error
	DeleteLike(like_id string) error
}

type DisLikeUsecase interface {
	GetDisLikes(context context.Context, post_id string) ([]DisLike, error)
	CreateDisLike(context context.Context, user_id string, post_id string) error
	DeleteDisLike(context context.Context, like_id string) error
}

type DisLikeRepository interface {
	GetDisLikes(post_id string) ([]DisLike, error)
	CreateDisLike(user_id string, post_id string) error
	DeleteDisLike(like_id string) error
}

type CommentUsecase interface {
	GetComments(ctx context.Context, BlogID string) ([]Comment, error)
	CreateComment(ctx context.Context, BlogID string, userID string, comment Comment) error
	UpdateComment(ctx context.Context, commentID string, comment Comment) error
	DeleteComment(ctx context.Context, commentID string) error
}

type CommentRepository interface {
	GetComments(BlogID string) ([]Comment, error)
	CreateComment(BlogID, userID string, comment Comment) error
	UpdateComment(commentID string, comment Comment) error
	DeleteComment(commentID string) error
}
