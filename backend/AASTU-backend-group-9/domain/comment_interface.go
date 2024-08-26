package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository interface {
	AddComment(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) *Error
	GetComments(ctx context.Context, post_id primitive.ObjectID) ([]Comment, *Error)
	DeleteComment(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID, userID primitive.ObjectID) *Error
	UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *Comment) *Error
	DeleteComments(ctx context.Context, post_id primitive.ObjectID) *Error
	AddReply(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, reply *Comment) *Error
	IncrementCommentPopularity(ctx context.Context, post_id primitive.ObjectID, commentID primitive.ObjectID, metric string) *Error
}
