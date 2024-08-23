package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog represents a blog post with flexible content.
type Blog struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`          // Unique identifier for the blog
	Author        primitive.ObjectID `json:"ownerID" bson:"ownerID"` // ID of the blog author
	AuthorName    string             `json:"author_name" bson:"author_name"`
	Title         string             `json:"title" bson:"title"`           // Title of the blog
	Content       string             `json:"content" bson:"content"`       // Array of any type of content
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the blog was created
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at"` // Timestamp for when the blog was last updated
	Tags          []BlogTag          `json:"tags" bson:"tags"`             // Tags for categorizing the blog
	LikesCount    int                `json:"likes_count" bson:"likes_count"`
	CommentsCount int                `json:"comments_count" bson:"comments_count"`
	ViewsCount    int                `json:"views_count" bson:"views_count"`
}

type CreateBlogDTO struct {
	Title   string    `json:"title" bson:"title"` // Title of the blog
	Content string    `json:"content" bson:"content"`
	Tags    []BlogTag `json:"tags" bson:"tags"`
}

type UpdateBlogDTO struct {
	Title      string    `json:"title" bson:"title"` // Title of the blog
	Content    string    `json:"content" bson:"content"`
	Tags       []BlogTag `json:"tags" bson:"tags"`
	AuthorName string    `json:"author_name" bson:"author_name"`
}

type GetBlogDTO struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Author     primitive.ObjectID `json:"ownerID" bson:"ownerID"`
	AuthorName string             `json:"author_name" bson:"author_name"`
	Title      string             `json:"title" bson:"title"`
	Content    string             `json:"content" bson:"content"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Tags       []BlogTag          `json:"tags" bson:"tags"`

	ViewsCount    int `json:"views_count" bson:"views_count"`
	LikesCount    int `json:"likes_count" bson:"likes_count"`
	CommentsCount int `json:"comments_count" bson:"comments_count"`
}

type GetSingleBlogDTO struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Author     primitive.ObjectID `json:"ownerID" bson:"ownerID"`
	AuthorName string             `json:"author_name" bson:"author_name"`
	Title      string             `json:"title" bson:"title"`
	Content    string             `json:"content" bson:"content"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Tags       []BlogTag          `json:"tags" bson:"tags"`

	ViewsCount    int `json:"views_count" bson:"views_count"`
	LikesCount    int `json:"likes_count" bson:"likes_count"`
	CommentsCount int `json:"comments_count" bson:"comments_count"`

	Comments []Comment `json:"comments" bson:"comments"`
	Likes    []Like    `json:"likes" bson:"likes"`
	Views    []View    `json:"views" bson:"views"`
}

type BlogTag struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`    // Unique identifier for the blog
	Name string             `json:"name" bson:"name"` // Name of the blog
}

type BlogFilter struct {
	AuthorID  *primitive.ObjectID // Filter by Author ID
	Tags      []string            // Filter by Tags
	Title     *string             // Filter by Title (exact or partial match)
	DateRange *DateRange          // Filter by Creation Date Range
	Content   *string             // Filter by Content (exact or partial match)
	Keyword   *string             // Filter by keyword in title, content, or tags
}

// DateRange represents a time range for filtering
type DateRange struct {
	From time.Time // Start date for the range
	To   time.Time // End date for the range
}
