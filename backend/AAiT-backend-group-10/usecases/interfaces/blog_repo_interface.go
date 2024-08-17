package interfaces

import "aait.backend.g10/domain"

type BlogRepoInterface interface {
	Create(blog *domain.Blog) error
	FindAll() ([]domain.Blog, error)
	FindByID(id string) (*domain.Blog, error)
	Update(blog *domain.Blog) error
	Delete(id string) error
}