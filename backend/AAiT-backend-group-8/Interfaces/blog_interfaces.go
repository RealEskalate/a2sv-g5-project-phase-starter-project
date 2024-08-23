package Interfaces

import domain "AAiT-backend-group-8/Domain"

type IBlogRepository interface {
	Search(criteria *domain.SearchCriteria) ([]domain.Blog, error)
	Create(blog *domain.Blog) error
	UpdateCommentCount(id string, inc bool) error
	UpdateLikeCount(id string, inc bool) error
	FindAll(page int, pageSize int, sortBy string) ([]domain.Blog, error)
	FindByID(ID string) (*domain.Blog, error)
	Delete(ID string) error
	UpdateViewCount(id string) error
	Update(blog *domain.Blog) error
}

type IBlogUseCase interface {
	SearchBlog(criteria *domain.SearchCriteria) ([]domain.Blog, error)
	CreateBlog(blog *domain.Blog) error
	UpdateBlogViewCount(id string) error
	UpdateBlogCommentCount(id string, inc bool) error
	UpdateBlogLikeCount(id string, inc bool) error
	GetAllBlogs(page int, pageSize int, sortBy string) ([]domain.Blog, error)
	GetBlogByID(ID string) (*domain.Blog, error)
	DeleteBlog(ID string) error
	UpdateBlog(blog *domain.Blog) error
}
