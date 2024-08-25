package domain

import (
	"AAiT-backend-group-6/domain/dtos"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlogs = "blogs"
)
const (
	CollectionComments = "comments"
)
const (
	CollectionReactions = "reactions"
)

type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Author    string             `json:"author" bson:"author"`
	Content   string             `json:"content" bson:"content"`
	Title     string             `json:"title" bson:"title"`
	Tags      []string           `json:"tags" bson:"tags"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type Reaction struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID  primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	IsLiked bool               `json:"is_liked" bson:"is_liked"`
}

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Author    User               `json:"author" bson:"author"`
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type BlogUseCase interface {
	CreateBlog(c context.Context, blog *Blog) (Blog, error)
	GetBlog(c context.Context, id string) (*Blog, error)

	GetBlogs(c context.Context, pagination *Pagination) ([]*Blog, error)
	UpdateBlog(c context.Context, blog *Blog, blog_id string) error
	DeleteBlog(c context.Context, id string) error
	LikeBlog(c context.Context, blogID string, userID string) error
	UnlikeBlog(c context.Context, blogID string, userID string) error
	CommentBlog(c context.Context, blogID string, comment *Comment) error
}

type BlogRepository interface {
	CreateBlog(c context.Context, blog *Blog) (Blog, error)
	GetBlog(c context.Context, id string) (*Blog, error)
	GetBlogs(c context.Context, pagination *Pagination) ([]*Blog, error)
	UpdateBlog(c context.Context, blog *Blog) error
	DeleteBlog(c context.Context, id string) error
	LikeBlog(c context.Context, blogID string, userID string) error
	UnlikeBlog(c context.Context, blogID string, userID string) error
	CommentBlog(c context.Context, blogID string, comment *Comment) error
}

type CommentUseCase interface {
	CreateComment(c context.Context, comment *Comment, userID primitive.ObjectID) error
	GetComment(c context.Context, id string) (*Comment, error)
	UpdateComment(c context.Context, comment *dtos.UpdateDto, commentID primitive.ObjectID) error
	DeleteComment(c context.Context, id string) error
}

type CommentRepository interface {
	CreateComment(c context.Context, comment *Comment) error
	GetComment(c context.Context, id string) (*Comment, error)
	UpdateComment(c context.Context, comment *dtos.UpdateDto, commentID primitive.ObjectID) error
	DeleteComment(c context.Context, id string) error
}

type ReactionRepository interface {
	LikeBlog(c context.Context, blogID, userID primitive.ObjectID) error
	UnLikeBlog(c context.Context, blogID, userID primitive.ObjectID) error
	DeleteLike(c context.Context, blogID, userID primitive.ObjectID) error
	GetLike(ctx context.Context, userID, blogID primitive.ObjectID) (*Reaction, error)
}
type ReactionUsecase interface {
	LikeBlog(c context.Context, blogID, userID primitive.ObjectID) error
	UnLikeBlog(c context.Context, blogID, userID primitive.ObjectID) error
	DeleteLike(c context.Context, blogID, userID primitive.ObjectID) error
	GetLike(ctx context.Context, userID, blogID primitive.ObjectID) (*Reaction, error)
}
