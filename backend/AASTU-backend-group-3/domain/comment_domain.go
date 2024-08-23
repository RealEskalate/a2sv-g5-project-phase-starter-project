package domain

import (
	"time"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment represents a comment made by a user on a post.
type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID             `json:"userId" bson:"userId"`
	PostID    primitive.ObjectID `json:"postId" bson:"postId"`
	Content   string             `json:"content" bson:"content"`
	LikesCount int               `json:"likesCount" bson:"likesCount"`
	Replies   []Reply            `json:"replies,omitempty" bson:"replies,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// unmarshal leave out updatedAt
type CommentResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID             `json:"userId" bson:"userId"`
	PostID    primitive.ObjectID `json:"postId" bson:"postId"`
	Content   string             `json:"content" bson:"content"`
	LikesCount int               `json:"likesCount" bson:"likesCount"`
	Replies   []Reply            `json:"replies,omitempty" bson:"replies,omitempty"`
}

func (c *Comment) MarshalJSON() ([]byte, error) {
	return json.Marshal(&CommentResponse{
		ID:        c.ID,
		UserID:    c.UserID,
		PostID:    c.PostID,
		Content:   c.Content,
		LikesCount: c.LikesCount,
		Replies:   c.Replies,
	})
}
type CommentRepository interface {
	CreateComment(comment *Comment) (*Comment, error)
	UpdateComment(comment *Comment) (*Comment, error)
	DeleteComment(id primitive.ObjectID) (*Comment, error)
	GetCommentByID(id primitive.ObjectID) (*Comment, error)
	GetCommentsByPostID(postID string, page, limit int64) ([]*Comment, error)

	CreateReply(reply *Reply) (*Reply, error)
	UpdateReply(reply *Reply) (*Reply, error)
	DeleteReply(id primitive.ObjectID) (*Reply, error)
	GetReplyByID(id primitive.ObjectID) (*Reply, error)
	GetRepliesByCommentID(commentID string, page, limit int64) ([]*Reply, error)

	LikeComment(commentID primitive.ObjectID, userID string) error
	UnlikeComment(commentID primitive.ObjectID, userID string) error

	LikeReply(replyID primitive.ObjectID, userID string) error
	UnlikeReply(replyID primitive.ObjectID, userID string) error
}

type CommentUsecase interface {
	CreateComment(comment *Comment) (*Comment, *CustomError)
	UpdateComment(comment *Comment, role_, userID string) (*Comment, *CustomError)
	DeleteComment(commentID, role_, userID string) (*Comment, *CustomError)
	GetCommentByID(commentID string) (*Comment, *CustomError)
	GetComments(postID string, page, limit int) ([]Comment, *CustomError)
	
	CreateReply(reply *Reply) (*Reply, *CustomError)
	UpdateReply(reply *Reply, userID string) (*Reply, *CustomError)
	DeleteReply(replyID , role_ , userID string) (*Reply, *CustomError)
	GetReplies(commentID string, page, limit int) ([]Reply, *CustomError)

	LikeComment(commentID string, userID string) *CustomError
	UnlikeComment(commentID string, userID string) *CustomError
	LikeReply(replyID string, userID string) *CustomError
	UnlikeReply(replyID string, userID string) *CustomError
}
