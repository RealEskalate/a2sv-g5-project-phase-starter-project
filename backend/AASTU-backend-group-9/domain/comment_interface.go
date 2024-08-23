package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository interface {
	AddComment(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) error
	GetComments(ctx context.Context, post_id primitive.ObjectID) ([]Comment, error) 
	DeleteComment(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID, userID primitive.ObjectID) error
	UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) error
}
