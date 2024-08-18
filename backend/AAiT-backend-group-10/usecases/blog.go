package usecases

import (
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type IBlogUseCase interface {
	CreateBlog(blog *domain.Blog) (*domain.Blog, error)
	GetAllBlogs() ([]domain.Blog, error)
	GetBlogByID(id uuid.UUID) (*domain.Blog, error)
	UpdateBlog(blog *domain.Blog) error
	DeleteBlog(id uuid.UUID) error
	AddView(id uuid.UUID) error
}

type BlogUseCase struct {
	blogRepo interfaces.IBlogRepository
}

func NewBlogUseCase(repo interfaces.IBlogRepository) *BlogUseCase {
	return &BlogUseCase{blogRepo: repo}
}

func (b *BlogUseCase) CreateBlog(blog *domain.Blog) (*domain.Blog, error) {
	blog.ID = uuid.New()
	blog.CreatedAt = time.Now().UTC()
	blog.UpdatedAt = time.Now().UTC()
	return blog, b.blogRepo.Create(blog)
}

func (b *BlogUseCase) GetAllBlogs() ([]domain.Blog, error) {
	return b.blogRepo.FindAll()
}

func (b *BlogUseCase) GetBlogByID(id uuid.UUID) (*domain.Blog, error) {
	return b.blogRepo.FindByID(id)
}

func (b *BlogUseCase) UpdateBlog(blog *domain.Blog) error {
	blog.UpdatedAt = time.Now().UTC()
	return b.blogRepo.Update(blog)
}

func (b *BlogUseCase) DeleteBlog(id uuid.UUID) error {
	return b.blogRepo.Delete(id)
}

func (b *BlogUseCase) AddView(id uuid.UUID) error {
	return b.blogRepo.AddView(id)
}
