package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reply struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CommentId primitive.ObjectID `bson:"comment_id"` // Correct BSON field name for consistency
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BlogID    primitive.ObjectID `bson:"blog_id,omitempty"` // Blog ID or parent comment ID
	UserID    primitive.ObjectID `bson:"user_id"`           // Renamed for consistency
	Content   string             `bson:"content"`
	Replies   []Reply            `bson:"replies,omitempty"` // Holds replies
	CreatedAt time.Time          `bson:"created_at"`
}

type CommentUsecaseInterface interface {
	AddComment(comment *Comment) error
	GetCommentsByBlogID(blogID primitive.ObjectID) ([]Comment, error)
	UpdateComment(commentID primitive.ObjectID, content string) error
	DeleteComment(commentID primitive.ObjectID) error
	AddReply(commentID primitive.ObjectID, reply Reply) error
	UpdateReply(commentID, replyID primitive.ObjectID, content string) error
	DeleteReply(commentID, replyID primitive.ObjectID) error
}
