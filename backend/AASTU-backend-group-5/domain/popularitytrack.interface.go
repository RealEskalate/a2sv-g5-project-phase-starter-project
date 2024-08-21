package domain

type BlogPopularityRepository interface {
	GetPopularBlogs(sortBy string, sortOrder int) ([]Blog, error)
}

type BlogPopularityUsecase interface {
	GetSortedPopularBlogs(sortBy string, sortOrder int) ([]Blog, error)
}
