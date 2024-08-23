package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title" validate:"required"`
	Content      string             `json:"content"`
	Slug         string             `json:"slug"`
	PublishedAt  time.Time          `json:"published_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	IsPublished  bool               `json:"is_published"`
	Views        uint               `json:"views" bson:"views"`
	LikeCount    uint               `json:"like_count" bson:"likecount"`
	DislikeCount uint               `json:"dislike_count" bson:"dislikecount"`
	Popularity   uint               `json:"popularity"`
	AuthorID     primitive.ObjectID `json:"author_id"`
	AuthorName   string             `json:"author_name"`
	Tags         []string           `json:"tags"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Content   string             `json:"content" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	AuthorID  primitive.ObjectID `json:"author_id"`
	PostID    primitive.ObjectID `json:"post_id"`
}

type Tag struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty" bson:"_id,omitempty"`
	Name  string               `json:"name" validate:"required"`
	Slug  string               `json:"slug"`
	Posts []primitive.ObjectID `json:"post_id"`
}

type LikeDislike struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	PostID primitive.ObjectID `json:"post_id"`
	UserID primitive.ObjectID `json:"user_id"`
	IsLike bool               `json:"is_like" bson:"islike"` // true for like, false for dislike
	// other fields
}
