package blog

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author"`
	Content   string             `json:"content" bson:"content"`
	AuthorID  primitive.ObjectID `json:"author_id" bson:"author_id"`
	Tags      []string           `json:"tags" bson:"tags"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type BlogUsecase interface {
	// CreateBlog(ctx context.Context, blog *Blog) error
	// UpdateBlog(ctx context.Context, authorID primitive.ObjectID, blogID primitive.ObjectID, updatedBlog *Blog) error
	// DeleteBlog(ctx context.Context, authorID primitive.ObjectID, blogID primitive.ObjectID) error
	// SearchBlog(ctx context.Context, blogTitle string, blogAuthor string) ([]*Blog, error)
	// GetBlogs(ctx context.Context) ([]*Blog, error)
	// GetBlog(ctx context.Context, blogid primitive.ObjectID) ([]*Blog, error)
}

type BlogRepository interface {
	// CreateBlog(ctx context.Context, blog *Blog) error
	// UpdateBlog(ctx context.Context, blogID primitive.ObjectID, authorID primitive.ObjectID, updatedBlog *Blog) error
	// DeleteBlog(ctx context.Context, blogID primitive.ObjectID, authorID primitive.ObjectID) error
	// SearchBlog(ctx context.Context, blogTitle string, blogAuthor string) ([]*Blog, error)
	// GetAllBlogs(ctx context.Context) ([]*Blog, error)
	// FilterBlogByTag(ctx context.Context, tag string) ([]*Blog, error)
}
