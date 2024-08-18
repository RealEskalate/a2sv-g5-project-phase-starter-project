package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Content     string             `json:"content"`
	Slug        string             `json:"slug"`
	PublishedAt time.Time          `json:"published_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	IsPublished bool               `json:"is_published"`
	Views       uint               `json:"views"`
	AuthorID    primitive.ObjectID `json:"author_id"`
	LikeCount   uint               `json:"like_count"`
	LikeDislike []LikeDislike      `json:"like_dislike"`
	Comments    []Comment          `json:"comments"`
	Tags        []*Tag             `json:"tags"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Content   string             `json:"content"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	AuthorID  uint               `json:"author_id"`
	PostID    uint               `json:"post_id"`
}

type Tag struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Slug  string             `json:"slug"`
	Posts []*Post            `json:"posts"`
}

type LikeDislike struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	PostID uint               `json:"post_id"`
	UserID uint               `json:"user_id"`
	IsLike bool               `json:"is_like"` // true for like, false for dislike
	// other fields
}
