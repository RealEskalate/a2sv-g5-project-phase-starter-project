package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title         string             `bson:"title" json:"title" binding:"required"`
	Content       string             `bson:"content" json:"content" binding:"required"`
	Author        string             `bson:"author" json:"author"`
	Tags          []string           `bson:"tags" json:"tags"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	LastUpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	ViewsCount    int                `bson:"views_count" json:"views_count"`
	LikesCount    int                `bson:"likes_count" json:"likes_count"`
	CommentsCount int                `bson:"comments_count" json:"comments_count"`
}

type View struct {
	BlogID primitive.ObjectID `bson:"blogid" json:"blogid"`
	User   string             `bson:"user" json:"user"`
}

type Like struct {
	BlogID primitive.ObjectID `bson:"blogid" json:"blogid"`
	User   string             `bson:"user" json:"user"`
	Like   bool               `bson:"like" json:"like"`
}

type Comment struct {
	BlogID  primitive.ObjectID `bson:"blogid" json:"blogid"`
	Author  string             `bson:"author" json:"author"`
	Content string             `bson:"content" json:"content"`
	Date    time.Time          `bson:"date" json:"date"`
}

type Tag struct {
	ID  string `bson:"_id" json:"id"`
}

type TagRepository interface {
	InsertTag(tag *Tag) error
	GetTags() ([]*Tag, error)
	DeleteTag(tag *Tag) error
	CheckTag(tags []string) error
}

type BlogRepository interface {
	InsertBlog(blog *Blog) (*Blog, error)
	GetBlogByID(id string) (*Blog, error)
	UpdateBlogByID(id string, blog *Blog) error
	DeleteBlogByID(id string) error
	SearchBlog(title, author string, tags []string) ([]*Blog, error)
	FilterBlog(tags []string, dateFrom, dateTo time.Time) ([]*Blog, error)
	AddView(view []*View) error
	AddLike(like *Like) error
	UpdateLike(like *Like) error
	AddComment(comment *Comment) error
	GetBlogsByPopularity(page, limit int, reverse bool) ([]*Blog, error)
	GetBlogsByRecent(page, limit int, reverse bool) ([]*Blog, error)
	GetLikebyAuthorAndBlogID(blogID string, author string) (*Like, error)
	GetBlogComments(blogID string) ([]*Comment, error)
	GetBlogLikes(blogID string) ([]*Like, error)
	IncrmentBlogViews(blogID string) error
	IncrmentBlogLikes(blogID string) error
	IncrmentBlogComments(blogID string) error
	DecrementBlogComments(blogID string) error
	DecrementBlogViews(blogID string) error
	DecrementBlogLikes(blogID string) error
	RemoveLike(blogID string, author string) error
	DeleteComment(commentID string) error
	GetCommentByID(commentID string) (*Comment, error)
	GetTotalBlogs() (int, error)

}

type BlogUsecase interface {
	InsertBlog(blog *Blog) (*Blog, error)
	GetBlogByID(id string) (*Blog, error)
	UpdateBlogByID(id string, blog *Blog, claim *LoginClaims) (*Blog, error)
	DeleteBlogByID(id string, claim *LoginClaims) error
	SearchBlog(title, author string, tags []string) ([]*Blog, error)
	FilterBlog(tags []string, dateFrom, dateTo time.Time) ([]*Blog, error)
	AddView(view []primitive.ObjectID, claim LoginClaims) error
	AddLike(like *Like) error
	AddComment(comment *Comment) error
	GetBlogs(sortBy string, page, limit int, reverse bool) ([]*Blog, int, error)
	GetBlogComments(blogID string) ([]*Comment, error)
	GetBlogLikes(blogID string) ([]*Like, error)
	RemoveLike(id string, claim *LoginClaims) error
	GenerateAiContent(prompt string) (string, error)
	DeleteComment(commentID string, claim *LoginClaims) error
	DeleteTag(tag *Tag, claim *LoginClaims) error
	InsertTag(tag *Tag, claim *LoginClaims) error
	GetTags() ([]*Tag, error)
	CheckTag(tags []string) error



}
