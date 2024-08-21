package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type IBlogRepository interface {
	Create(blog *domain.Blog) *domain.CustomError
	FindAll() ([]domain.Blog, *domain.CustomError)
	FindByID(id uuid.UUID) (*domain.Blog, *domain.CustomError)
	Update(blog *domain.Blog) *domain.CustomError
	Delete(id uuid.UUID) *domain.CustomError
	AddView(id uuid.UUID) *domain.CustomError
	Search(filter domain.BlogFilter) ([]domain.Blog, int, *domain.CustomError)
}
