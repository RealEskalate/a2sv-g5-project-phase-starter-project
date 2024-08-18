package usecases

import (
	domain "blogs/Domain"
	"strings"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUsecase(blogRepository domain.BlogRepository) domain.BlogUsecase {
	return BlogUsecase{
		blogRepository: blogRepository,
	}
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
func (b BlogUsecase) DeleteBlogByID(user_id string, blog_id string, role string) error {
	var err error
	if strings.ToLower(role) != "admin"{
		err = b.blogRepository.DeleteBlogByID("", blog_id)
	}else{
		err = b.blogRepository.DeleteBlogByID(user_id, blog_id)
	}
	if err != nil{
		return err
	}
	return nil
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
func (b BlogUsecase) GetBlogs(pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
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


