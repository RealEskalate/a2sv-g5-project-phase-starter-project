package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
)

type BlogRepository struct {
	PostCollection mongo.Collection
	env            interface{}
}

func NewBlogRepository(PostCollection mongo.Collection, env interface{}) *BlogRepository {
	return &BlogRepository{
		PostCollection: PostCollection,
		env:            env,
	}
}

func (b *BlogRepository) CreateBlog(user_id string, blog domain.Blog) error {

}
func (b *BlogRepository) GetBlogByID(blog_id string) (domain.Blog, error) {

}
func (b *BlogRepository) GetBlogs(pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogRepository) UpdateBlogByID(user_id string, blog_id string, blog *domain.Blog) error {

}
func (b *BlogRepository) DeleteBlogByID(user_id string, blog_id string) error {

}
func (b *BlogRepository) CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment *domain.Comment) error {

}
func (b *BlogRepository) SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogRepository) FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogRepository) GetMyBlogs(user_id string, pageNo string, pageSize string) ([]domain.Blog, Utils.Pagination, error) {

}
func (b *BlogRepository) GetMyBlogByID(user_id string, blog_id string) (domain.Blog, error) {

}
