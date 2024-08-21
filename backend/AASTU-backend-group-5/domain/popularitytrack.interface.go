package domain

type BlogPopularityRepository interface {
	GetPopularBlogs(sortBy []SortBy, sortOrder []SortOrder) ([]Blog, error)
}
type BlogPopularityUsecase interface {
	GetSortedPopularBlogs(sortBy []SortBy, sortOrder []SortOrder) ([]Blog, error)
}
