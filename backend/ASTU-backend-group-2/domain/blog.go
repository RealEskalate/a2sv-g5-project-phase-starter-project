package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this structure will be used when responding for blog request
type Blog struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	AuthorID     primitive.ObjectID `json:"author_id" bson:"author_id"`
	Title        string             `json:"title" bson:"title" binding:"required"`
	Tags         []string           `json:"tags" bson:"tags"`
	Content      string             `json:"content" bson:"content" binding:"required"`
	ViewCount    int                `json:"view_count" bson:"view_count"`
	LikeCount    int                `json:"like_count" bson:"like_count"`
	DislikeCount int                `json:"dislike_count" bson:"dislike_count"`
	Comments     []Comment          `json:"comments" bson:"comments"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// defines the structure for the blogs that will be  received from the request when creating and updating
type BlogIn struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	AuthorID  primitive.ObjectID `json:"author_id" bson:"author_id"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	Tags      []string           `json:"tags" bson:"tags"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type BlogUpdate struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	Tags      []string  `json:"tags" bson:"tags"`
	Content   string    `json:"content" bson:"content" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// this structure when creating and updating comments
type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// user reaction to the blog if liked or disliked
type Reaction struct {
	BlogID   primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Liked    bool               `json:"liked" bson:"liked"`
	Disliked bool               `json:"disliked" bson:"disliked"`
	Date     time.Time          `json:"date" bson:"date"`
}


type BlogRepository interface {
	GetAllBlogs(c context.Context) ([]Blog, error)                             // Retrieves all blogs
	GetBlogByID(c context.Context, blogID string) (Blog, error)                // Retrieves a single blog by ID
	CreateBlog(c context.Context, newBlog *Blog) error                         // Creates a new blog
	UpdateBlog(c context.Context, blogID string, updatedBlog *Blog) error      // Updates an existing blog by ID
	DeleteBlog(c context.Context, blogID string) error                         // Deletes a blog by ID
	GetComments(c context.Context, blogID string) ([]Comment, error)           // Retrieves comments for a specific blog
	CreateComment(c context.Context, blogID string, comment *Comment) error    // Adds a comment to a blog
	GetComment(c context.Context, blogID, commentID string) (Comment, error)   // Retrieves a specific comment by ID
	UpdateComment(c context.Context, blogID, commentID string, updatedComment *Comment) error // Updates a comment by ID
	DeleteComment(c context.Context, blogID, commentID string) error           // Deletes a comment by ID
	LikeBlog(c context.Context, blogID, userID string) error                   // Adds a like to a blog
}

type BlogUsecase interface {
	GetAllBlogs() ([]Blog, error)                             // Retrieves all blogs
	GetBlogByID(blogID string) (Blog, error)                  // Retrieves a single blog by ID
	CreateBlog(newBlog Blog) (Blog, error)                    // Creates a new blog
	UpdateBlog(blogID string, updatedBlog Blog) (Blog, error) // Updates an existing blog by ID
	DeleteBlog(blogID string) error                           // Deletes a blog by ID
	GetComments(blogID string) ([]Comment, error)             // Retrieves comments for a specific blog
	CreateComment(blogID string, comment Comment) (Comment, error) // Adds a comment to a blog
	GetComment(blogID, commentID string) (Comment, error)     // Retrieves a specific comment by ID
	UpdateComment(blogID, commentID string, updatedComment Comment) (Comment, error) // Updates a comment by ID
	DeleteComment(blogID, commentID string) error             // Deletes a comment by ID
	LikeBlog(blogID string, userID string) error              // Adds a like to a blog
}

