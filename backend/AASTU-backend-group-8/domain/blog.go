package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title     string             `json:"title" validate:"required,min=5,max=255"`
	Content   string             `json:"content" validate:"required"`
	AuthorID  primitive.ObjectID `json:"author_id"` // Foreign key to User model
	Tags      []string           `json:"tags"`
	Veiws     int                `json:"views"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	//PublishedAt *time.Time `json:"published_at"`  Optional
}

type SearchBlogPost struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BlogUsecaseInterface interface {
	CreateBlogPost(blog *BlogPost) (string, error)
	GetAllBlogPosts(Pagination, string, int) ([]BlogPost, error)
	GetBlogByID(id primitive.ObjectID) (*BlogPost, error)
	UpdateBlogPost(id primitive.ObjectID, blog *BlogPost) (*BlogPost, error)
	// SearchBlogPosts(query *SearchBlogPost) ([]BlogPost, error) // Add this method
	DeleteBlogPost(id primitive.ObjectID) error
}

// domain/blog_repository_interface.go

type BlogRepositoryInterface interface {
	Save(blog *BlogPost) (string, error)
	GetAllBlog(Pagination, string, int) ([]BlogPost, error)
	GetBlogByID(id primitive.ObjectID) (*BlogPost, error)
	Update(id primitive.ObjectID, blog *BlogPost) (*BlogPost, error)
	// Search(title string) ([]BlogPost, error) // Add this method
	Delete(id primitive.ObjectID) error
}
