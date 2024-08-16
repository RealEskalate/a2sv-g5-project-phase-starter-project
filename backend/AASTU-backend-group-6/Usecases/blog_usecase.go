package usecases

import (
	domain "blogs/Domain"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUsecase(blogRepository domain.BlogRepository) *BlogUsecase {
	return &BlogUsecase{
		blogRepository: blogRepository,
	}
}
func (b *BlogUsecase) CreateBlog(user_id string, blog domain.Blog) error {

}
func (b *BlogUsecase) GetBlogByID(blog_id string) (domain.Blog, error) {

}
func (b *BlogUsecase) GetBlogs(pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogUsecase) UpdateBlogByID(user_id string, blog_id string, blog domain.Blog) error {

}
func (b *BlogUsecase) DeleteBlogByID(user_id string, blog_id string) error {

}
func (b *BlogUsecase) CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment domain.Comment) error {

}

func (b *BlogUsecase) SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogUsecase) FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogUsecase) GetMyBlogs(user_id string, pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogUsecase) GetMyBlogByID(user_id string, blog_id string) (domain.Blog, error) {

}
