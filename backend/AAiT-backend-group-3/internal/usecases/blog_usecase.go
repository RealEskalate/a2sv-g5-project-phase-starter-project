package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/infrastructures/services"
	repository_interface "AAIT-backend-group-3/internal/repositories/interfaces"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type BlogUsecase struct {
	blogRepo repository_interface.BlogRepositoryInterface
	cacheService services.ICacheService
}
func NewBlogUsecase(blogRepo repository_interface.BlogRepositoryInterface, cacheSrv services.ICacheService) *BlogUsecase {
	return &BlogUsecase{
		blogRepo: blogRepo,
		cacheService: cacheSrv,
	}
}

func (u *BlogUsecase) CreateBlog(blog *models.Blog, authorID string) error {
	return u.blogRepo.CreateBlog(blog, authorID)
}

func (u *BlogUsecase) GetBlogByID(blogID string) (*models.Blog, error) {
	var cachedBlog models.Blog
	err := u.cacheService.GetBlog(blogID, &cachedBlog)
	if err == nil && (cachedBlog.ID).Hex() != "" {
		return &cachedBlog, nil
	}

	blog_id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}
	blog, err := u.blogRepo.GetBlogByID(blog_id)
	if err != nil {
		return nil, err
	}

	err = u.cacheService.SetBlog(blogID, blog, time.Hour)
	if err != nil {
		return nil, err
	}
	return blog, nil
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

