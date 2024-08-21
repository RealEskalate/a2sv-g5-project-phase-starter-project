package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	repository_interface "AAIT-backend-group-3/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type BlogUsecase struct {
	blogRepo repository_interface.BlogRepositoryInterface
	
}
func NewBlogUsecase(blogRepo repository_interface.BlogRepositoryInterface) *BlogUsecase {
	return &BlogUsecase{
		blogRepo: blogRepo,
	}
}

func (u *BlogUsecase) CreateBlog(blog *models.Blog, authorID string) error {
	return u.blogRepo.CreateBlog(blog, authorID)
}

func (u *BlogUsecase) GetBlogByID(blogID string) (*models.Blog, error) {
	blog_id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}
	return u.blogRepo.GetBlogByID(blog_id)
}

func (u *BlogUsecase) GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error) {
	return u.blogRepo.GetBlogs(filter, search, page, limit)
}

func (u *BlogUsecase) EditBlog(blogID string, newBlog *models.Blog) error {
	return u.blogRepo.EditBlog(blogID, newBlog)
}

func (u *BlogUsecase) DeleteBlog(blogID string) error {
	return u.blogRepo.DeleteBlog(blogID)
}

func (u *BlogUsecase) AddCommentToTheList(blogID string, commentID string) error {
	return u.blogRepo.AddCommentToTheList(blogID, commentID)
}

