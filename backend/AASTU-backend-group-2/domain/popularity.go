package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Popularity domain model
type Like struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
}

type DisLike struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
}

type Comment struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID  primitive.ObjectID `json:"post_id" bson:"post_id"`
	Content string             `json:"comment" bson:"comment"`
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
	GetComments(context context.Context, post_id string) ([]Comment, error)
	CreateComment(context context.Context, post_id string, user_id string, comment Comment) error
	DeleteComment(context context.Context, comment_id string) error
	UpdateComment(context context.Context, comment_id string) error
}

type CommentRepository interface {
	GetComments(post_id string) ([]Comment, error)
	CreateComment(post_id string, user_id string, comment Comment) error
	DeleteComment(comment_id string) error
	UpdateComment(comment_id string) error
}
