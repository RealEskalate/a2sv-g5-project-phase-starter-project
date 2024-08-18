package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type IBlogRepository interface {
	Create(blog *domain.Blog) error
	FindAll() ([]domain.Blog, error)
	FindByID(id uuid.UUID) (*domain.Blog, error)
	Update(blog *domain.Blog) error
	Delete(id uuid.UUID) error
	AddView(id uuid.UUID) error
}
