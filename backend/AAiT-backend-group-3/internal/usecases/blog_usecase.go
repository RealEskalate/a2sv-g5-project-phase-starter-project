package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
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

func (u *BlogUsecase) CreateBlog(blog *models.Blog) error {
	return u.blogRepo.CreateBlog(blog)
}

func (u *BlogUsecase) GetBlogByID(blogID primitive.ObjectID) (*models.Blog, error) {
	return u.blogRepo.GetBlogByID(blogID)
}

func (u *BlogUsecase) GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error) {
	return u.blogRepo.GetBlogs(filter, search, page, limit)
}

func (u *BlogUsecase) EditBlog(blogID primitive.ObjectID, newBlog *models.Blog) error {
	return u.blogRepo.EditBlog(blogID, newBlog)
}

func (u *BlogUsecase) DeleteBlog(blogID primitive.ObjectID) error {
	return u.blogRepo.DeleteBlog(blogID)
}

func (u *BlogUsecase) AddCommentToTheList(blogID primitive.ObjectID, commentID primitive.ObjectID) error {
	return u.blogRepo.AddCommentToTheList(blogID, commentID)
}

