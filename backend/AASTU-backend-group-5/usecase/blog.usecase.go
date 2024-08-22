package usecase

import (
	"time"

	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCase struct {
	BlogRepo domain.Blog_Repository_interface
}

func NewBlogUsecase(repo domain.Blog_Repository_interface) *BlogUseCase {
	return &BlogUseCase{
		BlogRepo: repo,
	}
}

func (uc *BlogUseCase) CreateBlog(iblog domain.PostBlog) (domain.Blog, error) {
	blog := domain.Blog{
		Title:     iblog.Title,
		Content:   iblog.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Tag:       iblog.Tag,
	}
	if blog.ID.IsZero() {
		blog.ID = primitive.NewObjectID()
	}
	createdBlog, err := uc.BlogRepo.CreateBlogDocument(blog)
	if err != nil {
		return domain.Blog{}, err
	}
	return createdBlog, nil
}

func (uc *BlogUseCase) GetBlogs(limit int, page_number int) ([]domain.Blog, error) {
	blogs, err := uc.BlogRepo.GetBlogDocuments(page_number, limit)
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}
func (uc *BlogUseCase) GetOneBlog(id string) (domain.Blog, error) {
	blogs, err := uc.BlogRepo.GetOneBlogDocument(id)
	if err != nil {
		return domain.Blog{}, err
	}
	return blogs, nil
}
func (uc *BlogUseCase) UpdateBlog(id string, updatedBlog domain.Blog) (domain.Blog, error) {
	blog, err := uc.BlogRepo.UpdateBlogDocument(id, updatedBlog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (uc *BlogUseCase) DeleteBlog(id string) error {
	return uc.BlogRepo.DeleteBlogDocument(id)
}
func (uc *BlogUseCase) FilterBlog(filters map[string]interface{}) ([]domain.Blog, error) {
	blogs, err := uc.BlogRepo.FilterBlogDocument(filters)
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}
