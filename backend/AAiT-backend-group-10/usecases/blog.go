package usecases

import (
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type BlogUseCase struct {
	BlogRepo interfaces.BlogRepoInterface
}

func (b *BlogUseCase) CreateBlog(blog *domain.Blog) (*domain.Blog, error) {
	blog.ID = uuid.New()
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	return blog, b.BlogRepo.Create(blog)
}

func (b *BlogUseCase) GetAllBlogs() ([]domain.Blog, error) {
	return b.BlogRepo.FindAll()
}

func (b *BlogUseCase) GetBlogByID(id uuid.UUID) (*domain.Blog, error) {
	return b.BlogRepo.FindByID(id)
}

func (b *BlogUseCase) UpdateBlog(blog *domain.Blog) error {
	blog.UpdatedAt = time.Now()
	return b.BlogRepo.Update(blog)
}

func (b *BlogUseCase) DeleteBlog(id uuid.UUID) error {
	return b.BlogRepo.Delete(id)
}

func (b *BlogUseCase) AddView(id uuid.UUID) error {
	return b.BlogRepo.AddView(id)
}