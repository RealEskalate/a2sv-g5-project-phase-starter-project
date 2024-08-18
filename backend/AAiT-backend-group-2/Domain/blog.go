package domain

import "time"

type Blog struct {
	ID        string   `bson:"_id" json:"id"`
	Title     string   `bson:"title" json:"title"`
	Content   string   `bson:"content" json:"content"`
	Author    string   `bson:"author" json:"author"`
	Tags      []string `bson:"tags" json:"tags"`
	ViewCount int      `bson:"view_count" json:"view_count"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time    `bson:"updated_at" json:"updated_at"`
}

type BlogRepository interface {
	FindAll() ([]Blog, error)
	FindByID(id string) (*Blog, error)
	Save(blog *Blog) error
	Update(blog *Blog) error
	Delete(id string) error
}

type BlogUsecase interface {
	GetAllBlogs() ([]Blog, error)
	GetBlogByID(id string) (*Blog, error)
	CreateBlog(blog *Blog) error
	UpdateBlog(blog *Blog) error
	DeleteBlog(id string) error
}
