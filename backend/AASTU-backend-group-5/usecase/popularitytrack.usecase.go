package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
)

type BlogPopularityUsecase struct {
	repo domain.BlogPopularityRepository
}

func NewBlogPopularityUsecase(repo domain.BlogPopularityRepository) domain.BlogPopularityUsecase {
	return &BlogPopularityUsecase{repo: repo}
}

func (bpu *BlogPopularityUsecase) GetSortedPopularBlogs(sortBy []domain.SortBy, sortOrder []domain.SortOrder) ([]domain.Blog, error) {
	return bpu.repo.GetPopularBlogs(sortBy, sortOrder)
}
