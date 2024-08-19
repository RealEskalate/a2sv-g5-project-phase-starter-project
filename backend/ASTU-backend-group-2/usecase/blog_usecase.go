package usecase

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
)

type blogUsecase struct {
	blogRepository domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository domain.BlogRepository, timeout time.Duration) domain.BlogUsecase {
	return &blogUsecase{
		blogRepository: blogRepository,
		contextTimeout: timeout,
	}
}

func (bu *blogUsecase) GetAllBlogs() ([]domain.Blog, error) {
	return nil, nil
}
func (bu *blogUsecase) GetBlogByID(blogID string) (domain.Blog, error) {
	return domain.Blog{}, nil
}
func (bu *blogUsecase) CreateBlog(newBlog domain.Blog) (domain.Blog, error) {
	return domain.Blog{}, nil
}
func (bu *blogUsecase) UpdateBlog(blogID string, updatedBlog domain.Blog) (domain.Blog, error) {
	return domain.Blog{}, nil
}
func (bu *blogUsecase) DeleteBlog(blogID string) error {
	return nil
}
func (bu *blogUsecase) GetComments(blogID string) ([]domain.Comment, error) {
	return []domain.Comment{}, nil
}
func (bu *blogUsecase) CreateComment(blogID string, comment domain.Comment) (domain.Comment, error) {
	return domain.Comment{}, nil
}
func (bu *blogUsecase) GetComment(blogID, commentID string) (domain.Comment, error) {
	return domain.Comment{}, nil
}
func (bu *blogUsecase) UpdateComment(blogID, commentID string, updatedComment domain.Comment) (domain.Comment, error) {
	return domain.Comment{}, nil
}
func (bu *blogUsecase) DeleteComment(blogID, commentID string) error {
	return nil
}
func (bu *blogUsecase) LikeBlog(blogID string, userID string) error {
	return nil
}
