package domain

import "github.com/google/uuid"

type UserRepository interface {
	FindById(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindAll() ([]User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id uuid.UUID) error
}

type BlogRepository interface {
	FindById(id uuid.UUID) (*Blog, error)
	FindAll() ([]Blog, error)
	Create(blog *Blog) (*Blog, error)
	Update(blog *Blog) (*Blog, error)
	Delete(id uuid.UUID) error
	Search(query string) ([]Blog, error)
	Filters(filters map[string]interface{}) ([]Blog, error)
	Like(id uuid.UUID) error
}

type CommentRepository interface {
	FindById(id uuid.UUID) (*Comment, error)
	FindAll() ([]Comment, error)
	Create(comment *Comment) (*Comment, error)
	Update(comment *Comment) (*Comment, error)
	Delete(id uuid.UUID) error
}