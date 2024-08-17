package usecases

import (
	domain "blogs/Domain"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
}

// CommentOnBlog implements domain.BlogRepository.
func (b BlogUsecase) CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment domain.Comment) error {
	panic("unimplemented")
}

// CreateBlog implements domain.BlogRepository.
func (b BlogUsecase) CreateBlog(user_id string, blog domain.Blog) (domain.Blog, error) {
	newBlog, err := b.blogRepository.CreateBlog(user_id, blog)
	if err != nil{
		return domain.Blog{}, err
	}
	return newBlog, nil
}

// DeleteBlogByID implements domain.BlogRepository.
func (b BlogUsecase) DeleteBlogByID(user_id string, blog_id string) error {
	panic("unimplemented")
}

// FilterBlogsByTag implements domain.BlogRepository.
func (b BlogUsecase) FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
}

// GetBlogByID implements domain.BlogRepository.
func (b BlogUsecase) GetBlogByID(blog_id string) (domain.Blog, error) {
	panic("unimplemented")
}

// GetBlogs implements domain.BlogRepository.
func (b *BlogUsecase) GetBlogs(pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	blogs, pagination, err := b.blogRepository.GetBlogs(pageNo, pageSize)
	if err != nil {
		return nil, domain.Pagination{}, err
	} else {
		return blogs, pagination, nil
	}
}

// GetMyBlogByID implements domain.BlogRepository.
func (b BlogUsecase) GetMyBlogByID(user_id string, blog_id string) (domain.Blog, error) {
	panic("unimplemented")
}

// GetMyBlogs implements domain.BlogRepository.
func (b BlogUsecase) GetMyBlogs(user_id string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
}

// SearchBlogByTitleAndAuthor implements domain.BlogRepository.
func (b BlogUsecase) SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
}

// UpdateBlogByID implements domain.BlogRepository.
func (b BlogUsecase) UpdateBlogByID(user_id string, blog_id string, blog domain.Blog) error {
	panic("unimplemented")
}

func NewBlogUsecase(blogRepository domain.BlogRepository) domain.BlogUsecase {
	return BlogUsecase{
		blogRepository: blogRepository,
	}
}
