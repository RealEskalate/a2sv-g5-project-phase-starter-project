package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
    // ID          string      `json:"id" gorm:"primaryKey"`
    ID          primitive.ObjectID   `json:"id"  bson:"_id,omitempty"`
    Title       string    `json:"title" validate:"required,min=5,max=255"`
    Content     string    `json:"content" validate:"required"`
    AuthorID    primitive.ObjectID      `json:"author_id"` // Foreign key to User model
    Tags        []string  `json:"tags" gorm:"type:text[]"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    //PublishedAt *time.Time `json:"published_at"`  Optional
}

type SearchBlogPost struct {
    Title string `json:"title"`
    Author string `json:"author"`
}

type BlogUsecaseInterface interface {
    CreateBlogPost(blog *BlogPost) (*BlogPost,error)
	GetAllBlogPosts() ([]BlogPost, error)
	GetBlogByID(id primitive.ObjectID) (*BlogPost, error)
	UpdateBlogPost(id primitive.ObjectID, blog *BlogPost) (*BlogPost, error)
	SearchBlogPosts(query *SearchBlogPost) ([]BlogPost, error)
	DeleteBlogPost(id primitive.ObjectID) error
}

type BlogRepositoryInterface interface {
    Save(blog *BlogPost) error
	FindAll() ([]BlogPost, error)
	FindByID(id int) (*BlogPost, error)
	Update(blog *BlogPost) (*BlogPost, error)
	Search(query string) ([]BlogPost, error)
	Delete(id int) error
}